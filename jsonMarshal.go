package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	movies := []Movie{
		{Title: "Casablanca", Year: 1942, Color: true, Actors: []string{"Jin Yang", "Mike"}},
		{Title: "God Father", Year: 1994, Actors: []string{"Jin Yang", "Mike"}},
		{Title: "Casablanca", Year: 1942, Color: true, Actors: []string{"Jin Yang", "Mike"}},
		{Title: "Casablanca", Year: 1942, Color: true, Actors: []string{"Jin Yang", "Mike"}},
	}
	var m []Movie
	data, _ := json.MarshalIndent(movies, "", "    ")
	fmt.Printf("%s\n", data)

	if err := json.Unmarshal(data, &m); err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
}
