package main

import (
    "fmt"
	"math/rand"
	"log"
	"encoding/json"
	"net/http"
)

func indexPage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the Index Page!")
}

func pingHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "pong")
}

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request){
    keys, ok := r.URL.Query()["name"]

	var nameVal = "";
    if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		nameVal = "Stranger";
    } else {
		nameVal = keys[0];
	}

	var id = rand.Intn(1000000000);
	var user = User{Id: id, Name: "Hello " + nameVal};
	bytes, err := json.Marshal(user);

	if err != nil {
		fmt.Println("Failed to serialize to json")
	} else {
		log.Println(string(bytes));
		fmt.Fprintf(w, string(bytes));
	}
}

type SumParams struct {
	X int `json:"x"`
	Y int `json:"y"`
}


type SumResult struct {
	Sum int `json:"sum"`
}

func sumHandler(w http.ResponseWriter, req *http.Request){
	decoder := json.NewDecoder(req.Body);
	var params SumParams
	err := decoder.Decode(&params)

	var res = SumResult{Sum: params.X + params.Y}

	bytes, err := json.Marshal(res);

	if err != nil {
		fmt.Println("Failed to serialize to json")
	} else {
		log.Println(string(bytes));
		fmt.Fprintf(w, string(bytes));
	}
}

func handleRequests() {
	http.HandleFunc("/", indexPage);
	http.HandleFunc("/ping", pingHandler);
	http.HandleFunc("/hello-world", helloWorldHandler);
	http.HandleFunc("/sum", sumHandler);
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    handleRequests()
}
