package main

import (
	"fmt"
)

var indexOutOfBounds = false

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

			// prev existence and avoiding pointing to the same node
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

	// All transactions are present in Global Transaction List
	fmt.Println("Task 1 complete")
	fmt.Println("Size of Transaction List: ", len(globalTransactionList))

	// All parents for each node are present in direct parent map
	fmt.Println("Task 2 complete")
	fmt.Println("Length of all transactions", len(directParentMap))

	// Direct Parent - Indirect Parent Count
	dipMapCount := make(map[string]int)

	// Iterate over all transactions and get their dip map
	// To get the sum of indirect and direct parents
	// We iterate over all transactions who have a direct parent
	// dip(A) = dp(A) + ip(A)
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

	// Task 3: All direct + indirect parent map count is present in dipMapCounts
	sortAndPrintTop10(dipMapCount)

}
