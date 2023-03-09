package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"os"
)

var (
	DB  *pgxpool.Pool
	Ctx context.Context
)

func init() {
	err := godotenv.Load("./prod.env")
	if err != nil {
		fmt.Println(".env dosyasi yuklenirken bir hata ile karsilasildi")
	}
	host := os.Getenv("db_host")
	port := os.Getenv("db_port")
	user := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	ConnStr := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v", host, port, user, password, dbName)
	DB, err = pgxpool.Connect(context.Background(), ConnStr)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Println("")
		fmt.Println("database baglantisi basarili bir sekilde gerceklestirildi")
	}
	Ctx = context.Background()
}
