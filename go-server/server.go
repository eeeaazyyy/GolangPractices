package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Age    int    `json:"age"`
	Name   string `json:"name"`
	Height int32  `json:"height"`
	Weight int8   `json:"weight"`
}

func (p Person) toString() string {
	return fmt.Sprintf("Name: %s, Age: %d, Height: %d cm, Weight: %d kg",
		p.Name, p.Age, p.Height, p.Weight)
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	person := Person{
		Age:    21,
		Name:   "Vlad Trifonov",
		Height: 175,
		Weight: 75,
	}

	w.Header().Set("Counter-Type", "application/json")

	if err := json.NewEncoder(w).Encode(person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Sent person data:", person.toString())

}

func main() {
	http.HandleFunc("/person", personHandler)

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server error: ", err)
	}
}
