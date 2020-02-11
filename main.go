//package main :
// File main.go
package main

import "iam/core/database"

func main() {
	idb := database.NewIAMDatabase("test.db")
	idb.Setup()
}
