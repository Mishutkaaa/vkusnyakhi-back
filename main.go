package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"vkusnyakhi-back/config"
	"vkusnyakhi-back/models"

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

	http.HandleFunc("/drinks", func(w http.ResponseWriter, r *http.Request) {
		var drinks []models.Drinks

		w.Header().Set("Content-Type", "application/json")
		rows, err := db.Query("select id, name, image, categories, brand from drinks")
		if err != nil {
			log.Println(err)
			http.Error(w, "cannot get drinks", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for rows.Next() {
			drink := models.Drinks{}
			if err := rows.Scan(&drink.ID, &drink.Name, &drink.Image, &drink.Category, &drink.Brand); err != nil {
				log.Println("scan err", err)
				return
			}
			drinks = append(drinks, drink)
		}
		json.NewEncoder(w).Encode(drinks)
	})
	addr := ":" + port
	http.ListenAndServe(addr, nil)
}
