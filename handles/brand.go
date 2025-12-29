package handles

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"vkusnyakhi-back/models"
)

func GetBrand(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}

func CreateBrand(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var brand models.Brand

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if err := json.NewDecoder(r.Body).Decode(&brand); err != nil {
			http.Error(w, "error decode json", http.StatusBadRequest)
			log.Println(err)
			return
		}

		_, err := db.Exec("insert into brand (name, type) values ($1, $2)",
			brand.Name, brand.Type)
		if err != nil {
			log.Println("cannot create brand", err)
			return
		}
	}
}
