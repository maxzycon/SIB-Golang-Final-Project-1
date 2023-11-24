package impl

import (
	"github.com/maxzycon/SIB-Golang-Final-Project-1/internal/config"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/internal/domain/global/repository"
	"gorm.io/gorm"
)

type NewGlobalServiceParams struct {
	Conf             *config.Config
	GlobalRepository repository.GlobalRepository
	Db               *gorm.DB
}
type GlobalService struct {
	conf             *config.Config
	globalRepository repository.GlobalRepository
	db               *gorm.DB
}

func New(params *NewGlobalServiceParams) *GlobalService {
	return &GlobalService{
		conf:             params.Conf,
		globalRepository: params.GlobalRepository,
		db:               params.Db,
	}
}
