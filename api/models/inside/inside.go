package inside

import (
	postgre "MapProject/api/adapters/database"
	"MapProject/internal/entities"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

var query string

var errDb error

//func Insidemod(d entities.Inside, id int64) error { //kullanici id si ile veri donusu
//	//fmt.Println("id : ", id)
//	//query = `INSERT INTO polygon(latitude, longitude, name, groupname, index, groupid) VALUES ($1,$2,$3,$4,$5,$6)`
//	query = `INSERT INTO polygon(id,name,description,comment,user_id,coordinates) VALUES ($1,$2,$3,$4,$5,$6)`
//	_, err := postgre.DB.Exec(postgre.Ctx, query, d.Pinid, d.Polid, id)
//	if err != nil {
//		return fmt.Errorf("insert ornegi yapilamadi: %v", err)
//	}
//	return nil
//}

func Insidemod(d entities.Inside, id int64) error {
	query := fmt.Sprintf(`
	INSERT INTO inside (pinid, polid, user_id)
	select '%v', '%v', %v
	WHERE NOT EXISTS (
		SELECT 1
		FROM inside
		WHERE pinid = '%v' AND polid = '%v'
	)`, d.Pinid, d.Polid, id, d.Pinid, d.Polid)
	_, err := postgre.DB.Exec(postgre.Ctx, query)
	if err != nil {
		return fmt.Errorf("databasede bu degerler zaten kayitli: %v", err)
	}
	return nil
}

func GetInfoFromPoly(d entities.Inside, id int64) (interface{}, error) {
	polygonRow := postgre.DB.QueryRow(postgre.Ctx, "SELECT id, name, description, comment, coordinates FROM polygon WHERE id = $1 AND user_id = $2", d.Polid, id)
	var polygon entities.Polygon
	if err := polygonRow.Scan(&polygon.ID, &polygon.Name, &polygon.Description, &polygon.Comment, &polygon.Coordinates); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// Query the database to get the pin data
	pinRows, err := postgre.DB.Query(postgre.Ctx, "SELECT id, name, description, comment, coordinates FROM pin WHERE id IN (SELECT pinid FROM inside WHERE polid = $1) AND user_id = $2", d.Polid, id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer pinRows.Close()
	pinsayisi := 0
	var pins []entities.Pin
	for pinRows.Next() {
		var pin entities.Pin
		if err := pinRows.Scan(&pin.ID, &pin.Name, &pin.Description, &pin.Comment, &pin.Coordinates); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error()), nil
		}
		pinsayisi = pinsayisi + 1
		pins = append(pins, pin)
	}

	// Construct the response object
	response := map[string]interface{}{
		"id":          polygon.ID,
		"name":        polygon.Name,
		"description": polygon.Description,
		"comment":     polygon.Comment,
		"coordinates": polygon.Coordinates,
		"pinnumber":   pinsayisi,
		"insidepins":  pins,
	}
	return response, nil
}

func GetInfoFromPin(d entities.Inside, id int64) (interface{}, error) {
	polygonRow := postgre.DB.QueryRow(postgre.Ctx, "SELECT id, name, description, comment, coordinates FROM pin WHERE id = $1 AND user_id = $2", d.Pinid, id)
	var polygon entities.Pin
	if err := polygonRow.Scan(&polygon.ID, &polygon.Name, &polygon.Description, &polygon.Comment, &polygon.Coordinates); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// Query the database to get the pin data
	pinRows, err := postgre.DB.Query(postgre.Ctx, "SELECT id, name, description, comment, coordinates FROM polygon WHERE id IN (SELECT polid FROM inside WHERE pinid = $1) AND user_id = $2", d.Pinid, id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer pinRows.Close()

	var pols []entities.Polygon
	for pinRows.Next() {
		var pol entities.Polygon
		if err := pinRows.Scan(&pol.ID, &pol.Name, &pol.Description, &pol.Comment, &pol.Coordinates); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error()), nil
		}
		pols = append(pols, pol)
	}

	// Construct the response object
	response := map[string]interface{}{
		"id":             polygon.ID,
		"name":           polygon.Name,
		"description":    polygon.Description,
		"comment":        polygon.Comment,
		"coordinates":    polygon.Coordinates,
		"insidepolygons": pols,
	}
	return response, nil
}

//func GetAllInfo(d entities.Inside, id int64) (interface{}, error) {
//	// Query the database to get all polygons and their corresponding pins
//	polygonRows, err := postgre.DB.Query(postgre.Ctx, "SELECT p.id, p.name, p.description, p.comment, p.coordinates, pin.id, pin.name, pin.description, pin.comment, pin.coordinates FROM polygon p LEFT JOIN inside i ON p.id = i.polid LEFT JOIN pin ON i.pinid = pin.id WHERE p.user_id = $1", id)
//	if err != nil {
//		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
//	}
//	defer polygonRows.Close()
//
//	polygons := make(map[string]interface{})
//	for polygonRows.Next() {
//		var polygon entities.Polygon
//		var pin entities.Pin
//		var insidePinList []entities.Pin
//
//		if err := polygonRows.Scan(&polygon.ID, &polygon.Name, &polygon.Description, &polygon.Comment, &polygon.Coordinates, &pin.ID, &pin.Name, &pin.Description, &pin.Comment, &pin.Coordinates); err != nil {
//			return fiber.NewError(fiber.StatusInternalServerError, err.Error()), nil
//		}
//
//		if _, ok := polygons[polygon.ID]; !ok {
//			insidePinList = make([]entities.Pin, 0)
//		} else {
//			insidePinList = polygons[polygon.ID].([]entities.Pin)
//		}
//
//		//if pin.ID.Valid {
//		//	insidePinList = append(insidePinList, pin)
//		//}
//		if pin.ID != "" {
//			insidePinList = append(insidePinList, pin)
//		}
//
//		polygons[polygon.ID] = map[string]interface{}{
//			"id":          polygon.ID,
//			"name":        polygon.Name,
//			"description": polygon.Description,
//			"comment":     polygon.Comment,
//			"coordinates": polygon.Coordinates,
//			"insidepins":  insidePinList,
//		}
//	}
//
//	// Convert the map to a slice and return the response object
//	var response []interface{}
//	for _, value := range polygons {
//		response = append(response, value)
//	}
//	return response, nil
//}

func GetAllInfo(d entities.Inside, id int64) (interface{}, error) {
	// Query the database to get all polygons and their corresponding pins
	polygonRows, err := postgre.DB.Query(postgre.Ctx, "SELECT p.id, p.name, p.description, p.comment, p.coordinates, pin.id, pin.name, pin.description, pin.comment, pin.coordinates FROM polygon p LEFT JOIN inside i ON p.id = i.polid LEFT JOIN pin ON i.pinid = pin.id WHERE p.user_id = $1", id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer polygonRows.Close()

	polygons := make(map[string]map[string]interface{})
	for polygonRows.Next() {
		var polygon entities.Polygon
		var pin entities.Pin
		var insidePinList []entities.Pin

		if err := polygonRows.Scan(&polygon.ID, &polygon.Name, &polygon.Description, &polygon.Comment, &polygon.Coordinates, &pin.ID, &pin.Name, &pin.Description, &pin.Comment, &pin.Coordinates); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error()), nil
		}

		if _, ok := polygons[polygon.ID]; !ok {
			insidePinList = make([]entities.Pin, 0)
		} else {
			insidePinList = polygons[polygon.ID]["insidepins"].([]entities.Pin)
		}

		if pin.ID != "" {
			insidePinList = append(insidePinList, pin)
		}

		polygons[polygon.ID] = map[string]interface{}{
			"id":          polygon.ID,
			"name":        polygon.Name,
			"description": polygon.Description,
			"comment":     polygon.Comment,
			"coordinates": polygon.Coordinates,
			"insidepins":  insidePinList,
		}
	}

	// Convert the map to a slice and return the response object
	var response []interface{}
	for _, value := range polygons {
		insidePinList := value["insidepins"].([]entities.Pin)
		value["insidepins"] = insidePinList
		response = append(response, value)
	}
	return response, nil
}
