package jwt

import (
	ent "MapProject/internal/entities"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	//jwt "github.com/golang-jwt/jwt/v4"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"os"
	"strings"
	"time"
)

type TokenMetadata struct {
	Exp int64 `json:"exp"`
}

func init() {
	err := godotenv.Load("./prod.env")
	println(err)
}

func GenerateAccessToken(claims *ent.TokenClaim) (string, error) {
	//println(os.Getenv("APP_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	claims.Iat = time.Now().Unix()
	claims.Exp = time.Now().Local().Add(time.Second * time.Duration(18000)).Unix()
	signingKey := []byte(os.Getenv("APP_SECRET"))
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func GenerateRefreshToken(claims *ent.TokenClaim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	claims.Iat = time.Now().Unix()
	claims.Exp = time.Now().Add(time.Hour * 168).Unix()
	signingKey := []byte(os.Getenv("APP_SECRET"))
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func ExtractToken(c *fiber.Ctx) (returnstring string) {
	bearToken := c.Get("Authorization")
	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}
	return
}

func VerifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, JwtKeyFunc)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func JwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}

// JWTMiddleware verifies the JWT token in the request and returns a 401 Unauthorized error if the token is invalid

func JWTMiddleware(c *fiber.Ctx) error {
	// Get the Authorization header from the request
	header := c.Get("Authorization")
	if header == "" {
		//c.Status(fiber.StatusUnauthorized).Send(errors.New("Authorization header is missing"))
		c.Status(fiber.StatusUnauthorized).Send([]byte("Authorization basligi bulunamadi"))
		return errors.New("Authorization basligi bulunamadi")
	}

	// Check if the header starts with "Bearer "
	if !strings.HasPrefix(header, "accesstoken : ") {
		c.Status(fiber.StatusUnauthorized).Send([]byte("Authorization basligi 'accesstoken ' ile baslamali"))
		return errors.New("Authorization basligi 'accesstoken ' ile baslamali")
	}

	// Extract the JWT token from the header
	tokenString := header[14:]
	if tokenString == "" {
		c.Status(fiber.StatusUnauthorized).Send([]byte("JWT token eksik"))
		return errors.New("JWT token eksik")
	}

	// Parse the JWT token
	secret := os.Getenv("app_secret")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Beklenmeyen giris yontemi: %v", token.Header["accesstoken : "])
		}
		return []byte(secret), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized).Send([]byte("jwt parse edilirken bir sorun ile karsilasildi"))
		return errors.New("JWT parse edilirken hata alindi")
	}

	// Check if the token is expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				c.Status(fiber.StatusUnauthorized).Send([]byte("JWT token suresi gecmis"))
				return errors.New("JWT token suresi gecmis")
			}
		}
	}
	// Check if the token is valid
	if !token.Valid {
		c.Status(fiber.StatusUnauthorized).Send([]byte("JWT gecersiz"))
		return errors.New("JWT gecersiz")
	}

	// If the token is valid and not expired, continue processing the request

	// Extract the "id" field from the claims
	claims, _ := token.Claims.(jwt.MapClaims)
	id, ok := claims["id"].(float64)
	if !ok {
		c.Status(fiber.StatusBadRequest).Send([]byte("JWT icerisinde id fieldi yok"))
		return errors.New("JWT icerisinde id field yok")
	}

	//c.Locals("id", id)
	idInt := int64(id)
	c.Locals("id", idInt)
	c.Next()

	return nil
}

func JwtMiddlewareWrapper(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return JWTMiddleware(c)
	}
}
