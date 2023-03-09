package polygon

import (
	postgre "MapProject/api/adapters/database"
	"MapProject/internal/entities"
	"fmt"
)

var query string

var errDb error

func CreatePolygon(d entities.Polygon, id int64) error { //kullanici id si ile veri donusu
	//fmt.Println("id : ", id)
	//query = `INSERT INTO polygon(latitude, longitude, name, groupname, index, groupid) VALUES ($1,$2,$3,$4,$5,$6)`
	query = `INSERT INTO polygon(id,name,description,comment,user_id,coordinates) VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := postgre.DB.Exec(postgre.Ctx, query, d.ID, d.Name, d.Description, d.Comment, id, d.Coordinates)
	if err != nil {
		return fmt.Errorf("insert ornegi yapilamadi: %v", err)
	}
	return nil
}

//func GetPolygonsById(d entities.Polygon2, id int64) error {
//	var polygons []entities.Polygon2
//	rows, err := postgre.DB.Query(postgre.Ctx, "SELECT id,name,description,comment,coordinates FROM actor")
//	if err != nil {
//		println(err.Error())
//	}
//	defer rows.Close()
//
//}

func GetPolygonsByUserID(d entities.Polygon, id int64) []entities.Polygon2 {
	var polys []entities.Polygon
	var polygon2s []entities.Polygon2
	rows, err := postgre.DB.Query(postgre.Ctx, "SELECT id,name,description,comment,coordinates FROM polygon WHERE user_id = $1", id)
	if err != nil {
		println(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var polygon entities.Polygon
		if err := rows.Scan(&polygon.ID, &polygon.Name, &polygon.Description, &polygon.Comment, &polygon.Coordinates); err != nil {
			println(errDb.Error())
		}
		polys = append(polys, polygon)
	}
	for _, polygon := range polys {
		var coords [][2]float64
		for _, coord := range polygon.Coordinates {
			coords = append(coords, [2]float64{coord.Long, coord.Lat})
		}
		polygon2s = append(polygon2s, entities.Polygon2{
			ID:          polygon.ID,
			Name:        polygon.Name,
			Description: polygon.Description,
			Comment:     polygon.Comment,
			UserID:      polygon.UserID,
			Coordinates: coords,
		})
	}
	return polygon2s
}

//func ConvertPolygons(polygons []entities.Polygon) []entities.Polygon2 {
//	var polygon2s []entities.Polygon2
//	for _, polygon := range polygons {
//		var coords [][2]float64
//		for _, coord := range polygon.Coordinates {
//			coords = append(coords, [2]float64{coord.Long, coord.Lat})
//		}
//		polygon2s = append(polygon2s, entities.Polygon2{
//			ID:          polygon.ID,
//			Name:        polygon.Name,
//			Description: polygon.Description,
//			Comment:     polygon.Comment,
//			UserID:      polygon.UserID,
//			Coordinates: coords,
//		})
//	}
//	return polygon2s
//}

func DeletePolygonOnePoint(d *entities.DeletePolygonOnePoint, id int64) error {
	quee := fmt.Sprintf(`UPDATE polygon
SET coordinates = (SELECT jsonb_agg(e) FROM (SELECT jsonb_array_elements(coordinates) AS e) subq WHERE e != '{"lng": %f, "lat": %f}')::jsonb
WHERE id = '%v' AND user_id = %v AND coordinates @> '[{"lng": %f, "lat": %f}]';`, d.Coordinates[0].Long, d.Coordinates[0].Lat, d.ID, id, d.Coordinates[0].Long, d.Coordinates[0].Lat)
	//fmt.Println(quee)
	_, err := postgre.DB.Exec(postgre.Ctx, quee)
	if err != nil {
		return fmt.Errorf("Silme islemi gerceklestirilemedi: %v", err)
	}
	return nil
}

func DeletePolygon(d *entities.DeletePolygonOnePoint, id int64) error {
	quee := fmt.Sprintf(`DELETE FROM polygon WHERE id = '%v'  AND user_id = %v ;`, d.ID, id)
	_, err := postgre.DB.Exec(postgre.Ctx, quee)
	if err != nil {
		return fmt.Errorf("Silme islemi gerceklestirilemedi: %v", err)
	}
	return nil
}

func UpdatePolygon(d *entities.Polygon, id int64) error {
	quee := `UPDATE polygon SET name = $1, description = $2, comment = $3, coordinates = $4 WHERE id = $5 AND user_id = $6;`
	_, err := postgre.DB.Exec(postgre.Ctx, quee, d.Name, d.Description, d.Comment, d.Coordinates, d.ID, id)
	//fmt.Println(quee)
	if err != nil {
		return fmt.Errorf("insert ornegi yapilamadi: %v", err)
	}
	return nil
}
