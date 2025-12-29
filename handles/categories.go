package handles

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"vkusnyakhi-back/models"
)

func GetCategories(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}

}

func CreateCategories(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var categories models.Categories

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if err := json.NewDecoder(r.Body).Decode(&categories); err != nil {
			http.Error(w, "error decode json", http.StatusBadRequest)
			log.Println(err)
			return
		}

		_, err := db.Exec("insert into categories (name, type) values ($1, $2)",
			categories.Name, categories.Table)
		if err != nil {
			log.Println("cannot edit food", err)
			return
		}
	}
}

