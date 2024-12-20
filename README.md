# AURA: The pure Go implementation of SSE

The go implementation of Aura, cite from the paper [Practical Non-Interactive Searchable Encryption with Forward and Backward Privacy](https://www.ndss-symposium.org/wp-content/uploads/ndss2021_2C-4_24162_paper.pdf).

## Features

- Secure search over encrypted data
- Dynamic updates (insert and delete operations)
- Bloom filter optimization for deletion
- GGM tree based key derivation
- AES encryption for data protection

## Architecture

The system consists of three main components:

- **SSE Client**: Handles encryption, search token generation and update operations
- **SSE Server**: Stores encrypted index and performs search operations
- **GGM Tree**: Provides efficient key derivation mechanism

## Installation
To Use AURA, make sure you have Go 1.21 or later installed, then run:
```bash
go get github.com/ZBCccc/Aura
```

## Usage

1. **Initialize the Client**: Set up the SSE client to handle encryption and token generation.
2. **Perform Searches**: Use the client to generate search tokens and query the SSE server.
3. **Update Data**: Insert or delete data dynamically using the client.

```go
import "github.com/ZBCccc/aura/Core/SSEClient"
// Create a new SSE client
client := sseclient.NewSSEClient()
// Insert a document
client.Update(util.Insert, "keyword", "document1")
// Search for documents
results := client.Search("keyword")
// Delete a document
client.Update(util.Delete, "keyword", "document1")
```