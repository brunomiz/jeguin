package ports

import "github.com/brunomiz/jeguin/domain"

type JobsRepository interface {
	GetJobs() []domain.Job
}
