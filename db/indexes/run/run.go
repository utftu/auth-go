package main

import "auth-go/db/indexes"

func main() {
	indexes.CreateUserIndexes()
	indexes.CreateClientIndexes()
}