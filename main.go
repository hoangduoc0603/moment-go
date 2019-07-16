package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hoangduoc0603/moment/driver"
	ah "github.com/hoangduoc0603/moment/handler/http"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Active Server Successfully")

	// Lấy về connection tới DB
	connection, err := driver.ConnectSQL("127.0.0.1", "3306", "host", "Kevin123", "test")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	myRouter := mux.NewRouter().StrictSlash(true)
	aHandler := ah.NewPostHandler(connection)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/createAccount", aHandler.Create).Methods("POST")

	fmt.Println("Server listen at :8081")
	http.ListenAndServe(":8081", myRouter)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
