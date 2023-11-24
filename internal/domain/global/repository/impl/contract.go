package impl

import (
	"github.com/maxzycon/SIB-Golang-Final-Project-1/internal/config"
	"gorm.io/gorm"
)

type NewGlobalRepository struct {
	Conf *config.Config
	Db   *gorm.DB
}
type GlobalRepository struct {
	conf *config.Config
	db   *gorm.DB
}

func New(params *NewGlobalRepository) *GlobalRepository {
	return &GlobalRepository{
		conf: params.Conf,
		db:   params.Db,
	}
}
