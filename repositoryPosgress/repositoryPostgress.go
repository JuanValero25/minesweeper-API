package repositoryPosgress

import (
	"errors"
	"github.com/go-pg/pg"
	"minesweeper-API/models"
)

type MineSweeperPostgresRepo struct {
	*pg.DB
}

func New() *MineSweeperPostgresRepo {
	return &MineSweeperPostgresRepo{buildPgConnection()}
}

func (repo *MineSweeperPostgresRepo) UpdateGame(game *models.Game) error {
	return repo.Update(game)
}
func (repo *MineSweeperPostgresRepo) NewGame(game *models.Game) error {
	return repo.Insert(game)
}

func (repo *MineSweeperPostgresRepo) NewPlayer(player *models.Player) error {
	err := repo.Insert(player)
	return player, err
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
