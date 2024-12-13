package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"atmail"
	"atmail/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// start server
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT cannot be blank!")
	}

	// setup database
	databaseUrl := os.Getenv("MYSQL_URL")

	if databaseUrl == "" {
		log.Fatal("MYSQL_URL cannot be blank!")
	}

	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		log.Fatalf("failed to connect to the mysql database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("database ping failed: %v", err)
	}

	store := atmail.NewStore(db)

	server := server.New(store)

	fmt.Printf("Starting at port %s...\n", port)

	if err := http.ListenAndServe(":"+port, server); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
