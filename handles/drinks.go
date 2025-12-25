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

func GetDrinks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var drinks []models.Drinks
		var query string = "select id, name, image, categories, brand from drinks"
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
			if err := rows.Scan(&drink.ID, &drink.Name, &drink.Image, &drink.Categories, &drink.Brand); err != nil {
				log.Println("scan err", err)
				return
			}
			drinks = append(drinks, drink)
		}
		json.NewEncoder(w).Encode(drinks)
	}
}
