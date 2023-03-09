package entities

import (
	"github.com/golang-jwt/jwt/v4"
)

// @Success 200 {object} responseWrapper{data=entities.Polygon2, error=false, message="İşlem başarılı"}
// @Failure 400 {object} responseWrapper{data=null, error=true, message="Hatalı istek"}

type ResponseWrapper struct {
	Data       interface{} `json:"data" swaggertype:"object" x-nullable:"true"`
	StatusCode int         `json:"statuscode" example:"200"`
	Error      bool        `json:"error" example:"false"`
	Message    string      `json:"message" example:"islem basarili"`
}

type ErrorResponseWrapper struct {
	StatusCode int    `json:"statuscode" example:"400"`
	Error      bool   `json:"error" example:"true"`
	Message    string `json:"message" example:"error message (islem basarili degil)"`
}

type TokenClaim struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Exp     int64  `json:"exp"`
	Iat     int64  `json:"iat"`
	jwt.RegisteredClaims
}
