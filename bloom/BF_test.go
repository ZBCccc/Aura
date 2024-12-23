package bloom

import (
	"slices"
	"testing"
)

func TestBf(t *testing.T) {
	filter := New(10000, 10)
	filter.Add([]byte("Love"))
	if filter.Test([]byte("Love")) {
		t.Log("元素存在")
	} else {
		t.Log("元素不存在")
	}

	// test get index
	a := filter.GetIndex([]byte("Love"))
	t.Log(a)
	slices.Sort(a)

	// test search
	b := filter.Search()
	t.Log(b)

	// Verify the index and the bit set

}
