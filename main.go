package main

import (
	"fmt"
	pg "github.com/predixdeveloperACN/postgres-template/postgres"
	sr "github.com/predixdeveloperACN/postgres-template/server"
)

func main() {

	fmt.Println("---------------------------------------")
	fmt.Println("Starting service...")
	fmt.Println("---------------------------------------")

	//init DB
	pg.Open_postgres("user=postgres password=predix dbname=postgres sslmode=disable")
	defer pg.Database.Close()

	//  start rest interface here
	sr.SetupRestServer()
	sr.StartRestServer()
}