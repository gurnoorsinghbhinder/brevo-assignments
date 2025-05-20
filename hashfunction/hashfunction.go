package HashFunction

import (
    "errors"
    "fmt"
    //"strings"
)

const (
    c1           uint64 = 0xcc9e2d51
    c2           uint64 = 0x1b873593
    Prime1       uint64 = 0x85ebca6b
    Prime2       uint64 = 0xc2b2ae35
    Prime3       uint64 = 0xe6546b64
    base62              = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    base62Length        = 62
    seed         uint64 = 0x9747b28c
)

// Performs a left rotation on 64-bit unsigned integer
func rotateLeft(num uint64, k int) uint64 {
    return (num << k) | (num >> (64 - k))
}

// Converts 64-bit hash to base62 encoded string
func convertBase62(hash uint64, length int, trimLeadingZeros bool) string {
    var ansi string
    for hash > 0 { // Only process non-zero hash values
        keyLength := hash % base62Length
        ansi = string(base62[keyLength]) + ansi
        hash /= base62Length}
    
    return ansi
}

// Error if input key is empty
func checkEmptyKey(key string, keyLength int) error {
    if keyLength == 0 {
        return errors.New("empty key! please try again!")
    }
    return nil
}

// Processes full 8-byte chunks
func handleInitialBits(key string, keyLength, i int, hash uint64) (uint64, int) {
    for ; i+8 <= keyLength; i += 8 {
        curr := uint64(key[i]) | uint64(key[i+1])<<8 | uint64(key[i+2])<<16 | uint64(key[i+3])<<24 |
            uint64(key[i+4])<<32 | uint64(key[i+5])<<40 | uint64(key[i+6])<<48 | uint64(key[i+7])<<56
        curr *= c1
        curr = rotateLeft(curr, 31)
        curr *= c2
        hash ^= curr
        hash = rotateLeft(hash, 27)
        hash = hash*5 + Prime3
    }
    return hash, i
}

// Processes leftover bytes
func handleRemainingBits(key string, keyLength, i int, hash uint64) (uint64, int) {
    var remainingBytes uint64
    switch keyLength % 8 {
    case 7:
        remainingBytes |= uint64(key[i+6]) << 48
        fallthrough
    case 6:
        remainingBytes |= uint64(key[i+5]) << 40
        fallthrough
    case 5:
        remainingBytes |= uint64(key[i+4]) << 32
        fallthrough
    case 4:
        remainingBytes |= uint64(key[i+3]) << 24
        fallthrough
    case 3:
        remainingBytes |= uint64(key[i+2]) << 16
        fallthrough
    case 2:
        remainingBytes |= uint64(key[i+1]) << 8
        fallthrough
    case 1:
        remainingBytes |= uint64(key[i])
        remainingBytes *= c1
        remainingBytes = rotateLeft(remainingBytes, 31)
        remainingBytes *= c2
        hash ^= remainingBytes
    }
    return hash, i
}

// Final bit-mixing (avalanche effect)
func handleAvalancheEffect(keyLength int, hash uint64) uint64 {
    hash ^= uint64(keyLength)

    // More aggressive bit mixing
    hash ^= hash >> 33
    hash *= Prime1
    hash ^= hash >> 29
    hash *= Prime2
    hash ^= hash >> 33

    // Add additional mixing rounds
    hash ^= hash >> 15
    hash *= 0x9e3779b97f4a7c15 // Another prime constant (golden ratio for 64-bit)
    hash ^= hash >> 21

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

func printHash(str ...string){
    for _,val:=range str{
        ansi,err:=hashOf(val)
        if err!=nil{
            fmt.Println("error in calculating hash of: ",ansi," with Err:",err)
        }else{
            fmt.Println("Hash of ",val," is: ",ansi)
        }

    }
    fmt.Println()
} 

func HashFunction() {
 
    //same examples
    printHash("harshit","harshit")

    //Slightly different
    printHash("harshit a","harshit b")

    //Permutated examples:
    printHash("1234","1243","123")

    //Edge cases 
    printHash("","a")

    //Additional Edge Cases
    printHash("a very long string that exceeds normal length for testing purposes","!@#$%^&*()_+","!@#$%^&*()_+","   ")

}