package main

import (
	"github.com/brunomiz/jeguin/domain"
	"github.com/brunomiz/jeguin/ports"
	"net/http"
)

func main() {
	var jobsRepository ports.JobsRepository = nil
	jobs := jobsRepository.GetJobs()
	serveJobs(jobs)
}

func serveJobs(jobs []domain.Job) {
	for _, job := range jobs {
		token := job.GetToken()
		execution := job.GetExecution()
		handler := generateHandler(execution)
		http.HandleFunc(token, handler)
	}
}

func generateHandler(execution func()) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		execution()
	}
}
