package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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
		var query string = "select id, name, image from drinks"
		var where []string

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		categories := r.URL.Query().Get("categories")
		brand := r.URL.Query().Get("brand")

		if categories != "" {
			where = append(where, fmt.Sprintf(" categories @> array [%s]", categories))
		}
		if brand != "" {
			where = append(where, fmt.Sprintf("brand = %s", brand))
		}
		if len(where) != 0 {
			query += " where " + strings.Join(where, "and")
		}
		rows, err := db.Query(query)
		if err != nil {
			log.Println(err)
			http.Error(w, "cannot get drinks", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for rows.Next() {
			drink := models.Drinks{}
			if err := rows.Scan(&drink.ID, &drink.Name, &drink.Image); err != nil {
				log.Println("scan err", err)
				return
			}
			drinks = append(drinks, drink)
		}
		json.NewEncoder(w).Encode(drinks)
	})

	http.HandleFunc("/food", func(w http.ResponseWriter, r *http.Request) {
		var foods []models.Food
		var query string = "select id, name, image from food"
		var where []string

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		categories := r.URL.Query().Get("categories")
		brand := r.URL.Query().Get("brand")

		if categories != "" {
			where = append(where, fmt.Sprintf(" categories @> array [%s]", categories))
		}
		if brand != "" {
			where = append(where, fmt.Sprintf("brand = %s", brand))
		}
		if len(where) != 0 {
			query += " where " + strings.Join(where, "and")
		}

		rows, err := db.Query(query)
		if err != nil {
			log.Println(err)
			http.Error(w, "cannot get food", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for rows.Next() {
			food := models.Food{}
			if err := rows.Scan(&food.ID, &food.Name, &food.Image); err != nil {
				log.Println("scan err", err)
				return
			}
			foods = append(foods, food)
		}
		json.NewEncoder(w).Encode(foods)
	})

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		var categories []models.Categories

		table := r.URL.Query().Get("type")

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		rows, err := db.Query("select id, name from categories where type =$1 or type ='any'", table)
		if err != nil {
			log.Println(err)
			http.Error(w, "cannot get categories", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for rows.Next() {
			category := models.Categories{}
			if err := rows.Scan(&category.ID, &category.Name); err != nil {
				log.Println("scan err", err)
				return
			}
			categories = append(categories, category)
		}
		json.NewEncoder(w).Encode(categories)
	})

	http.HandleFunc("/brand", func(w http.ResponseWriter, r *http.Request) {
		var brands []models.Brand

		table := r.URL.Query().Get("type")

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		rows, err := db.Query("select id, name from brand where type = $1 or type = 'any'", table)
		if err != nil {
			log.Println(err)
			http.Error(w, "cannot get brand", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for rows.Next() {
			brand := models.Brand{}
			if err := rows.Scan(&brand.ID, &brand.Name); err != nil {
				log.Println("scan err", err)
				return
			}
			brands = append(brands, brand)
		}
		json.NewEncoder(w).Encode(brands)
	})

	addr := ":" + port
	http.ListenAndServe(addr, nil)
}
