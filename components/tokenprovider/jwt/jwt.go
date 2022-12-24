package jwt

import (
	"FoodDelivery/components/tokenprovider"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type jwtProvider struct {
	secret string
}

func (j *jwtProvider) Validate(token string) (*tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(token, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil || !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)

	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	return &claims.Payload, nil
}

func (j jwtProvider) String() string {
	return "JWT implement Provider"
}

func NewTokenJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.RegisteredClaims
}

func (j *jwtProvider) Generate(
	data tokenprovider.TokenPayload,
	expiry int) (*tokenprovider.Token, error) {
	// generate token
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Second * time.Duration(expiry))),
			IssuedAt:  jwt.NewNumericDate(time.Now().Local()),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))

	if err != nil {
		return nil, err
	}

	return &tokenprovider.Token{
		Token:   myToken,
		Created: time.Now(),
		Expiry:  expiry,
	}, nil
}