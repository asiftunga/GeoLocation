package pin

import (
	postgre "MapProject/api/adapters/database"
	"MapProject/internal/entities"
	"fmt"
)

var query string

var errDb error

func CreatePin(d entities.Pin, id int64) error { //kullanici id si ile veri donusu

	//query = `INSERT INTO pin(id,name,description,comment,user_id,coordinates) VALUES ($1,$2,$3,$4,$5,$6)`
	query = `
INSERT INTO pin (id, name, description, comment, user_id, coordinates) 
SELECT $1, $2, $3, $4, $5, $6 
WHERE NOT EXISTS (
  SELECT 1
  FROM pin
  WHERE coordinates = $6
);
`
	_, err := postgre.DB.Exec(postgre.Ctx, query, d.ID, d.Name, d.Description, d.Comment, id, d.Coordinates)
	//_, err := postgre.DB.Exec(postgre.Ctx, query)
	if err != nil {
		return fmt.Errorf("insert ornegi yapilamadi: %v", err)
	}
	return nil
}

//	func GetPinsByUserID(d entities.Pin, id int64) []entities.Pin2 {
//		var pins2 []entities.Pin2
//		var pins []entities.Pin
//		rows, err := postgre.DB.Query(postgre.Ctx, "SELECT id,name,description,comment,coordinates FROM pin WHERE user_id = $1", id)
//		if err != nil {
//			println(err.Error())
//		}
//		defer rows.Close()
//		for rows.Next() {
//			if err := rows.Scan(&d.ID, &d.Name, &d.Description, &d.Comment, &d.Coordinates); err != nil {
//				println(errDb.Error())
//			}
//			pins = append(pins, d)
//		}
//		for _, pin := range pins {
//			fmt.Println(pin)
//			var coords [][2]float64
//			for _, coord := range pin.Coordinates {
//				coords = append(coords, [2]float64{coord.Long, coord.Lat})
//			}
//			pins2 = append(pins2, entities.Pin2{
//				ID:          pin.ID,
//				Name:        pin.Name,
//				Description: pin.Description,
//				Comment:     pin.Comment,
//				UserID:      pin.UserID,
//				Coordinates: coords,
//			})
//		}
//		return pins2
//	}
func GetPinsByUserID(d entities.Pin, id int64) []entities.Pin2 {
	var pins2 []entities.Pin2
	var pins []entities.Pin
	rows, err := postgre.DB.Query(postgre.Ctx, "SELECT id,name,description,comment,coordinates FROM pin WHERE user_id = $1", id)
	if err != nil {
		println(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var pin entities.Pin
		if err := rows.Scan(&pin.ID, &pin.Name, &pin.Description, &pin.Comment, &pin.Coordinates); err != nil {
			println(errDb.Error())
		}
		pins = append(pins, pin)
	}
	for _, pin := range pins {
		var coords [][2]float64
		for _, coord := range pin.Coordinates {
			coords = append(coords, [2]float64{coord.Long, coord.Lat})
		}
		pins2 = append(pins2, entities.Pin2{
			ID:          pin.ID,
			Name:        pin.Name,
			Description: pin.Description,
			Comment:     pin.Comment,
			UserID:      pin.UserID,
			Coordinates: coords,
		})
	}
	return pins2
}

func DeletePin(d *entities.Pin, id int64) error {
	quee := fmt.Sprintf(`DELETE FROM pin WHERE id = '%v' AND user_id = %v ;`, d.ID, id)
	_, err := postgre.DB.Exec(postgre.Ctx, quee)
	if err != nil {
		return fmt.Errorf("Silme islemi gerceklestirilemedi: %v", err)
	}
	return nil
}

func DeleteAllPin(id int64) error {
	quee := fmt.Sprintf(`DELETE FROM pin WHERE user_id = %v ;`, id)
	_, err := postgre.DB.Exec(postgre.Ctx, quee)
	if err != nil {
		return fmt.Errorf("Silme islemi gerceklestirilemedi: %v", err)
	}
	return nil
}

func UpdatePin(d *entities.Pin, id int64) error {
	quee := `UPDATE pin SET name = $1, description = $2, comment = $3, coordinates = $4 WHERE id = $5 AND user_id = $6;`
	_, err := postgre.DB.Exec(postgre.Ctx, quee, d.Name, d.Description, d.Comment, d.Coordinates, d.ID, id)
	//fmt.Println(quee)
	if err != nil {
		return fmt.Errorf("insert ornegi yapilamadi: %v", err)
	}
	return nil
}
