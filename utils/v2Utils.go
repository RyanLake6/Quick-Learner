package utils

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/labstack/gommon/log"
)

func RunJobs(keywords []string) (res []map[string]string, err error) {
	numJobs, numOfKeywordsInJobs := GetOptimizedSplit(keywords)

	jobs := make(chan []string, numJobs)
	results := make(chan map[string]string, numJobs)

	// starting workers (ready to take in jobs)
	numWorkers := GetOptimalNumWorkers(numJobs)
	for w := 1; w <= numWorkers; w++ {
		go Worker(w, jobs, results)
	}

	// Building jobs to send to workers
	splitJobs, err := SplitArrayEvenly(keywords, numOfKeywordsInJobs)
	if err != nil {
		log.Errorf("Failed to split up keywords: %v", err)
		return nil, err
	}

	// Jobs are sent to the workers
	for _, job := range splitJobs {
		jobs <- job
	}
	close(jobs)

	res = []map[string]string{}
	for a := 1; a <= numJobs; a++ {
		res = append(res, <-results)
	}

	// TODO: put this output into one large map not an array of maps
	return res, nil
}

// TODO return err channel from GetAllWikiLinks
func Worker(id int, jobs <-chan []string, results chan<- map[string]string) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		// time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		links, _ := GetAllWikiLinks(j)
		results <- links
	}
}

func SplitArrayEvenly(array []string, keywordsInSplit int) (splitArray [][]string, err error) {
	if len(array) == 0 {
		return nil, errors.New("Can't split up empty array")
	}

	var j int
	ret := [][]string{}

	for i := 0; i < len(array); i += keywordsInSplit{
    	j += keywordsInSplit
    	if j > len(array) {
        	j = len(array)
    	}
    	ret = append(ret, array[i:j])
	}
	return ret, nil
}

func GetOptimizedSplit(keywords []string) (numOfJobs int, numOfKeywordsInJobs int) {
	// TODO: get env variable of how many splits it should do (if empty optimize)
	// Fails if the split is too high for a small ingestion of data
	numOfKeywordsInJobs = len(keywords) / 2
	if len(keywords) < 2 {
		numOfKeywordsInJobs = len(keywords)
	}

	numOfJobs = len(keywords) / numOfKeywordsInJobs

	return numOfJobs, numOfKeywordsInJobs
}

func GetOptimalNumWorkers(numJobs int) (workers int) {
	workers = numJobs / 2
	if workers > runtime.NumCPU() {
		workers = runtime.NumCPU()
	}

	return workers
}