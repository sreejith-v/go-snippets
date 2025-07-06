package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type ResponseModel struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type ChanStruct struct {
	Title     string
	Completed string
	Id        int
}

const todoBaseUrl = "https://jsonplaceholder.typicode.com/todos/"
const limit = 100
const (
	NotComplete = "not complete"
	Complete    = "complete"
)

func getURLResponses(url string) (ResponseModel, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status code: %d, body: %s\n", resp.StatusCode, string(body))
		os.Exit(1)
	}

	// Unmarshal the JSON response into the struct
	var data ResponseModel
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		os.Exit(1)
	}

	return data, nil
}

func printURLDataResponse(url string) {
	isCompleteStatus := NotComplete
	responses, err := getURLResponses(url)
	if err != nil {
		return
	}
	if responses.Completed {
		isCompleteStatus = Complete
	}
	log.Printf("Todo item %v with title: %v; is %v", responses.Id, responses.Title, isCompleteStatus)
}

func laymanApproach() {
	for i := 0; i < limit; i++ {
		url := todoBaseUrl + strconv.Itoa(i+1)
		printURLDataResponse(url)
	}
}

func printURLDataResponseWithChannel(url string, ch chan ChanStruct, wg *sync.WaitGroup) {
	defer wg.Done()
	isCompleteStatus := NotComplete
	responses, err := getURLResponses(url)
	if err != nil {
		return
	}
	if responses.Completed {
		isCompleteStatus = Complete
	}
	ch <- ChanStruct{Title: responses.Title, Completed: isCompleteStatus, Id: responses.Id}
}

func simpleGoRoutine() {
	ch := make(chan ChanStruct)
	wg := sync.WaitGroup{}
	for i := 0; i < limit; i++ {
		url := todoBaseUrl + strconv.Itoa(i+1)
		log.Printf("calling go routine for %v", i+1)
		wg.Add(1)
		go printURLDataResponseWithChannel(url, ch, &wg)
	}

	go func() {
		wg.Wait() // Wait for all goroutines to finish
		close(ch) // Close the channel to signal completion
	}()

	for results := range ch {
		log.Printf("Todo item %v with title: %v; is %v", results.Id, results.Title, results.Completed)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	startTime := time.Now()
	log.Println("Approach 1 Starting")
	laymanApproach()
	//simpleGoRoutine()
	endTime := time.Now()
	log.Printf("Completed in %v", endTime.Sub(startTime))
}
