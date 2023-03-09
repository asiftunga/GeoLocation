package polygon

import (
	mod "MapProject/api/models/polygon"
	"MapProject/internal/entities"
	"github.com/gofiber/fiber/v2"
)

var (
	Responser  interface{}
	StatusCode int
)

// @Summary Create Polygon With User ID and Polygon
// @Description Create Polygon with expected request info
// @Tags Polygons
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT authorization token"
// @Param Polygon_info body entities.Polygon true "Request body"
// @Success 200 {object} entities.ResponseWrapper
// @Failure 502 {object} entities.ErrorResponseWrapper
// @Router /create/polygon [post]
func CreatePolygon(c *fiber.Ctx) error {
	id := c.Locals("id").(int64)
	dat := entities.Polygon{}
	// Read the data from the request body
	errParser := c.BodyParser(&dat)
	if errParser != nil {
		StatusCode = fiber.StatusBadGateway
		Responser = entities.ErrorResponseWrapper{
			StatusCode: StatusCode,
			Error:      true,
			Message:    errParser.Error(),
		}
	} else {
		Data := mod.CreatePolygon(dat, id)
		StatusCode = fiber.StatusOK
		Responser = entities.ResponseWrapper{
			Data:       Data,
			StatusCode: StatusCode,
			Error:      false,
			Message:    "islem başarılı",
		}
	}
	return c.Status(StatusCode).JSON(Responser)
}

// @Summary Get polygons by user ID
// @Description Get a list of polygons for the specified user ID
// @Tags Polygons
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT authorization token"
// @Success 200 {array} entities.Polygon2
// @Failure 400 {object} error
// @Router /get/polygons [get]
func GetPolygonsByUserID(c *fiber.Ctx) error {
	id := c.Locals("id").(int64)
	dat := entities.Polygon{}
	Data := mod.GetPolygonsByUserID(dat, id)
	StatusCode = fiber.StatusOK
	return c.Status(StatusCode).JSON(Data)
}

// @Summary Delete Polygon one coordinate With User ID and Polygon ID
// @Description Delete spesific Polygon with Polygon ID and User ID from header
// @Tags Polygons
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT authorization token"
// @Param body body entities.DeletePolygonOnePoint1 true "Request body"
// @Success 200 {object} entities.ResponseWrapper
// @Failure 400 {object} entities.ErrorResponseWrapper
// @Router /delete/onepoint [delete]
func DeleteOnePointByUserID(c *fiber.Ctx) error {
	id := c.Locals("id").(int64)
	dat := entities.DeletePolygonOnePoint{}
	// Read the data from the request body
	errParser := c.BodyParser(&dat)
	if errParser != nil {
		StatusCode = fiber.StatusBadGateway
		Responser = entities.ErrorResponseWrapper{
			StatusCode: StatusCode,
			Error:      true,
			Message:    errParser.Error(),
		}
	} else {
		Data := mod.DeletePolygonOnePoint(&dat, id)
		StatusCode = fiber.StatusOK
		Responser = entities.ResponseWrapper{
			Data:       Data,
			StatusCode: StatusCode,
			Error:      false,
			Message:    "poligon coordinat silme islemi başarılı",
		}
	}
	return c.Status(StatusCode).JSON(Responser)
}

// @Summary Delete Polygon With User ID and Polygon ID
// @Description Delete spesific Polygon with Polygon ID and User ID from header
// @Tags Polygons
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT authorization token"
// @Param Polygon_Id body entities.DeletePolygonOnePoint true "Request body"
// @Success 200 {object} entities.ResponseWrapper
// @Failure 400 {object} entities.ErrorResponseWrapper
// @Router /delete/polygon [delete]
func DeletePolygon(c *fiber.Ctx) error {
	id := c.Locals("id").(int64)
	dat := entities.DeletePolygonOnePoint{}
	// Read the data from the request body
	errParser := c.BodyParser(&dat)
	if errParser != nil {
		StatusCode = fiber.StatusBadGateway
		//Responser = common.HTTPResponser(nil, StatusCode, true, errParser.Error())
		Responser = entities.ErrorResponseWrapper{
			StatusCode: StatusCode,
			Error:      true,
			Message:    errParser.Error(),
		}
	} else {
		Data := mod.DeletePolygon(&dat, id)
		StatusCode = fiber.StatusOK
		Responser = entities.ResponseWrapper{
			Data:       Data,
			StatusCode: StatusCode,
			Error:      false,
			Message:    "poligon silme islemi basarili",
		}
	}
	return c.Status(StatusCode).JSON(Responser)
}

// @Summary Update Polygon With User ID and Polygon
// @Description Update spesific Polygon with Polygon infos and User ID from header
// @Tags Polygons
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT authorization token"
// @Param Polygon_info body entities.Polygon true "Request body"
// @Success 200 {object} entities.ResponseWrapper
// @Failure 502 {object} entities.ErrorResponseWrapper
// @Router /update/polygon [put]
func UpdatePolygon(c *fiber.Ctx) error {
	id := c.Locals("id").(int64)
	dat := entities.Polygon{}
	err := c.BodyParser(&dat)
	if err != nil {
		StatusCode = fiber.StatusBadGateway
		Responser = entities.ErrorResponseWrapper{
			StatusCode: StatusCode,
			Error:      true,
			Message:    err.Error(),
		}
	} else {
		Data := mod.UpdatePolygon(&dat, id)
		StatusCode = fiber.StatusOK
		Responser = entities.ResponseWrapper{
			Data:       Data,
			StatusCode: StatusCode,
			Error:      false,
			Message:    "poligon  guncelleme islemi başarılı",
		}
	}
	return c.Status(StatusCode).JSON(Responser)
}
