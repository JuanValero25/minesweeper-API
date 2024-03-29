package repositoryPosgress

import (
	"errors"
	"fmt"
	"github.com/go-pg/pg"
	"minesweeper-API/models"
)

type MineSweeperPostgresRepo struct {
	*pg.DB
}

func New() *MineSweeperPostgresRepo {
	pgConnection := buildPgConnection()
	pgConnection.AddQueryHook(dbLogger{})
	return &MineSweeperPostgresRepo{pgConnection}
}

func (repo *MineSweeperPostgresRepo) UpdateGame(game *models.Game) error {
	return repo.Update(game)
}
func (repo *MineSweeperPostgresRepo) NewGame(game *models.Game) error {
	return repo.Insert(game)
}

func (repo *MineSweeperPostgresRepo) NewPlayer(player *models.Player) error {
	return repo.Insert(player)
}

func (repo *MineSweeperPostgresRepo) GetGameById(player string) (game *models.Game, err error) {
	game = &models.Game{GameId: player}
	err = repo.Select(game)
	return
}

func (repo *MineSweeperPostgresRepo) PauseGame(game *models.Game) (err error) {
	result, err := repo.Model(game).Column("status").WherePK().Update()
	if result != nil && result.RowsAffected() < 1 {
		err = errors.New("no rows affected on DB")
	}
	return
}
func (repo *MineSweeperPostgresRepo) GetGamesByPlayerId(playerId string) (games *[]models.Game, err error) {
	err = repo.Model(games).Where("player_id = ?", playerId).Select(games)
	return
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {
	fmt.Println(q.Query)
}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	fmt.Println(q.FormattedQuery())
}
