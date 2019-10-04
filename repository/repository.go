package repository

import "minesweeper-API/models"

type MinesweeperRepository interface {
	UpdateGame(game *models.Game) error
	PauseGame(game *models.Game) error
	NewGame(game *models.Game) error
	NewPlayer(User *models.Player) error
	GetGamesByPlayerId(playerId string) (games *[]models.Game, err error)
}
