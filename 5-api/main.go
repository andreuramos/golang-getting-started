package main

import "fmt"

type Album struct {
	ID		string 	`json:"id"`
	Title	string	`json:"title"`
	Artist	string	`json:"artist"`
	Price	float64	`json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Raining Blood", Artist: "Slayer", Price: 49.99},
	{ID: "2", Title: "Toxicity", Artist: "System of a Down", Price: 25.00},
	{ID: "3", Title: "Master of Puppets", Artist: "Metallica", Price: 39.95},
	{ID: "4", Title: "Battle of Los Angeles", Artist: "Rage Against the Machine", Price: 56.99},
}

func main() {
	fmt.Println(albums)
}