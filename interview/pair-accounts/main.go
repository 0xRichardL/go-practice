package main

import (
	"fmt"
	"sort"
)

// A set of customer/account pairs indicates which customers owns which accounts.
// Task
// In your choice of programming language, write a function that finds all customers who share all the
// same accounts.
// Example
// The following example indicates which customers own which accounts. For instance, the customer with
// id 1 owns the account with id 10 and the account with id 11.
// Customer    Account
// 1           10
// 1           11
// 2           13
// 3           11
// 4           14
// 3           10
// 4           13
// For this example:
// Customers 1 and 3 share all their accounts, 10 and 11, so they are a match.
// Customers 2 and 4 share account 13 but customer 4 also owns account 14, which customer 2 does
// not. They are not a match.

// Input : 2 arrays / slices customers and accounts with the same length.
// Output: print matching customer pairs .

// Customer    Account
// 1           10
// 1           11
// 2           13
// 3           11
// 4           14
// 3           10
// 4           13
// 5 					 10
// 5 					 11
// 6 					 13

type Pair struct {
	CID int
	AID int
}

func pairAccounts(pairs []Pair) {
	cusM := make(map[int][]int)
	hashM := make(map[string][]int)
	for _, p := range pairs {
		if _, ok := cusM[p.CID]; !ok {
			cusM[p.CID] = make([]int, 0)
		}
		cusM[p.CID] = append(cusM[p.CID], p.AID)
	}
	for cID, aIDs := range cusM {
		sort.Ints(aIDs)
		key := fmt.Sprint(aIDs)
		if _, ok := hashM[key]; !ok {
			hashM[key] = make([]int, 0)
		}
		hashM[key] = append(hashM[key], cID)
	}
	for _, cIDs := range hashM {
		if len(cIDs) > 1 {
			fmt.Println(cIDs)
		}
	}
}

func main() {
	input := []Pair{
		// {1, 10},
		// {1, 11},
		// {2, 13},
		// {3, 11},
		// {4, 14},
		// {3, 10},
		// {4, 13},
		////////
		{1, 10},
		{1, 11},
		{2, 13},
		{3, 11},
		{4, 14},
		{3, 10},
		{4, 13},
		{5, 10},
		{5, 11},
		{6, 13},
	}
	pairAccounts(input)
}
