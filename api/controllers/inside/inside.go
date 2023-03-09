package inside

import (
	mod "MapProject/api/models/inside"
	"MapProject/internal/common"
	"MapProject/internal/entities"
	"github.com/gofiber/fiber/v2"
)

var (
	Responser  interface{}
	StatusCode int
)

// @Summary Insert inside table  With Polygon ID and Pin ID with User ID
// @Description Insert inside table  With Polygon ID and Pin ID with User ID from request headers
// @Tags InsideInfo
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT authorization token"
// @Param PolAndPinInfo body entities.Inside true "Request body"
// @Success 200 {object} entities.ResponseWrapper
// @Failure 502 {object} entities.ErrorResponseWrapper
// @Router /create/inside [post]
func Inside(c *fiber.Ctx) error {
	id := c.Locals("id").(int64)
	dat := entities.Inside{}
	// Read the data from the request body
	errParser := c.BodyParser(&dat)
	if errParser != nil {
		StatusCode = fiber.StatusBadGateway
		Responser = common.HTTPResponser(StatusCode, true, errParser.Error())
	}
	Data := mod.Insidemod(dat, id)
	StatusCode = fiber.StatusOK
	Responser = common.HTTPResponser(StatusCode, false, "İşlem başarılı", Data)
	//Responser = Data
	return c.Status(StatusCode).JSON(Responser)
}

// @Summary get pin information that inside the polygon
// @Description get pin information contained within a particular polygon
// @Tags InsideInfo
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT authorization token"
// @Param Polygon_Id body entities.Inside2 true "Request body"
// @Success 200 {object} entities.ResponseWrapper
// @Failure 502 {object} entities.ErrorResponseWrapper
// @Router /create/getinfofrompolid [post]
func GetInfoFromPolygon(c *fiber.Ctx) error {
	id := c.Locals("id").(int64)
	dat := entities.Inside{}
	errParser := c.BodyParser(&dat)
	if errParser != nil {
		StatusCode = fiber.StatusBadGateway
		Responser = common.HTTPResponser(StatusCode, true, errParser.Error())
	}
	Data, err := mod.GetInfoFromPoly(dat, id)
	if err != nil {
		StatusCode = fiber.StatusBadGateway
		Responser = common.HTTPResponser(StatusCode, true, errParser.Error())
	}
	StatusCode = fiber.StatusOK
	Responser = common.HTTPResponser(StatusCode, false, "İşlem başarılı", Data)
	//Responser = Data
	return c.Status(StatusCode).JSON(Responser)
}

// @Summary get polygons information that the pin is in
// @Description according to the pin information, the information of the polygons in which the pin is located inside of polygons
// @Tags InsideInfo
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT authorization token"
// @Param Pin_Id body entities.Inside1 true "Request body"
// @Success 200 {object} entities.ResponseWrapper
// @Failure 502 {object} entities.ErrorResponseWrapper
// @Router /create/getinfofrompinid [post]
func GetInfoFromPin(c *fiber.Ctx) error {
	id := c.Locals("id").(int64)
	dat := entities.Inside{}
	errParser := c.BodyParser(&dat)
	if errParser != nil {
		StatusCode = fiber.StatusBadGateway
		Responser = common.HTTPResponser(StatusCode, true, errParser.Error())
	}
	Data, err := mod.GetInfoFromPin(dat, id)
	if err != nil {
		StatusCode = fiber.StatusBadGateway
		Responser = common.HTTPResponser(StatusCode, true, errParser.Error())
	}
	StatusCode = fiber.StatusOK
	Responser = common.HTTPResponser(StatusCode, false, "İşlem başarılı", Data)
	//Responser = Data
	return c.Status(StatusCode).JSON(Responser)
}

//func GetAllInfo(c *fiber.Ctx) error {
//	id := c.Locals("id").(int64)
//	dat := entities.Inside{}
//	errParser := c.BodyParser(&dat)
//	if errParser != nil {
//		StatusCode = fiber.StatusBadGateway
//		Responser = common.HTTPResponser(nil, StatusCode, true, errParser.Error())
//	}
//	Data, err := mod.GetAllInfo(dat, id)
//	if err != nil {
//		StatusCode = fiber.StatusBadGateway
//		Responser = common.HTTPResponser(nil, StatusCode, true, errParser.Error())
//	}
//	StatusCode = fiber.StatusOK
//	Responser = common.HTTPResponser(Data, StatusCode, false, "İşlem başarılı")
//	//Responser = Data
//	return c.Status(StatusCode).JSON(Responser)
//}
