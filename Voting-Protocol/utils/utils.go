package utils

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"voting_protocol/node"
)

func NewPriorityQueue() *node.PriorityQueue {
	pq := make(node.PriorityQueue, 0)
	heap.Init(&pq)
	return &pq
}

func ReadNodesList() map[int]string {
	jsonFile, err := os.Open("nodes-list.json")
	if err != nil {
		fmt.Println("Error opening nodes-list.json file:", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var nodesList map[int]string

	json.Unmarshal(byteValue, &nodesList) // Puts the byte value into the nodesList map

	return nodesList
}

// Calculate the time taken from the first node to request to the last node to exist the critical section
func CalculateTimeTaken(n *node.Node, numRequests int) {
	startTime := time.Now()

	if n.ID == 0 {
		n.Finished = make([]bool, numRequests)
		for {
			if all(n.Finished) {
				fmt.Printf("Time taken for all nodes to exit the critical section: %v\n", time.Since(startTime))
				break
			}
		}
	}
}

func all(arr []bool) bool {
	for _, v := range arr {
		if !v {
			return false
		}
	}
	return true
}