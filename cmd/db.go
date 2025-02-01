package cmd

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func InitDB() {
	// Fix the connection string
	connPool, err := pgxpool.NewWithConfig(context.Background(), DBConfig(""))
	if err != nil {
		log.Fatal("Error while creating connection to the database!!")
	}

	// Acquire a connection from the pool
	connection, err := connPool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error while acquiring connection from database pool!")
	}
	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		log.Fatal("Could not PING database")
	}
	fmt.Println("Connected to the database")
}
