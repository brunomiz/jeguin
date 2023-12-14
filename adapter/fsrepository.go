package adapter

import "github.com/brunomiz/jeguin/domain"

type fileSystemRepository struct{}

func (fs *fileSystemRepository) GetJobs() []domain.Job {
	return []domain.Job{}
}
