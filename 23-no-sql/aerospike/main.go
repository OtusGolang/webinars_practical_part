package main

import (
	"fmt"
	"log"

	as "github.com/aerospike/aerospike-client-go"
)

func main() {
	// 1. Connect to an Aerospike cluster
	client, err := as.NewClient("127.0.0.1", 3000)
	if err != nil {
		log.Fatalf("Failed to connect to the cluster: %v", err)
	}
	defer client.Close()

	namespace := "test"
	set := "demoSet"
	keyValue := "demoKey"

	// 2. Creating a new record
	writePolicy := as.NewWritePolicy(0, 0)
	recordKey, err := as.NewKey(namespace, set, keyValue)
	if err != nil {
		log.Fatalf("Failed to create a key: %v", err)
	}

	binMap := as.BinMap{
		"name":    "John Doe",
		"age":     30,
		"isAlive": 1, // use 1 for true and 0 for false
	}

	err = client.Put(writePolicy, recordKey, binMap)
	if err != nil {
		log.Fatalf("Failed to put record: %v", err)
	}
	fmt.Println("Record created!")

	// 3. Reading a record
	record, err := client.Get(nil, recordKey)
	if err != nil {
		log.Fatalf("Failed to get record: %v", err)
	}
	if record == nil {
		log.Fatalf("No record found with key: %v", keyValue)
	}
	fmt.Printf("Record read: Name=%s, Age=%d, IsAlive=%v\n", record.Bins["name"], record.Bins["age"], record.Bins["isAlive"])

	// 4. Updating a record
	record.Bins["age"] = 31
	err = client.PutBins(writePolicy, recordKey, as.NewBin("age", 31))
	if err != nil {
		log.Fatalf("Failed to update record: %v", err)
	}
	fmt.Println("Record updated!")

	// 5. Deleting a record
	existed, err := client.Delete(writePolicy, recordKey)
	if err != nil {
		log.Fatalf("Failed to delete record: %v", err)
	}
	if !existed {
		log.Fatalf("No record found with key: %v", keyValue)
	}
	fmt.Println("Record deleted!")
}
