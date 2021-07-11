package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"go_api/models"

	"github.com/joho/godotenv" // package used to read the .env file
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// create connection with postgres db
func createConnection() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// GetAllArticles will return all the articles
func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get all the users in the db
	articles, err := getAllArticles()

	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}

	// send all the users as response
	json.NewEncoder(w).Encode(articles)
}

func getAllArticles() ([]models.Article, error) {
	// create the postgres db connection
	db := createConnection()
	// close the db connection
	defer db.Close()

	var articles []models.Article

	// create the select sql query
	sqlStatement := `SELECT * FROM articles`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the queries", err)
	}

	defer rows.Close()

	for rows.Next() {
		var article models.Article
		// unmarshal the row object to user
		err = rows.Scan(&article.ID, &article.Name)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the article in the articles slice
		articles = append(articles, article)
	}

	return articles, err
}
