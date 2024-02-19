package token

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"github.com/o1egl/paseto"
)

type PasetoPublic struct {
	paseto *paseto.V2
}

func NewPasetoPublic() *PasetoPublic {
	return &PasetoPublic{
		paseto: paseto.NewV2(),
	}
}

func (maker *PasetoPublic) CreateToken(claims Claims, privateKey string) (string, *Payload, error) {
	payload, err := NewPayload(claims)
	if err != nil {
		return "", nil, err
	}
	token, err := maker.paseto.Sign(privateKey, payload, nil)
	if err != nil {
		return "", nil, err
	}

	return token, payload, nil
}

func (maker *PasetoPublic) VerifyToken(token string, publicKey string) (*Payload, error) {
	payload := &Payload{}
	err := maker.paseto.Verify(token, publicKey, &payload, nil)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func GenerateKeyPair() (privateKey string, publicKey string, err error) {
	pbk, pvk, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return "", "", fmt.Errorf("error generating key pair: %v\n", err)
	}
	privateKey = string(pvk)
	publicKey = string(pbk)
	return
}
