package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"minesweeper-API/gameService"
	"minesweeper-API/repositoryPosgress"
	"minesweeper-API/routes"
	"net/http"
)

func main() {

	gameRouter := routes.GameRouter{
		GameService: &gameService.GameService{
			GameEngine:&gameService.MinesweeperGameEngine{Repository:repositoryPosgress.New()},
		},
	}
	router := httprouter.New()

	router.POST("/newPlayer", gameRouter.NewPlayer)
	router.POST("/newGame", gameRouter.NewGame)
	router.POST("/ClickGame", gameRouter.ClickGame)

	log.Fatal(http.ListenAndServe(":8080", router))
}