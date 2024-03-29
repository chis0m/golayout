package services

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"go-layout/config"
	"go-layout/internal/appx"
	"go-layout/internal/models"
	"go-layout/internal/types"
	"go-layout/pkg/token"
	"go-layout/utils"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type AuthServiceInterface interface {
	SignUp(raw types.UserSignUpRequestDTO) (*models.User, error)
	Login(agent string, clientIp string, raw types.UserLoginRequestDTO) (*types.UserLoginResponse, error)
	Renew(raw types.RefreshTokenRequestDTO) (*types.RefreshTokenResponse, error)
	Logout(raw types.RefreshTokenRequestDTO) error
}

type AuthService struct {
	env        *config.Config
	db         *gorm.DB
	tokenMaker token.Maker
}

func NewAuthService(env *config.Config, db *gorm.DB, tokenMaker token.Maker) AuthServiceInterface {
	return &AuthService{
		env:        env,
		db:         db,
		tokenMaker: tokenMaker,
	}
}

func (as *AuthService) SignUp(raw types.UserSignUpRequestDTO) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(raw.Password)
	if err != nil {
		log.Err(err).Msg("Error while hashing password")
		return nil, appx.NewError(http.StatusInternalServerError, "internal server error")
	}
	user := &models.User{
		FirstName:    utils.PointerString(raw.FirstName),
		LastName:     utils.PointerString(raw.LastName),
		Email:        raw.Email,
		PasswordHash: hashedPassword,
	}
	err = as.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (as *AuthService) Login(agent string, clientIp string, raw types.UserLoginRequestDTO) (*types.UserLoginResponse, error) {
	var user models.User
	err := as.db.Where("email = ?", raw.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appx.NewError(http.StatusBadRequest, "user not found")
		} else {
			return nil, fmt.Errorf("could not fetch user %+v", err)
		}
	}

	err = utils.VerifyPassword(user.PasswordHash, raw.Password)
	if err != nil {
		log.Err(err).Msg("Password verification failed")
		return nil, fmt.Errorf("invlid credentials %s", err)
	}
	claims := token.Claims{
		Issuer:   as.env.App.Name,
		Subject:  user.Email,
		Audience: as.env.App.Url,
		Duration: as.env.Token.AccessDuration,
		Data:     map[string]interface{}{"type": "original"},
	}
	accessToken, accessPayload, err := as.tokenMaker.CreateToken(claims)
	if err != nil {
		log.Err(err).Msg("AuthService Login: token creation failed")
		return nil, fmt.Errorf("internal server error %s", err)
	}

	claims = token.Claims{
		Issuer:   as.env.App.Name,
		Subject:  user.Email,
		Audience: as.env.App.Url,
		Duration: as.env.Token.RefreshDuration,
		Data:     make(map[string]interface{}),
	}
	refreshToken, refreshPayload, err := as.tokenMaker.CreateToken(claims)

	session := models.Session{
		ID:           refreshPayload.ID,
		UserId:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    agent,
		ClientIP:     clientIp,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.Exp,
	}

	err = as.db.Create(session).Error
	if err != nil {
		log.Err(err).Msg("AuthService Login: failed inserting session to database")
		return nil, appx.NewError(http.StatusBadRequest, "internal server error")
	}

	return &types.UserLoginResponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.Exp,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.Exp,
		User: types.UserResponse{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		},
	}, nil
}

func (as *AuthService) Renew(raw types.RefreshTokenRequestDTO) (*types.RefreshTokenResponse, error) {
	refreshPayload, err := as.tokenMaker.VerifyToken(raw.RefreshToken)
	if err != nil {
		return nil, appx.NewError(http.StatusBadRequest, "invalid access token")
	}

	var session models.Session
	err = as.db.Where("id = ?", refreshPayload.ID).First(&session).Error
	if err != nil {
		log.Err(err).Msg("AuthService Logout: Failed to fetch session from DB")
		return nil, appx.NewError(http.StatusBadRequest, "internal server error")
	}

	if session.IsBlocked {
		return nil, appx.NewError(http.StatusBadRequest, "this session is blocked")
	}

	var user models.User
	err = as.db.Where("id = ?", session.UserId).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appx.NewError(http.StatusBadRequest, "user not found")
		} else {
			return nil, fmt.Errorf("could not fetch user %+v", err)
		}
	}
	if refreshPayload.Sub != user.Email {
		return nil, appx.NewError(http.StatusBadRequest, "incorrect session user")
	}

	if session.RefreshToken != raw.RefreshToken {
		return nil, appx.NewError(http.StatusBadRequest, "mismatched session token")
	}

	if time.Now().After(session.ExpiresAt) {
		return nil, appx.NewError(http.StatusBadRequest, "expired session")
	}

	claims := token.Claims{
		Issuer:   as.env.App.Name,
		Subject:  user.Email,
		Audience: as.env.App.Url,
		Duration: as.env.Token.AccessDuration,
		Data:     map[string]interface{}{"type": "renewed"},
	}
	newAccessToken, newAccessPayload, err := as.tokenMaker.CreateToken(claims)
	return &types.RefreshTokenResponse{
		AccessToken:          newAccessToken,
		AccessTokenExpiresAt: newAccessPayload.Exp,
	}, nil
}

func (as *AuthService) Logout(raw types.RefreshTokenRequestDTO) error {
	refreshPayload, err := as.tokenMaker.VerifyToken(raw.RefreshToken)
	if err != nil {
		return appx.NewError(http.StatusBadRequest, "invalid access token")
	}
	var session models.Session
	err = as.db.Where("id = ?", refreshPayload.ID).First(&session).Error
	if err != nil {
		log.Err(err).Msg("AuthService Logout: Failed to fetch session from DB")
		return appx.NewError(http.StatusBadRequest, "internal server error")
	}
	result := as.db.Delete(&session)
	if result.Error != nil {
		log.Err(result.Error).Msg("AuthService Logout: Failed to Delete session from DB")
		return result.Error
	}

	if result.RowsAffected == 0 {
		log.Error().Msg("AuthService Logout: session data was not deleted")
		return appx.NewError(http.StatusBadRequest, "failed to logout")
	}

	return nil
}
