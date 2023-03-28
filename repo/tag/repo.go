package tag

import (
	"grpc-project/service/tag"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(
	NewRepo,
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) tag.Repo {
	return &Repo{
		db: db,
	}
}
