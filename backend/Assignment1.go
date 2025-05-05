package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	c1           uint32 = 0xcc9e2d51
	c2           uint32 = 0x1b873593
	Prime1       uint32 = 0x85ebca6b
	Prime2       uint32 = 0xc2b2ae35
	Prime3       uint32 = 0xe6546b64
	base62              = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	base62Length        = 62
	seed         uint32 = 0x9747b28c
)

// Performs a left rotation on 32-bit unsigned integer
func rotateLeft(num uint32, k int) uint32 {
	return (num << k) | (num >> (32 - k))
}

// Converts 32-bit hash to base62 encoded string
func convertBase62(hash uint32, length int, trimLeadingZeros bool) string {
	var ansi string
	for i := 0; i < length; i++ {
		keyLength := hash % base62Length
		ansi = string(base62[keyLength]) + ansi
		hash /= base62Length
	}
	if trimLeadingZeros {
		ansi = strings.TrimLeft(ansi, "0")
	}
	return ansi
}

// Error if input key is empty
func checkEmptyKey(key string, keyLength int) error {
	if keyLength == 0 {
		return errors.New("empty key! please try again!")
	}
	return nil
}

// Processes full 4-byte chunks
func handleInitialBits(key string, keyLength, i int, hash uint32) (uint32, int) {
	for ; i+4 <= keyLength; i += 4 {
		curr := uint32(key[i]) | uint32(key[i+1])<<8 | uint32(key[i+2])<<16 | uint32(key[i+3])<<24
		curr *= c1
		curr = rotateLeft(curr, 15)
		curr *= c2
		hash ^= curr
		hash = rotateLeft(hash, 13)
		hash = hash*5 + Prime3
	}
	return hash, i
}

// Processes leftover bytes
func handleRemainingBits(key string, keyLength, i int, hash uint32) (uint32, int) {
	var remainingBytes uint32
	switch keyLength % 4 {
	case 3:
		remainingBytes |= uint32(key[i+2]) << 16
		fallthrough
	case 2:
		remainingBytes |= uint32(key[i+1]) << 8
		fallthrough
	case 1:
		remainingBytes |= uint32(key[i])
		remainingBytes *= c1
		remainingBytes = rotateLeft(remainingBytes, 15)
		remainingBytes *= c2
		hash ^= remainingBytes
	}
	return hash, i
}

// Final bit-mixing (avalanche effect)
func handleAvalancheEffect(keyLength int, hash uint32) uint32 {
	hash ^= uint32(keyLength)
	hash ^= hash >> 16
	hash *= Prime1
	hash ^= hash >> 13
	hash *= Prime2
	hash ^= hash >> 16
	return hash
}

// Main hash function with base62 output
func hashOf(key string) (string, error) {
	keyLength := len(key)
	if err := checkEmptyKey(key, keyLength); err != nil {
		return "", err
	}

	hash := seed
	i := 0
	hash, i = handleInitialBits(key, keyLength, i, hash)
	hash, i = handleRemainingBits(key, keyLength, i, hash)
	hash = handleAvalancheEffect(keyLength, hash)

	return convertBase62(hash, 10, true), nil
}

func main() {
	fmt.Println("Same examples:")
	fmt.Println(hashOf("harshit"))
	fmt.Println(hashOf("harshit"))

	fmt.Println("\nSlightly different:")
	fmt.Println(hashOf("harshit a"))
	fmt.Println(hashOf("harshit b"))

	fmt.Println("\nPermutated examples:")
	fmt.Println(hashOf("1234"))
	fmt.Println(hashOf("1243"))
	fmt.Println(hashOf("123"))

	fmt.Println("\nEdge Cases:")
	fmt.Println(hashOf(""))

	fmt.Println("\nAdditional Edge Cases:")
	fmt.Println(hashOf("a very long string that exceeds normal length for testing purposes"))
	fmt.Println(hashOf("!@#$%^&*()_+"))
	fmt.Println(hashOf("     "))
}
