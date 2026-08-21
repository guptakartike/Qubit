package main

import (
	"fmt"

	"os"
	"strconv"

	"github.com/guptakartike/qubit/internal/database"
	"github.com/guptakartike/qubit/internal/server"
	"github.com/joho/godotenv"
)


func main(){

	err:=godotenv.Load()
	if err != nil{
		fmt.Println("Warning: .env file not found")
	}
	port:=os.Getenv("PORT")
	if port == ""{
		port="8080"
	}

	portInt,err := strconv.Atoi(port)
	if err!=nil{
		panic("invalid port: ")
	}

	srv:=server.New(portInt)
	
	fmt.Println("Qubit api running on http://localhost:"+port)


	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		panic("DATABASE_URL environment variable is not set")
	}

	
	pgxPool,err:= database.NewPool(databaseURL)
	if err!=nil{
		panic("Pool connection failed: "+err.Error())
	}
	defer pgxPool.Close()
	fmt.Println("Database coonection Succesful")
	

	err = srv.ListenAndServe()
	if err!=nil{
		panic(err)
	}

	


}
