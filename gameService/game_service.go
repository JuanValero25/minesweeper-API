package gameService

import (
	"github.com/satori/go.uuid"
	"minesweeper-API/models"
)

type GameService struct {
	GameEngine *MinesweeperGameEngine
}

func (gameService *GameService) CreateNewGame(rowNumber, colNumber, minesCount int, gameOwner string) (newGame *models.Game, err error) {
	return gameService.GameEngine.NewBoardGame(rowNumber, colNumber, minesCount, gameOwner)
}

func (gameService *GameService) ClickCell(game *models.Game, i, j int) error {
	return gameService.GameEngine.ClickCell(game, i, j)
}

func (gameService *GameService) PauseGame(gameId string) error {
	gameToPause := &models.Game{
		GameId: gameId,
		Status: models.PAUSED,
	}
	return gameService.GameEngine.Repository.PauseGame(gameToPause)
}

func (gameService *GameService) NewPlayer(player *models.Player) (err error) {
	id := uuid.NewV4()
	player.Id = id.String()
	return gameService.GameEngine.Repository.NewPlayer(player)
}
