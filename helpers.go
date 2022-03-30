package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

func makeGetRequest(uri string) ([]byte, error) {
	// Step 1: Make a GET call using Block ID number to get hash from Block ID
	resp, err := http.Get(host + uri)
	if err != nil {
		log.Fatalln(err)
		return []byte{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		indexOutOfBounds = true
		return nil, errors.New(errorString)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return body, err
	}

	return body, nil

}

func fetchHashUsingBlockHeight(blockNumber string) (string, error) {
	// Step 2: Use hash to get
	uri := "/block-height/" + blockNumber
	blockHeightBody, err := makeGetRequest(uri)

	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	//Convert the body to type string
	return string(blockHeightBody), nil

}

func fetchTransactionListUsingHash(txHash string, index int) ([]Transaction, error) {
	var transactionList []Transaction
	offset := ""
	if index != 0 {
		offset = "/" + fmt.Sprintf("%d", index*25)
	}

	uri := "/block/" + txHash + "/txs" + offset
	body, err := makeGetRequest(uri)

	if err != nil {
		return transactionList, err
	}

	err = json.Unmarshal(body, &transactionList)
	if err != nil {
		log.Fatalln(err)
		return transactionList, err
	}

	return transactionList, nil
}

func enqueue(queue []string, element string) []string {
	queue = append(queue, element) // Simply append to enqueue.
	return queue
}

func dequeue(queue []string) []string {
	// element := queue[0] // The first element is the one to be dequeued.
	return queue[1:] // Slice off the element once it is dequeued.
}

func front(queue []string) string {
	return queue[0]
}

// Helper function to sort and print top 10 values from key value pair
func sortAndPrintTop10(dipMapCount map[string]int) {
	var ss []kv
	for k, v := range dipMapCount {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	count := 0
	fmt.Println("Task 3 complete")
	for _, kv := range ss {
		if count > 10 {
			break
		}
		fmt.Println(kv.Key, kv.Value)
		count += 1
	}

}
