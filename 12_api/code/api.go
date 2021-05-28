package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const BaseURL = "https://swapi.dev/api/"

type People struct{
	People []Person `json:"results"`
}

type Person struct{
	Name string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homeworld Planet
}

type Planet struct{
	Name string `json:"name"`
	Terrain string `json:"terrain"`
	Population string `json:"population"`
}

func (people *People) updatePeople(homeworld Planet){
	for i := range people.People {
		people.People[i].Homeworld = homeworld
	}
}

func (person *Person) getHomeWorld() Person{
	res, err := http.Get(person.HomeworldURL)

	if err != nil{
		log.Print("Failed to fetch resources.", err)
	}

	var bytes []byte
	if bytes, err = ioutil.ReadAll(res.Body); err != nil {
		log.Print("Failed to read response body", err)
	}

	json.Unmarshal(bytes, &person.Homeworld)

	return *person
}

func people(w http.ResponseWriter, r * http.Request){
	response , err := http.Get(BaseURL + "people")

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to fetch resources.")
	}

	//fmt.Println(response)

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to parse request body")
	}

	//fmt.Println(bytes)

	var people People
	if err := json.Unmarshal(bytes, &people); err != nil {
		fmt.Println("Error parsing json", err)
	}

	for _, person := range people.People{
		p := person.getHomeWorld()
		people.updatePeople(p.Homeworld)
	}

	fmt.Println(people)
}

func main() {
	var PORT = ":8080"
	http.HandleFunc("/people", people)
	fmt.Println("Server on", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
