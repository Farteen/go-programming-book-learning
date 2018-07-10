package main

import (
	"os"
	"fmt"
	"log"
	"github.com/Farteen/Go-Programming-Book-Learning/4.5/github"
)

type Movie struct {
	Title string 
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issue: \n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	// var movies = []Movie {
	// 	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	// 	{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	// 	{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// }
	// jsonStr, err := json.MarshalIndent(movies, "", "  ")
	// if err != nil {
	// 	return
	// }
	// fmt.Printf("%s", jsonStr)
}