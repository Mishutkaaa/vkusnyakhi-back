package handles

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"vkusnyakhi-back/models"
)

func GetFood(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var foods []models.Food
		var query string = "select id, name, image, categories, brand from food"
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
			where = append(where, fmt.Sprintf("brand in (%s)", brand))
		}
		if len(where) != 0 {
			query += " where " + strings.Join(where, " and ")
		}
		log.Println(query)
		rows, err := db.Query(query)
		if err != nil {
			log.Println(err)
			http.Error(w, "cannot get food", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for rows.Next() {
			food := models.Food{}
			if err := rows.Scan(&food.ID, &food.Name, &food.Image, &food.Categories, &food.Brand); err != nil {
				log.Println("scan err", err)
				return
			}
			foods = append(foods, food)
		}
		json.NewEncoder(w).Encode(foods)
	}
}
func EditFood(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var food models.Food

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(&food); err != nil {
			http.Error(w, "error decode json", http.StatusBadRequest)
			log.Println(err)
			return
		}

		_, err := db.Exec("update food set name = $1, image = $2, categories = $3, brand = $4 where id = $5",
			food.Name, food.Image, food.Categories, food.Brand, food.ID)
		if err != nil {
			log.Println("cannot edit food", err)
			return
		}
		log.Println("всё ок")
	}
}
