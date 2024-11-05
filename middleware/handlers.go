package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/anujmritunjay/go-postgres/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// type response struct {
// 	ID      int64  `json:"id,omitempty"`
// 	Message string `json:"message,omitempty"`
// }

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to Postgres Database")
	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var stock models.Stock
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	id, err := insertStock(stock)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "message": err.Error()})
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": id})

}

func insertStock(stock models.Stock) (int, error) {
	db := CreateConnection()
	defer db.Close()
	insertQuery := `INSERT INTO stocks(name, price, company) VALUES ($1, $2, $3) RETURNING stockid`
	var id int
	err := db.QueryRow(insertQuery, stock.Name, stock.Price, stock.Company).Scan(&id)
	return id, err
}
