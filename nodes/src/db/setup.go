package db

import "os"

func Setup() {
	db := Connect()
	input, _ := os.ReadFile("./db/init.sql")
	db.Exec(string(input))
}
