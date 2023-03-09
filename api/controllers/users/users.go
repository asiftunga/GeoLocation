package users

import (
	mod "MapProject/api/models/users"
	"MapProject/internal/common"
	"MapProject/internal/entities"
	"github.com/gofiber/fiber/v2"
)

var (
	Responser  interface{}
	StatusCode int
)

func GetUserList(c *fiber.Ctx) error {
	dat := entities.User{}
	errParser := c.BodyParser(&dat) //note "Body Parser" is a function that helps in extracting data from the body of an HTTP request and storing it in a struct, its like bodyparser takes http body and put into your struct (which you gave to the bodyparser method as paramater). //kisaca parametre olarak koydugum struct icerisine (pointer seklinde verdigimden dolayi) http govdesini ayristirir ve bilgileri koyar
	if errParser != nil {
		StatusCode = fiber.StatusBadGateway
		Responser = common.HTTPResponser(StatusCode, true, errParser.Error())
	} else {
		Data := mod.GetAllUser(dat)
		StatusCode = fiber.StatusOK
		Responser = common.HTTPResponser(StatusCode, false, "İşlem başarılı", Data)
	}
	return c.Status(StatusCode).JSON(Responser)
}

func GetUserById(c *fiber.Ctx) error {
	dat := entities.User{}
	errParser := c.BodyParser(&dat)
	if errParser != nil {
		StatusCode = fiber.StatusBadGateway
		Responser = common.HTTPResponser(StatusCode, true, errParser.Error())
	} else {
		Data := mod.GetUserById(dat)
		StatusCode = fiber.StatusOK
		Responser = Data
	}
	return c.Status(StatusCode).JSON(Responser)
}
