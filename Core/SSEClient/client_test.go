package sseclient

import (
	util "aura/Util"
	"fmt"
	"testing"
	"time"
)

func TestSSEClient(t *testing.T) {
	sseClient := NewSSEClient()

	// insert
	start := time.Now()
	for i := 0; i < 200; i++ {
		sseClient.Update(util.Insert, "test", fmt.Sprintf("%d", i))
	}
	elapsed := time.Since(start)
	t.Logf("Time taken to insert 200 items: %v", elapsed)

	// search
	start = time.Now()
	indexes := sseClient.Search("test")
	elapsed = time.Since(start)
	t.Logf("Time taken to search 200 items: %v", elapsed)
	t.Log(indexes)

	// delete
	start = time.Now()
	for i := 100; i < 200; i++ {
		sseClient.Update(util.Delete, "test", fmt.Sprintf("%d", i))
	}
	elapsed = time.Since(start)
	t.Logf("Time taken to delete 100 items: %v", elapsed)

	// search
	start = time.Now()
	indexes = sseClient.Search("test")
	elapsed = time.Since(start)
	t.Logf("Time taken to search 100 items: %v", elapsed)
	t.Log(indexes)
}
