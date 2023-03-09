package auth

import (
	postgre "MapProject/api/adapters/database"
	"MapProject/internal/common"
	ent "MapProject/internal/entities"
)

func Login(d ent.Login) (Record int8, auth ent.Auth) {
	//=====
	runes := []rune(d.UserPass)
	b := make([]byte, len(runes))
	for i, r := range runes {
		b[i] = byte(r)
	}
	//for converting string to byte slice
	//=======
	query := `SELECT id,name,surname FROM users WHERE email=$1 AND password=$2`
	//b == d.UserPass boyle kisa bir isimlendirme verdim byte slice donusturulurken
	errDb := postgre.DB.QueryRow(postgre.Ctx, query, d.Email, b).Scan(&auth.AuthId, &auth.Name, &auth.Surname) //degisken kisimlari queryrow icinde, select kisimlari ise scan icine ve selectten donen degeri o structa koyuyor
	/*
		The postgre.DB.QueryRow method is then used to execute the query. It takes three arguments:
		postgre.Ctx: a context object that contains information about the context in which the query is being executed.
		query: the SQL query to be executed.
		d.Email and d.UserPass: the values to be used in the query to filter the results.
		The result of the query is then stored in the auth.AuthId variable using the Scan method. The Scan method takes the first row of the result set and stores it in the specified variable. In this case, the auth.AuthId variable is being used to store the id of the user.
		If an error occurs during the execution of the query, it will be stored in the errDb variable.
		In simpler terms, this code is fetching the id of a user from the database, based on their email and password, and storing it in the auth.AuthId variable.
	*/
	if errDb != nil && common.Contains(errDb.Error(), "no rows in result set") == false {
		println(errDb.Error())
		Record = -1 //this record for lets say i dont have any row in db, so this record works like a flag and tell me did the sent json information match with the database or not?
	} else if errDb != nil && common.Contains(errDb.Error(), "no rows in result set") == true {
		Record = 0
	} else {
		Record = 1
	}
	return
}
