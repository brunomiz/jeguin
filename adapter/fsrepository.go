package adapter

import "github.com/brunomiz/jeguin/domain"

type FileSystemRepository struct{}

func (fs *FileSystemRepository) GetJobs() []domain.Job {
	return ConvertToJob("jobs.txt")
}

func NewJobsRepository() *FileSystemRepository {
	return &FileSystemRepository{}
}
