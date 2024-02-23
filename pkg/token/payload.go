package token

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Claims struct {
	Issuer   string                 `json:"issuer"`
	Subject  string                 `json:"subject"`
	Audience string                 `json:"audience"`
	Duration time.Duration          `json:"duration"`
	Data     map[string]interface{} `json:"data,omitempty"`
}

type Payload struct {
	ID   uuid.UUID              `json:"id"`
	Iss  string                 `json:"iss"`
	Sub  string                 `json:"sub"`
	Aud  string                 `json:"aud"`
	Exp  time.Time              `json:"exp"`
	Data map[string]interface{} `json:"data,omitempty"`
}

func NewPayload(claim Claims) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		ID:   tokenID,
		Iss:  claim.Issuer,
		Sub:  claim.Subject,
		Aud:  claim.Audience,
		Exp:  time.Now().Add(claim.Duration),
		Data: claim.Data,
	}, nil
}

func (payload *Payload) IsValid() error {
	if time.Now().After(payload.Exp) || time.Now().Equal(payload.Exp) {
		return errors.New("token has expired")
	}
	return nil
}
