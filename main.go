package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"vkusnyakhi-back/config"
	"vkusnyakhi-back/handles"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env not found")
	}
	conf := config.NewConfigFromEnv()
	conn := conf.Conf()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Println("err conn db", err)
		return
	}
	defer db.Close()

	http.HandleFunc("/drinks", handles.GetDrinks(db))
	http.HandleFunc("/food", handles.GetFood(db))
	http.HandleFunc("/categories", handles.GetCategories(db))
	http.HandleFunc("/brand", handles.GetBrand(db))
	http.HandleFunc("/newProduct", handles.CreateNewProduct(db))
	http.HandleFunc("/editFood", handles.EditFood(db))

	addr := ":" + port
	http.ListenAndServe(addr, nil)
}
