package main

import (
	"fmt"
	"time"

	sseclient "github.com/ZBCccc/Aura/Core/SSEClient"
	util "github.com/ZBCccc/Aura/Util"
)

func main() {
	sseClient := sseclient.NewSSEClient()

	// insert
	start := time.Now()
	for i := 0; i < 200; i++ {
		sseClient.Update(util.Insert, "test", fmt.Sprintf("%d", i))
		sseClient.Update(util.Insert, "aura", fmt.Sprintf("%d", i))
	}
	elapsed := time.Since(start)
	fmt.Printf("Time taken to insert 200 items: %v\n", elapsed)

	// search
	start = time.Now()
	indexes := sseClient.Search("test")
	elapsed = time.Since(start)
	fmt.Printf("Time taken to search 200 items: %v\n", elapsed)
	fmt.Println(indexes[:10])

	// delete
	start = time.Now()
	for i := 100; i < 200; i++ {
		sseClient.Update(util.Delete, "test", fmt.Sprintf("%d", i))
	}
	elapsed = time.Since(start)
	fmt.Printf("Time taken to delete 100 items: %v\n", elapsed)

	// search
	start = time.Now()
	indexes = sseClient.Search("test")
	elapsed = time.Since(start)
	fmt.Printf("Time taken to search 100 items: %v\n", elapsed)
	fmt.Println(indexes)
	indexes = sseClient.Search("aura")
	fmt.Println(indexes)
}
