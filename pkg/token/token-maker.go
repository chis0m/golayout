package token

type Maker interface {
	CreateToken(claims Claims) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
