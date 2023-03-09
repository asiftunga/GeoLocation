package users

import (
	postgre "MapProject/api/adapters/database"
	"MapProject/internal/entities"
)

var (
	errDb error
)

func GetAllUser(d entities.User) []entities.User {
	//query = `SELECT row_to_json(X) FROM (SELECT firstname, lastname, gender, email FROM person WHERE id=$1) AS X`
	var users []entities.User
	//query = `SELECT first_name, last_name FROM actor`
	//errDb = postgre.DB.QueryRow(postgre.Ctx, query).Scan(&d.Name)
	rows, err := postgre.DB.Query(postgre.Ctx, "SELECT first_name, last_name FROM actor")
	if err != nil {
		println(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&d.Name, &d.Surname); err != nil {
			return nil
		}
		users = append(users, d)
	}
	return users
}

func GetUserById(d entities.User) (returndata entities.User) { //kullanici id si ile veri donusu
	query := `SELECT first_name FROM actor WHERE actor_id = $1`
	errDb = postgre.DB.QueryRow(postgre.Ctx, query, d.ID).Scan(&returndata.Name)
	if errDb != nil {
		println(errDb.Error())
	}
	return
}
