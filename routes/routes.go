package routes

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"minesweeper-API/gameService"
	"minesweeper-API/models"
	"net/http"
)

type GameRouter struct {
	GameService *gameService.GameService
}

type SimpleResponse struct {
	Message string `json:"message"`
}


func (router *GameRouter) NewPlayer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	player := &models.Player{}
	err := json.NewDecoder(r.Body).Decode(player)
	if err != nil {
		http.Error(w, "Error encoding the body", 500)
	}
	err = router.GameService.NewPlayer(player)
	if err != nil {
		http.Error(w, "error creating new player", 404)
	}
	_, err = fmt.Fprint(w, "player created OK\n")
	if err != nil {
		fmt.Println("error printing response ")
	}

}

func (router *GameRouter) NewGame(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	game := &models.GameRequest{}
	err := json.NewDecoder(r.Body).Decode(game)
	if err != nil {
		http.Error(w, "Error encoding the body", 500)
	}
	gameCreated, err := router.GameService.CreateNewGame(game.Rows, game.Cols, game.Mines, game.PlayerId)
	if err != nil {
		http.Error(w, "error creating new player", 404)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(gameCreated)
	if err != nil {
		fmt.Println("error printing response ")
	}
}

func (router *GameRouter) ClickGame(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	game := &models.GameClick{}
	err := json.NewDecoder(r.Body).Decode(game)
	if err != nil {
		http.Error(w, "Error encoding the body", 500)
		return
	}
	err = router.GameService.ClickCell(game.Game, game.PositionX, game.PositionY)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(game.Game)
	if err != nil {
		fmt.Println("error printing response ")
	}

}

func (router *GameRouter) PauseGame(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	gameId := params.ByName("gameId")

	err := router.GameService.PauseGame(gameId)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(SimpleResponse{Message:"game paused"})
	if err != nil {
		fmt.Println("error printing response ")
	}

}
