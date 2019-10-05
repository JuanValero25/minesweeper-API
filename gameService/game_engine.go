package gameService

import (
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"math/rand"
	"minesweeper-API/models"
	"minesweeper-API/repository"
	"minesweeper-API/repositoryPosgress"
	"time"
)

type MinesweeperGameEngine struct {
	Repository repository.MinesweeperRepository
}

// this a simple builder to make new engine it ca create new repos base on diferentes DB and connectors
func NewGameEngine(repo *repositoryPosgress.MineSweeperPostgresRepo) *MinesweeperGameEngine {
	return &MinesweeperGameEngine{Repository: repo}
}

//function to create new Game
func (GameEngine *MinesweeperGameEngine) NewBoardGame(rowNumber, colNumber, minesCount int, gameOwner string) (newGame *models.Game, err error) {
	//newGame = &models.Game{}
	if rowNumber < 6 && colNumber < 6 {
		err = errors.New("RowNumber and ColNumber must be bigger than 3")
		return
	}
	cellNumber := colNumber * rowNumber
	cells := make(models.CellGrid, cellNumber)

	i := 0
	for i < minesCount {
		index := rand.Intn(cellNumber)
		if !cells[index].Mine {
			cells[index].Mine = true
			i++
		}
	}
	id := uuid.NewV4()

	newGame = &models.Game{
		GameId:   id.String(),
		PlayerId: gameOwner,
		Rows:     rowNumber,
		Cols:     colNumber,
		Mines:    minesCount,
		Status:   models.STARTED,
		Timer:    time.Now(),
		Clicks:   0,
		Duration: 0,
	}

	//now make slice to matrix
	newGame.Grid = make([]models.CellGrid, rowNumber)
	for row := range newGame.Grid {
		newGame.Grid[row] = cells[(colNumber * row):(colNumber * (row + 1))]
	}
	initPotition(newGame)

	err = GameEngine.Repository.NewGame(newGame)

	if err != nil {
		fmt.Printf("error creating new models: %s", err)
		newGame = nil
		return
	}
	cleanMinesForResponse(newGame)
	return
}

// this function check if player win
func (GameEngine *MinesweeperGameEngine) ClickCell(game *models.Game, i, j int) error {

	currentGame, err := GameEngine.Repository.GetGameById(game.GameId)
	if err != nil {
		return errors.New("error geting current game")
	}
	game.Grid = currentGame.Grid
	if !isValidRevealable(game, i, j) {
		return errors.New("position out of the grid")
	}

	if game.Grid[i][j].Clicked {
		return errors.New("cell already clicked")
	}
	// not valid UUID
	if len(game.GameId) < 32 {
		return errors.New("gameId invalid")
	}

	if game.Status != models.STARTED {
		return errors.New("game finish or paused")
	}
	game.Grid[i][j].Clicked = true
	game.Clicks += 1
	if game.Grid[i][j].Mine {
		game.Status = models.DEFEAT
	}
	if checkIfWin(game) {
		game.Status = models.WON
	}
	revealAdjacent(game, i, j)
	return GameEngine.Repository.UpdateGame(game)
}

func (GameEngine *MinesweeperGameEngine) GamePause(game *models.Game) error {
	game.Status = models.PAUSED
	game.Timer = time.Now()
	return GameEngine.Repository.UpdateGame(game)
}

func checkIfWin(game *models.Game) bool {
	return game.Clicks == ((game.Rows * game.Cols) - game.Mines)
}

func cleanMinesForResponse(game *models.Game) {
	for y, col := range game.Grid {
		for x, _ := range col {
			game.Grid[x][y].Mine = false
		}

	}
}

func initPotition(game *models.Game) {
	for y, col := range game.Grid {
		for x, _ := range col {
			game.Grid[x][y].RowValue = x
			game.Grid[x][y].ColumnValue = y
		}

	}
}

func revealAdjacent(game *models.Game, x, y int) {
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if isValidRevealable(game, i, j) {
				fmt.Printf("points revealed  %d %d \n", i, j)
				game.Grid[i][j].Clicked = true
			}
		}
	}

}

//this check
func isValidRevealable(game *models.Game, x, y int) bool {
	return (y <= game.Cols-1 && y > 0) && (x <= game.Rows-1 && x > 0) && !game.Grid[x][y].Mine
}
