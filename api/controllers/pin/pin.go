package pin

import (
	mod "MapProject/api/models/pin"
	"MapProject/internal/common"
	"MapProject/internal/entities"
	"github.com/gofiber/fiber/v2"
)

var (
	Responser  interface{}
	StatusCode int
)

// @Summary Create Pin With User ID and Pin information
// @Description Create Pins with expected request info
// @Tags Pins
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT authorization token"
// @Param Pin_info body entities.Pin true "Request body"
// @Success 200 {object} entities.ResponseWrapper
// @Failure 502 {object} entities.ErrorResponseWrapper
// @Router /create/pin [post]
func CreatePin(c *fiber.Ctx) error {
	id := c.Locals("id").(int64)
	dat := entities.Pin{}
	// Read the data from the request body
	errParser := c.BodyParser(&dat)
	if errParser != nil {
		StatusCode = fiber.StatusBadGateway
		Responser = common.HTTPResponser(StatusCode, true, errParser.Error())
	}
	Data := mod.CreatePin(dat, id)
	StatusCode = fiber.StatusOK
	Responser = common.HTTPResponser(StatusCode, false, "İşlem başarılı", Data)
	//Responser = Data
	return c.Status(StatusCode).JSON(Responser)
}

// @Summary Get array of pins by user ID
// @Description Get a list of pins for the specified user ID
// @Tags Pins
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT authorization token"
// @Success 200 {array} entities.Pin2
// @Failure 400 {object} error
// @Router /get/pins [get]
func GetPinsByUserID(c *fiber.Ctx) error {
	id := c.Locals("id").(int64)
	dat := entities.Pin{}
	Data := mod.GetPinsByUserID(dat, id)
	StatusCode = fiber.StatusOK
	//Responser = common.HTTPResponser(Data, StatusCode, false, "İşlem başarılı")
	return c.Status(StatusCode).JSON(Data)
}

// @Summary Delete Pin With User ID and Pin ID
// @Description Delete spesific pin with pin ID and User ID from header
// @Tags Pins
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT authorization token"
// @Param Pin_Id body entities.DeletePolygonOnePoint true "Request body"
// @Success 200 {object} entities.ResponseWrapper
// @Failure 400 {object} entities.ErrorResponseWrapper
// @Router /delete/pin [delete]
func DeletePin(c *fiber.Ctx) error {
	id := c.Locals("id").(int64)
	dat := entities.Pin{}
	// Read the data from the request body
	errParser := c.BodyParser(&dat)
	if errParser != nil {
		StatusCode = fiber.StatusBadGateway
		Responser = common.HTTPResponser(StatusCode, true, errParser.Error())
	}
	Data := mod.DeletePin(&dat, id) //delete with only id
	StatusCode = fiber.StatusOK
	Responser = common.HTTPResponser(StatusCode, false, "pin  silme islemi başarılı", Data)
	//Responser = Data
	return c.Status(StatusCode).JSON(Responser)
}

// @Summary Delete All Pins With User ID and Pin ID
// @Description Delete all pins with pin ID and User ID from header
// @Tags Pins
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT authorization token"
// @Success 200 {object} entities.ResponseWrapper
// @Failure 400 {object} entities.ErrorResponseWrapper
// @Router /delete/deleteallpins [delete]
func DeleteAllPin(c *fiber.Ctx) error {
	id := c.Locals("id").(int64)
	// Read the data from the request body
	Data := mod.DeleteAllPin(id)
	StatusCode = fiber.StatusOK
	Responser = common.HTTPResponser(StatusCode, false, "pin  silme islemi başarılı", Data)
	//Responser = Data
	return c.Status(StatusCode).JSON(Responser)
}

// @Summary Update Pin With User ID and Polygon
// @Description Update spesific Pin with Pin infos and User ID from header
// @Tags Pins
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT authorization token"
// @Param Pin_info body entities.Pin true "Request body"
// @Success 200 {object} entities.ResponseWrapper
// @Failure 502 {object} entities.ErrorResponseWrapper
// @Router /update/pin [put]
func UpdatePin(c *fiber.Ctx) error {
	id := c.Locals("id").(int64)
	dat := entities.Pin{}
	errParser := c.BodyParser(&dat)
	if errParser != nil {
		StatusCode = fiber.StatusBadGateway
		Responser = common.HTTPResponser(StatusCode, true, errParser.Error())
	}
	Data := mod.UpdatePin(&dat, id)
	StatusCode = fiber.StatusOK
	Responser = common.HTTPResponser(StatusCode, false, "Pin guncelleme islemi basarili", Data)
	return c.Status(StatusCode).JSON(Responser)
}
