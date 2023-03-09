package auth

import (
	mod "MapProject/api/models/auth"
	"MapProject/internal/common"
	"MapProject/internal/entities"
	"MapProject/pkg/utils/jwt"
	"github.com/gofiber/fiber/v2"
)

var (
	Responser  interface{}
	Data       interface{}
	Result     int8
	StatusCode int
	Claim      entities.Auth
)

// Auth Login for login.
// @Description Login Function
// @Summary Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body entities.Login true "Login"
// @Success 200 {object} entities.ResponseWrapper
// @Failure 502 {object} entities.ErrorResponseWrapper
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
	dat := entities.Login{}
	errParser := c.BodyParser(&dat)
	if errParser != nil {
		StatusCode = fiber.StatusBadGateway
		Responser = common.HTTPResponser(StatusCode, true, errParser.Error())
	}
	Result, Claim = mod.Login(dat)
	if Result == -1 {
		StatusCode = fiber.StatusInternalServerError
		Responser = common.HTTPResponser(StatusCode, true, "Veri tabaninda hata olustu")
	} else if Result == 0 {
		StatusCode = fiber.StatusUnauthorized
		Responser = common.HTTPResponser(StatusCode, true, "Kullanici bulunamadi")
	} else {
		claims := entities.TokenClaim{
			Id:      Claim.AuthId, //bu authid models/auth/auth.go icindeki scan fonksiyonu sayesinde structa yazilan degerden geliyor
			Name:    Claim.Name,
			Surname: Claim.Surname,
			Exp:     0,
			Iat:     0,
		}

		// Set the cookie in the response (experimental) ===================================
		AccessToken, Error := jwt.GenerateAccessToken(&claims)
		//RefreshToken, Error1 := jwt.GenerateRefreshToken(&claims)
		//cookie := &fiber.Cookie{
		//	Name:     "jwt",
		//	Value:    Token,
		//	Expires:  time.Now().Add(24 * time.Hour),
		//	Path:     "/",
		//	SameSite: "Lax",
		//}
		//=====================================================
		if Error != nil {
			StatusCode = fiber.StatusConflict
			Responser = common.HTTPResponser(StatusCode, true, "access  token olusturma esnasinda porblem yasandi")
		} else {
			StatusCode = fiber.StatusOK
			c.Set(`Authorization`, "accesstoken : "+AccessToken)
			//c.Set("Refresh-Token", RefreshToken)
			//c.Cookie(cookie)
			//Responser = "kullanici girisi basarili"
			Responser = common.HTTPResponser(StatusCode, false, "Kullanici girisi basarili", Data)

		}
	}
	return c.Status(StatusCode).JSON(Responser)
}
