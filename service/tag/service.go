package tag

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewService,
)

type Service struct {
	repo Repo
}

func NewService(repo Repo) *Service {
	return &Service{repo: repo}
}

type Repo interface {
}
