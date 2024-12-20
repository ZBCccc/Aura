package sseclient

import (
	sseserver "aura/Core/SSEServer"
	ggmtree "aura/GGM"
	util "aura/Util"
	"encoding/base64"
	"fmt"
	"log"
	"slices"

	"github.com/bits-and-blooms/bloom/v3"
)

type SSEClient struct {
	key    []byte
	iv     []byte
	tree   *ggmtree.GGMTree
	bf     *bloom.BloomFilter
	C      map[string]int
	server *sseserver.Server
}

// NewSSEClient creates a new SSEClient.
func NewSSEClient() *SSEClient {
	return &SSEClient{
		key:    []byte("0123456789123456"),
		iv:     []byte("0123456789123456"),
		tree:   ggmtree.NewGGMTree(util.GGMSize),
		bf:     bloom.New(util.GGMSize, util.HashSize),
		C:      make(map[string]int),
		server: sseserver.NewServer(),
	}
}

// Update updates the client with the given data.
func (c *SSEClient) Update(op util.Operation, keyword string, ind string) {
	// compute the tag and generate the digest of tag
	pair := []byte(keyword + ind)
	tag := util.Sha256Digest(pair)

	// process the operator
	if op == util.Insert {
		// get all offsets in BF
		indexes := c.bf.GetIndex(tag)
		slices.Sort(indexes)

		// get SRE ciphertext list
		ciphertexts := make([]string, len(indexes))
		for i, index := range indexes {
			// derive a key from the offset
			derivedKey := make([]byte, len(c.key))
			copy(derivedKey, c.key)
			ggmtree.DeriveKeyFromTree(derivedKey, index, c.tree.GetLevel(), 0)

			// use the key to encrypt the id
			encryptedId, err := util.AesEncrypt([]byte(ind), derivedKey, c.iv)
			if err != nil {
				log.Fatal("Failed to AesEncrypt", err)
			}

			// save the encrypted id in the list
			ciphertexts[i] = base64.StdEncoding.EncodeToString(encryptedId)
		}

		// token
		token := util.HmacDigest([]byte(keyword), c.key)

		// label
		if _, exists := c.C[keyword]; !exists {
			c.C[keyword] = 0
		}
		label := util.HmacDigest([]byte(fmt.Sprintf("%d", c.C[keyword])), token)

		// update the counter
		c.C[keyword]++

		//save the list on the server
		c.server.AddEntries(base64.StdEncoding.EncodeToString(label), base64.StdEncoding.EncodeToString(tag), ciphertexts)
	} else if op == util.Delete {
		// insert the tag into BF
		c.bf.Add([]byte(tag))
	}
}

func (c *SSEClient) Search(keyword string) []string {
	// token
	token := util.HmacDigest([]byte(keyword), c.key)

	// search all deleted positions
	bfPos := make([]int, util.GGMSize)
	for i := 0; i < util.GGMSize; i++ {
		bfPos[i] = i
	}
	deletePos := c.bf.Search()
	remainPos := setDifference(bfPos, deletePos)

	// generate GGM Node for the remain position
	nodeList := make([]ggmtree.GGMNode, len(remainPos))
	for i, pos := range remainPos {
		nodeList[i] = ggmtree.GGMNode{Index: pos, Level: c.tree.GetLevel()}
	}
	remainNode := ggmtree.MinCoverage(nodeList)

	// compute the key set and send to the server
	for i := range remainNode {
		nodeKey := make([]byte, len(c.key))
		copy(nodeKey, c.key)
		ggmtree.DeriveKeyFromTree(nodeKey, uint(remainNode[i].Index), remainNode[i].Level, 0)
		copy(remainNode[i].Key[:], nodeKey)
	}

	// give all results to the server for search
	res := c.server.Search(token, remainNode, c.tree.GetLevel())

	// remove duplicates
	seen := make(map[string]bool)
	unique := make([]string, 0)
	for _, v := range res {
		if !seen[v] {
			seen[v] = true
			unique = append(unique, v)
		}
	}
	res = unique
	return res
}

func setDifference(a, b []int) []int {
	mb := make(map[int]bool, len(b))
	for _, x := range b {
		mb[x] = true
	}
	diff := make([]int, 0, len(a))
	for _, x := range a {
		if !mb[x] {
			diff = append(diff, x)
		}
	}
	return diff
}
