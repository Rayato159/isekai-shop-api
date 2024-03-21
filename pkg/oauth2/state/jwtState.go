package state

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtState struct {
	secret   []byte
	expireAt *jwt.NumericDate
	issuer   string
}

func NewJwtState(secret []byte, expireAt time.Duration, issuer string) State {
	return &jwtState{
		secret:   secret,
		expireAt: jwt.NewNumericDate(time.Now().Add(expireAt * time.Second)),
		issuer:   issuer,
	}
}

func (s *jwtState) GenerateRandomState() (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: s.expireAt,
		Issuer:    s.issuer,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(s.secret)
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (s *jwtState) ParseState(state string) error {
	_, err := jwt.Parse(state, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return s.secret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return errors.New("malformed token")
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return errors.New("expired token")
		} else {
			return errors.New("unknown claims type, cannot proceed")
		}
	}

	return nil
}
