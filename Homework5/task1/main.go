package main

import (
	"fmt"
	"time"
)

func Worker(id int, jobs <-chan string, results chan<- string) {
	for v := range jobs {
		fmt.Println("Worker with id: ", id, " start job ", v)
		time.Sleep(35 * time.Second)
		results <- v + " is done."
	}
}

func main() {
	numberOfWokers := 10
	jobList := []string{"Manager", "Developer", "CEO", "Doctor", "Actor"}
	numberOfJobs := len(jobList)

	jobs := make(chan string, numberOfJobs)
	results := make(chan string, numberOfJobs)

	for i := 0; i < numberOfWokers; i++ {
		go Worker(i, jobs, results)
	}

	for _, job := range jobList {
		jobs <- job
	}
	close(jobs)

	for a := 1; a <= numberOfJobs; a++ {
		fmt.Println(<-results)
	}

	fmt.Println("All jobs have been finished.")
}