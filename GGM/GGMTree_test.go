package ggmtree

import (
	"testing"
)

func TestNewGGMTree(t *testing.T) {
	tests := []struct {
		numNode int64
		expectedLevel int
	}{
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 2},
		{8, 3},
	}

	for _, test := range tests {
		tree := NewGGMTree(test.numNode)
		if tree.Level != test.expectedLevel {
			t.Errorf("NewGGMTree(%d) = %d; want %d", test.numNode, tree.Level, test.expectedLevel)
		}
	}
}

func TestDeriveKeyFromTree(t *testing.T) {

	tests := []struct {
		currentKey   []byte
		offset       uint
		startLevel   int
		targetLevel  int
	}{
		{[]byte{0x00}, 0, 2, 0},
		{[]byte{0x01}, 1, 2, 0},
	}

	for _, test := range tests {
		originalKey := make([]byte, len(test.currentKey))
		copy(originalKey, test.currentKey)
		DeriveKeyFromTree(test.currentKey, test.offset, test.startLevel, test.targetLevel)
		t.Log(test.currentKey)
	}
} 