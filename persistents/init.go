package persistents

import (
	"github.com/naufalfmm/csv-transaction/persistents/repositories"
	"github.com/naufalfmm/csv-transaction/resources/config"
	"github.com/naufalfmm/csv-transaction/resources/db"
	"github.com/naufalfmm/csv-transaction/utils/logger"
)

type Persistents struct {
	Repositories repositories.Repositories
}

func Init(d *db.DB, l logger.Logger, c *config.EnvConfig) (Persistents, error) {
	repo, err := repositories.Init(d, l)
	if err != nil {
		return Persistents{}, err
	}

	return Persistents{
		Repositories: repo,
	}, nil
}
