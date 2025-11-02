# go-anagram

[![Go Reference](https://pkg.go.dev/badge/github.com/abdihaikal/go-anagram.svg)](https://pkg.go.dev/github.com/abdihaikal/go-anagram)
[![Codecov](https://codecov.io/gh/abdihaikal/go-anagram/branch/development/graph/badge.svg)](https://codecov.io/gh/abdihaikal/go-anagram)
[![Go Report Card](https://goreportcard.com/badge/github.com/abdihaikal/go-anagram)](https://goreportcard.com/report/github.com/abdihaikal/go-anagram)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

A Go package for checking if two strings are anagrams of each other.

## Description

The `go-anagram` package provides a simple utility function to determine if two strings are anagrams. An anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

This implementation:
- Ignores case differences
- Ignores whitespace characters
- Handles Unicode characters properly

## Installation

```bash
go get github.com/abdihaikal/go-anagram
```

## Usage

### As a Package

```go
package main

import (
    "fmt"
    "github.com/abdihaikal/go-anagram/anagram"
)

func main() {
    fmt.Println(anagram.IsAnagram("listen", "silent"))         // true
    fmt.Println(anagram.IsAnagram("Debit Card", "Bad Credit")) // true
    fmt.Println(anagram.IsAnagram("hello", "world"))           // false
}
```

### Function Signature

```go
func IsAnagram(val1, val2 string) bool
```

Returns `true` if the two strings are anagrams of each other, `false` otherwise.

## Examples

```go
// Simple anagrams
anagram.IsAnagram("listen", "silent") // true

// Case insensitive
anagram.IsAnagram("Listen", "Silent") // true

// Whitespace ignored
anagram.IsAnagram("Debit Card", "Bad Credit") // true

// Not anagrams
anagram.IsAnagram("hello", "world") // false

// Unicode support
anagram.IsAnagram("a gentleman", "elegant man") // true
```

## Testing

To run the tests:

```bash
go test ./...
```

Or with verbose output:

```bash
go test -v ./...
```

## Implementation Details

The algorithm uses a frequency map approach:
1. First checks if the strings have the same length (after ignoring whitespace)
2. Creates a frequency map of characters from the first string
3. Decrements the frequency for each character in the second string
4. If any frequency goes negative, the strings are not anagrams

Time complexity: O(n) where n is the length of the strings
Space complexity: O(k) where k is the number of unique characters

## License

This project is licensed under the Apache License, Version 2.0 - see the [LICENSE](LICENSE) file for details.

## Author

Abdi Haikal