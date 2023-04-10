package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
)

type Payload struct {
	ID        uuid.UUID        `json:"id"`
	Username  string           `json:"username"`
	IssuedAt  *jwt.NumericDate `json:"issued_at"`
	ExpiresAt *jwt.NumericDate `json:"expires_at"`
	NotBefore *jwt.NumericDate `json:"not_before"`
	Issuer    string           `json:"issuer"`
	Subject   string           `json:"subject"`
	Audience  jwt.ClaimStrings `json:"audience"`
}

func (payload Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return payload.ExpiresAt, nil
}

func (payload Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return payload.IssuedAt, nil
}

func (payload Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return payload.NotBefore, nil
}

func (payload Payload) GetIssuer() (string, error) {
	return payload.Issuer, nil
}

func (payload Payload) GetSubject() (string, error) {
	return payload.Subject, nil
}

func (payload Payload) GetAudience() (jwt.ClaimStrings, error) {
	return payload.Audience, nil
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
	}
	return payload, nil
}
