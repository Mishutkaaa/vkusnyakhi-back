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

func CreateNewProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newProduct models.NewProduct
		var categories string

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
			http.Error(w, "error decode json", http.StatusBadRequest)
			log.Println(err)
			return
		}

		if newProduct.Categories != nil {
			categories = "{" + strings.Join(*newProduct.Categories, ",") + "}"
		} else {
			categories = "{}"
		}

		query := fmt.Sprintf("insert into %s ", *newProduct.Table)
		_, err := db.Exec(query+"(name, image, categories, brand) values ($1, $2, $3, $4)",
			newProduct.Name, newProduct.Image, categories, newProduct.Brand)
		if err != nil {
			log.Println("cannot add new product", err)
			return
		}

	}
}
