package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/TDiblik/project-template/api/constants"
	"github.com/TDiblik/project-template/api/models"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func CompareHashAndPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

type JWTInfo struct {
	Sub           string
	UserId        uuid.UUID
	UserEmail     string
	UserFirstName string
	UserLastName  string
	UserHandle    string
	Exp           int64
}

func GenerateJWT(user models.UserModelDB) (string, error) {
	token_insecure := jwt.New(jwt.SigningMethodHS256)

	claims := token_insecure.Claims.(jwt.MapClaims)
	claims["sub"] = "project-template.inc"
	claims["jti"] = uuid.New()
	claims["user_id"] = user.Id.String()
	claims["user_email"] = user.Email
	claims["user_first_name"] = user.FirstName
	claims["user_last_name"] = user.LastName
	claims["user_handle"] = user.Handle
	if EnvData.Debug {
		claims["exp"] = time.Now().AddDate(20, 0, 0).Unix()
	} else {
		claims["exp"] = time.Now().AddDate(0, 3, 0).Unix()
	}

	token, err := token_insecure.SignedString(EnvData.AUTH_SECRET_BYTES)

	return token, err
}

func TokenClaimsToJwtInfo(claims jwt.MapClaims) (*JWTInfo, error) {
	user_id, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		return nil, err
	}

	return &JWTInfo{
		Sub:           claims["sub"].(string),
		UserId:        user_id,
		UserEmail:     claims["user_email"].(string),
		UserFirstName: claims["user_first_name"].(string),
		UserLastName:  claims["user_last_name"].(string),
		UserHandle:    claims["user_handle"].(string),
		Exp:           int64(claims["exp"].(float64)),
	}, nil
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return EnvData.AUTH_SECRET_BYTES, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to validate token: %w", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unable to get claims from the token")
	}

	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return nil, fmt.Errorf("token has expired")
		}
	} else {
		return nil, fmt.Errorf("token does not contain a valid 'exp' claim")
	}

	return claims, nil
}

func GetJWTFromLocals(c fiber.Ctx) (*JWTInfo, error) {
	val := c.Locals(constants.TOKEN_CLAIMS_LOCALS_KEY)
	if val == nil {
		return nil, errors.New("JWT claims not found in context locals")
	}
	jwtInfo, ok := val.(*JWTInfo)
	if !ok {
		return nil, errors.New("invalid JWT claims type in context locals")
	}
	return jwtInfo, nil
}

func SetJWTToLocals(c fiber.Ctx, tokenInfo *JWTInfo) {
	c.Locals(constants.TOKEN_CLAIMS_LOCALS_KEY, tokenInfo)
}

type ErrJWTNoToken struct{}
type ErrJWTInvalidToken struct{}
type ErrJWTTokenConversion struct{}

func (e *ErrJWTNoToken) Error() string {
	return "jwt auth token is missing"
}
func (e *ErrJWTInvalidToken) Error() string {
	return "jwt auth token is invalid"
}
func (e *ErrJWTTokenConversion) Error() string {
	return "jwt auth token failed to convert to token claims"
}

var JWTNoTokenErr *ErrJWTNoToken = &ErrJWTNoToken{}
var JWTInvalidTokenErr *ErrJWTInvalidToken = &ErrJWTInvalidToken{}
var JWTConversionErr *ErrJWTTokenConversion = &ErrJWTTokenConversion{}

func GetUserInfoFromJWT(c fiber.Ctx) (*JWTInfo, error) {
	tokenRaw := c.Get(constants.TOKEN_HEADER_NAME)
	if len(tokenRaw) == 0 {
		return nil, JWTNoTokenErr
	}
	tokenClaims, err := ValidateJWT(tokenRaw)
	if err != nil {
		return nil, JWTInvalidTokenErr
	}
	tokenInfo, err := TokenClaimsToJwtInfo(tokenClaims)
	if err != nil {
		return nil, JWTConversionErr
	}
	return tokenInfo, nil
}
