package main

import (
	"fmt"
	"errors"
)

const (
	c1 uint32 = 0xcc9e2d51
	c2 uint32 = 0x1b873593
	Prime1 = 0x85ebca6b
    Prime2 = 0xc2b2ae35
    Prime3 = 0xe6546b64
)

var seed uint32 = 0x9747b28c

//this function performs a left rotation on 32-bit unsigned integer.
func rotateLeft(num uint32, k int) uint32 {
	return (num << k) | (num >> (32 - k))
}

func convertBase62(hash uint32) string {
	const Base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	const Base62Length=62
	var ansi string
	for i := 0; i < 10; i++ {
		keyLength := hash % Base62Length
		ansi = string(Base62[keyLength]) + ansi
		hash /= Base62Length
	}
	return ansi
}

func hashOf(key string) (string,error) {
	keyLength := len(key)
    if keyLength==0{
        err:=errors.New("Empty Key! Please Try Again !")
		return "",err
	}

	hash := seed
	

	i := 0 //so that left bits can be processed later
	for ; i+4 <= keyLength; i += 4 {
		curr := uint32(key[i]) | uint32(key[i+1])<<8 | uint32(key[i+2])<<16 | uint32(key[i+3])<<24
		curr *= c1
		curr = rotateLeft(curr, 15)
		curr *= c2
		hash ^= curr
		hash = rotateLeft(hash, 13)
		hash = hash*5 + Prime3
	}

	//handling final bits which are not processed above
	var remainingBytes uint32 = 0
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
		//we do not used last two steps of above loop to make it over scattered so dont use it
	}

	//avalanche effect
	hash ^= uint32(keyLength)
	hash ^= hash >> 16
	hash *= Prime1 //random prime no to increase uniqueness
	hash ^= hash >> 13
	hash *= Prime2 //random prime no
	hash ^= hash >> 16

	return convertBase62(hash),nil
}

func main() {
	//testing identical inputs
	fmt.Println("Same examples:")
	if hash,err:=hashOf("harshit");err==nil{
		fmt.Println(hash)
	} else{
		fmt.Println(err)
	}
	if hash,err:=hashOf("harshit");err==nil{
		fmt.Println(hash)
	} else{
		fmt.Println(err)
	}
	
	//testing avalache effect  
	fmt.Println()
	fmt.Println("slightly different:")
	if hash,err:=hashOf("harshit a");err==nil{
		fmt.Println(hash)
	} else {
		fmt.Println(err)
	}
	if hash,err:=hashOf("harshit b");err==nil{
		fmt.Println(hash)
	} else {
		fmt.Println(err)
	}
	
	//testing permutations 
	fmt.Println()
	fmt.Println("permutated examples:")
	if hash,err:=hashOf("1234");err==nil{
		fmt.Println(hash)
	} else{
		fmt.Println(err)
	}
	if hash,err:=hashOf("1243");err==nil{
		fmt.Println(hash)
	} else{
		fmt.Println(err)
	}
	if hash,err:=hashOf("123");err==nil{
		fmt.Println(hash)
	} else{
		fmt.Println(err)
	}
	
	//testing edge cases
    fmt.Println()
	fmt.Println("Edge Cases:")
	if hash,err:=hashOf("");err==nil{
		fmt.Println(hash)
	} else{
		fmt.Println(err)
	}

	//additional test cases
	fmt.Println()
	fmt.Println("Additional Edge Cases:")
longString := "a very long string that exceeds normal length for testing purposes"
if hash,err:=hashOf(longString);err==nil{
	fmt.Println(hash)
} else{
	fmt.Println(err)
}

specialChars := "!@#$%^&*()_+"
if hash,err:=hashOf(specialChars);err==nil{
	fmt.Println(hash)
} else{
	fmt.Println(err)
}

whitespace := "     "
if hash,err:=hashOf(whitespace);err==nil{
	fmt.Println(hash)
} else{
	fmt.Println(err)
}
}