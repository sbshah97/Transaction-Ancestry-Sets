package main

import (
	"fmt"
	"sort"
)

var indexOutOfBounds = false

type kv struct {
	Key   string
	Value int
}

func main() {

	// Fetch hash using Block height
	blockHash, _ := fetchHashUsingBlockHeight(initialBlockNumber)

	var globalTransactionList []Transaction

	indexOutOfBounds := false
	index := 0
	// This is an index, will keep of all transactions in the chain
	indexMap := make(map[string]bool)
	for !indexOutOfBounds {
		// Fetch all the transactions for that block using the hash
		transactionList, err := fetchTransactionListUsingHash(blockHash, index)
		index += 1

		for index := range transactionList {
			indexMap[transactionList[index].TxId] = true
		}

		globalTransactionList = append(globalTransactionList, transactionList...)

		if err != nil {
			break
		}
	}

	// Count direct parents
	directParentCount := make(map[string]int)

	// Map of direct parents
	directParentMap := make(map[string][]string)

	for index := range globalTransactionList {
		for _, inputTransaction := range globalTransactionList[index].VIn {

			// prev existence
			if indexMap[inputTransaction.TxId] && globalTransactionList[index].TxId != inputTransaction.TxId {
				// count++
				directParentCount[globalTransactionList[index].TxId] += 1
				//
				directParentMap[globalTransactionList[index].TxId] = append(
					directParentMap[globalTransactionList[index].TxId],
					inputTransaction.TxId,
				)
			}
		}
	}

	fmt.Println("Task 1 complete")
	fmt.Println("Size of Transaction List: ", len(globalTransactionList))

	fmt.Println("Task 2 complete")
	fmt.Println("Length of all transactions", len(directParentMap))

	// Direct Parent - Indirect Parent Count
	dipMapCount := make(map[string]int)

	// Iterate over all transactions and get their dip map
	for child := range directParentMap {

		searchQueue := []string{}
		searchQueue = enqueue(searchQueue, child)

		for len(searchQueue) != 0 {
			// The direct + indirect map count will be combination of
			// direct and indirect nodes
			dipMapCount[child] += directParentCount[front(searchQueue)]

			//
			for index := range directParentMap[front(searchQueue)] {
				// Enqueue the children
				searchQueue = enqueue(searchQueue, directParentMap[front(searchQueue)][index])
			}

			// Dequeue the front node
			searchQueue = dequeue(searchQueue)
		}
	}

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
