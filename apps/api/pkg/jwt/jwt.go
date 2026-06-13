package jwt

import (
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "github.com/google/uuid"
)

type TokenPayload struct {
    UserID uuid.UUID `json:"user_id"`
    Role   string    `json:"role"`
}

type TokenMaker interface {
    CreateToken(userID uuid.UUID, role string, duration time.Duration) (string, *jwt.RegisteredClaims, error)
    VerifyToken(token string) (*TokenPayload, error)
}

type JWTMaker struct {
    secretKey string
}

func NewJWTMaker(secretKey string) (TokenMaker, error) {
    if len(secretKey) < 32 {
        return nil, errors.New("invalid key size: must be at least 32 characters")
    }
    return &JWTMaker{secretKey}, nil
}

type customClaims struct {
    UserID uuid.UUID `json:"user_id"`
    Role   string    `json:"role"`
    jwt.RegisteredClaims
}

func (m *JWTMaker) CreateToken(userID uuid.UUID, role string, duration time.Duration) (string, *jwt.RegisteredClaims, error) {
    claims := &customClaims{
        UserID: userID,
        Role:   role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
            Subject:   userID.String(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenStr, err := token.SignedString([]byte(m.secretKey))
    return tokenStr, &claims.RegisteredClaims, err
}

func (m *JWTMaker) VerifyToken(token string) (*TokenPayload, error) {
    keyFunc := func(token *jwt.Token) (interface{}, error) {
        _, ok := token.Method.(*jwt.SigningMethodHMAC)
        if !ok {
            return nil, errors.New("invalid token signing method")
        }
        return []byte(m.secretKey), nil
    }

    jwtToken, err := jwt.ParseWithClaims(token, &customClaims{}, keyFunc)
    if err != nil {
        return nil, err
    }

    claims, ok := jwtToken.Claims.(*customClaims)
    if !ok {
        return nil, errors.New("invalid token claims")
    }

    return &TokenPayload{
        UserID: claims.UserID,
        Role:   claims.Role,
    }, nil
}
