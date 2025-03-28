package main

import "fmt"

const (
	c1 uint32 = 0xcc9e2d51
	c2 uint32 = 0x1b873593
)

var seed uint32 = 0x9747b28c

func rotateLeft(num uint32, k int) uint32 {
	return (num << k) | (num >> (32 - k))
}

func convertBase62(hash uint32) string {
	Base62 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var ansi string
	for i := 0; i < 10; i++ {
		n := hash % 62
		ansi = string(Base62[n]) + ansi
		hash /= 62
	}
	return ansi
}

func hashOf(key string) string {
	hash := seed
	n := len(key)

	i := 0 //so that left bits can be processed later
	for ; i+4 <= n; i += 4 {
		curr := uint32(key[i]) | uint32(key[i+1])<<8 | uint32(key[i+2])<<16 | uint32(key[i+3])<<24
		curr *= c1
		curr = rotateLeft(curr, 15)
		curr *= c2
		hash ^= curr
		hash = rotateLeft(hash, 13)
		hash = hash*5 + 0xe6546b64
	}

	//handling final bits which are not processed above
	var tail uint32 = 0
	switch n % 4 {
	case 3:
		tail |= uint32(key[i+2]) << 16
		fallthrough
	case 2:
		tail |= uint32(key[i+1]) << 8
		fallthrough
	case 1:
		tail |= uint32(key[i])
		tail *= c1
		tail = rotateLeft(tail, 15)
		tail *= c2
		hash ^= tail
		//we do not used last two steps of above loop to make it over scattered so dont use it
	}

	//avalanche effect
	hash ^= uint32(n)
	hash ^= hash >> 16
	hash *= 0x85ebca6b //random prime no to increase uniqueness
	hash ^= hash >> 13
	hash *= 0xc2b2ae35 //random prime no
	hash ^= hash >> 16

	return convertBase62(hash)
}

func main() {
	fmt.Println("Same examples:")
	fmt.Println(hashOf("harshit"))
	fmt.Println(hashOf("harshit"))
	fmt.Println()
	fmt.Println("slightly different:")
	fmt.Println(hashOf("harshit a"))
	fmt.Println(hashOf("harshit b"))
	fmt.Println()
	fmt.Println("permutated examples:")
	fmt.Println(hashOf("1234"))
	fmt.Println(hashOf("1243"))
	fmt.Println(hashOf("123"))

}
