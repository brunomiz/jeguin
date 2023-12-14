package adapter

import (
	"bufio"
	"github.com/brunomiz/jeguin/domain"
	"log"
	"os"
	"strings"
)

func convertToJob(filePath string) []domain.Job {
	lines := readFile(filePath)
	var jobs []domain.Job
	var job domain.Job
	for _, line := range lines {
		if strings.Contains("[", line) {
			if &job == nil {
				job = domain.Job{}
				job.SetToken(strings.Trim(line, "[]"))
			} else {
				jobs = append(jobs, job)
			}
		} else {
			job.AddExecutionCommand(line)
		}

	}
	return jobs
}

func readFile(filePath string) []string {
	var lines []string
	file, err := os.Open(filePath)
	if os.IsNotExist(err) {
		log.Fatal(err)
	}
	check(err)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
