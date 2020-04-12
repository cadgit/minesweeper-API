package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"minesweeper-API/types"
)

type Author struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	//redis://h:p83bf83e9dbccfa8103e03cf22fc5b67a7c90da66c31a9a204f86db7784d8d6cd@ec2-34-194-198-64.compute-1.amazonaws.com:20849
	client := redis.NewClient(&redis.Options{
		Addr: "ec2-34-194-198-64.compute-1.amazonaws.com:20849",
		Password: "p83bf83e9dbccfa8103e03cf22fc5b67a7c90da66c31a9a204f86db7784d8d6cd",
		DB: 0,
	})

	var game =  types.Game{
		Name: "test-game",
		Rows: 3,
		Cols: 2,
		Mines: 4,
		Status: "new",
	}


	//json, err := json.Marshal(Author{Name: "Elliot", Age: 25})
	//if err != nil {
	//	fmt.Println(err)
	//}

	gameAsJson, err := json.Marshal(game)
	if err != nil {
		fmt.Println(err)
	}

	//err = client.Set("id1234", json, 0).Err()
	err = client.Set(game.Name, gameAsJson, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	//val, err := client.Get("id1234").Result()
	val, err := client.Get(game.Name).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

	//birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	//var bird Bird
	var gameStored types.Game
	json.Unmarshal([]byte(val), &gameStored)
	fmt.Printf("Name: %s, Status: %s", gameStored.Name, gameStored.Status)
}
