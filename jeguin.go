package main

import (
	"fmt"
	"github.com/brunomiz/jeguin/adapter"
	"github.com/brunomiz/jeguin/domain"
	"github.com/brunomiz/jeguin/ports"
	"net/http"
)

func main() {
	var jobsRepository ports.JobsRepository = adapter.NewJobsRepository()
	jobs := jobsRepository.GetJobs()
	serveJobs(jobs)
}

func serveJobs(jobs []domain.Job) {
	for _, job := range jobs {
		token := job.GetToken()
		handler := generateHandler(job)
		http.HandleFunc("/"+token, handler)
	}
	port := "8080"
	fmt.Println("Listening on port", port)
	err := http.ListenAndServe("localhost: "+port, nil)
	panic(err)
}

func generateHandler(job domain.Job) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		execution := job.GetExecution()
		execution()
	}
}
