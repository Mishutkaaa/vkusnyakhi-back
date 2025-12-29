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
	http.HandleFunc("/editDrinks", handles.EditDrinks(db))
	http.HandleFunc("/drinks/{id}", handles.DeleteDrinks(db))

	http.HandleFunc("/food", handles.GetFood(db))
	http.HandleFunc("/editFood", handles.EditFood(db))
	http.HandleFunc("/food/{id}", handles.DeleteFood(db))

	http.HandleFunc("/categories", handles.GetCategories(db))
	http.HandleFunc("/createCategories", handles.CreateCategories(db))
	http.HandleFunc("/editCategories", handles.EditCategories(db))
	http.HandleFunc("/categories/{id}", handles.DeleteCategories(db))

	http.HandleFunc("/brand", handles.GetBrand(db))
	http.HandleFunc("/createBrand", handles.CreateBrand(db))
	http.HandleFunc("/editBrand", handles.EditBrand(db))
	http.HandleFunc("/brand/{id}", handles.DeleteBrand(db))

	http.HandleFunc("/newProduct", handles.CreateNewProduct(db))

	addr := ":" + port
	http.ListenAndServe(addr, nil)
}
