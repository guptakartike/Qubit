package main

import (

	"fmt"

	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/guptakartike/qubit/cmd/internal/server"
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

	err = srv.ListenAndServe()
	if err!=nil{
		panic(err)
	}

	
}