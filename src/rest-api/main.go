package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)


type User struct {
    ID int `json:"id"`
    Avatar string `json:"avatar"`
    First_name string `json:"first_name"`
    Last_name string `json:"last_name"`
}

type Users []User

func getAllUser(w http.ResponseWriter, r *http.Request) {
	users := Users{
        User{
            ID: 1,
            Avatar: "https://s3.amazonaws.com/uifaces/faces/twitter/calebogden/128.jpg",
            First_name: "George",
            Last_name: "Bluth",
        },
        User{
            ID: 2,
            Avatar: "https://s3.amazonaws.com/uifaces/faces/twitter/josephstein/128.jpg",
            First_name: "Janeth",
            Last_name: "Weaver",
        },
        User{
            ID: 3,
            Avatar: "https://s3.amazonaws.com/uifaces/faces/twitter/olegpogodaev/128.jpg",
            First_name: "Emma",
            Last_name: "Wong",
        },
    }
    fmt.Println(users)
    fmt.Println("Endpoint for all users")
    json.NewEncoder(w).Encode(users)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", getAllUser)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	
	handleRequests()
}