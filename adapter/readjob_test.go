package adapter

import (
	"testing"
)

func TestConvertToJob(t *testing.T) {
	got := ConvertToJob("resources/testJobs.txt")
	job := got[0]
	jobToken := job.GetToken()
	want := "myjob"
	if jobToken != want {
		t.Errorf("wanted [%s] but got [%s]", want, jobToken)
	}
	if len(got) != 2 {
		t.Errorf("wanted 2 jobs but got %d", len(got))
	}
}
