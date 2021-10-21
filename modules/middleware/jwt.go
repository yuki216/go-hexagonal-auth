package middlewareJWT

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-hexagonal/config"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
	"go-hexagonal/constants"
)


type JWTInterface interface {
	GenerateJWTToken(ctx context.Context, request JWTRequest) (string, error)
	ExtractJWTClaims(ctx context.Context, authBearer string) (claims *JWTClaims, err error)
	ValidateTokenIssuer(claims *JWTClaims) (err error)
	ValidateTokenExpire(ctx context.Context, claims *JWTClaims, reqToken string) (err error)

	SaveTokenToRedis(ctx context.Context, id, hour int, token, authKey string) error
	GetTokenFromRedis(ctx context.Context, id int, authKey string) (string, error)
	DeleteTokenFromRedis(ctx context.Context, id int, authKey string) error
}

const (
	AuthKeyUser  = "auth-user"
	AuthKeyAdmin = "auth-admin"
)

type jwtObj struct {
	config *config.JWTConfig
	redis  *redis.Client
}

func NewJWT(cfg *config.JWTConfig, redis *redis.Client) JWTInterface {
	return &jwtObj{
		config: cfg,
		redis:  redis,
	}
}

func (j *jwtObj) GenerateJWTToken(ctx context.Context, request JWTRequest) (string, error) {
	JWTSignatureKey := []byte(j.config.Secret)
	expireTime := time.Now().Add(time.Duration(j.config.TokenLifeTimeHour) * time.Hour)

	key := AuthKeyUser
	claims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    j.config.Issuer,
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		ID:           request.ID,
		Email:        request.Email,
		Name:         request.Name,
	}

	token := jwt.NewWithClaims(
		JWTSigningMethod,
		claims,
	)

	// create token client
	signedToken, err := token.SignedString(JWTSignatureKey)
	if err != nil {
		return "", err
	}

	// Save token to redis
	err = j.SaveTokenToRedis(ctx, request.ID, j.config.TokenLifeTimeHour, signedToken, key)
	if err != nil {
		err = constants.ErrSaveTokenToRedis
		log.Error(constants.ErrSaveTokenToRedis)
		return "", err
	}
	return signedToken, nil
}

func (j *jwtObj) ExtractJWTClaims(ctx context.Context, token string) (claims *JWTClaims, err error) {
	// check authorization
	splitToken := strings.Split(token, bearer)
	if len(splitToken) != 2 {
		return nil, constants.ErrTokenIsRequired
	}
	reqToken := strings.TrimSpace(splitToken[1])

	t, err := jwt.ParseWithClaims(reqToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.config.Secret, nil
	})

	if err != nil && err.Error() != constants.ErrKeyIsNotInvalidType.Error() {
		return nil, err
	}

	claims = t.Claims.(*JWTClaims)
	// Validate Issuer Token
	err = j.ValidateTokenIssuer(claims)
	if err != nil {
		return nil, err
	}

	// Validate token expire
	err = j.ValidateTokenExpire(ctx, claims, reqToken)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

// ValidateTokenIssuer is for validate token issuer
func (j *jwtObj) ValidateTokenIssuer(claims *JWTClaims) (err error) {
	if claims.Issuer != j.config.Issuer {
		return err
	}
	return nil
}

// ValidateTokenExpire is for validate Token Expire
func (j *jwtObj) ValidateTokenExpire(ctx context.Context, claims *JWTClaims, reqToken string) (err error) {
	// check token to redis
	key := AuthKeyUser
	token, err := j.GetTokenFromRedis(ctx, claims.ID, key)
	if err != nil {
		log.Error(constants.ErrGetTokenToRedis)
		return constants.ErrGetTokenToRedis
	}
	if token == "" {
		return constants.ErrTokenAlreadyExpired
	}

	if token != reqToken {
		return constants.ErrTokenAlreadyExpired
	}

	return nil
}

func (j *jwtObj) SaveTokenToRedis(ctx context.Context, id, hour int, token, authKey string) error {
	key := fmt.Sprintf("%s:%d", authKey, id)
	ttl := time.Duration(hour) * time.Hour
	err := j.redis.Set(ctx, key, token, ttl).Err()
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (j *jwtObj) GetTokenFromRedis(ctx context.Context, id int, authKey string) (string, error) {
	key := fmt.Sprintf("%s:%d", authKey, id)
	val, err := j.redis.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (j *jwtObj) DeleteTokenFromRedis(ctx context.Context, id int, authKey string) error {
	key := fmt.Sprintf("%s:%d", authKey, id)
	_, err := j.redis.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	return nil
}
