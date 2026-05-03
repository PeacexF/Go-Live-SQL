package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	connStr := "postgres://postgres:complex_and_hard_password@localhost:5432/postgres?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/api/query", handleUniversalQuery)

	fmt.Println("Live DB Monitor running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleUniversalQuery(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SQL string `json:"sql"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	rows, err := db.Query(req.SQL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer rows.Close()

	cols, _ := rows.Columns()
	results := make([]map[string]interface{}, 0)

	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		results = append(results, m)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
