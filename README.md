# Assignment : MurMur Hash Implementation
Completed the assignment 1, which was to implementing a 10 character unique hash function generation service that accepts alphanumeric input which should be made modular as much as possible implementing clean code. 
For that i implemented the MurmurHash algorithm, which relies on multiplication, bit rotation, and XOR operations to achieve a well-distributed hash.

## About the algorithm
The algorithm uses c1 and c2 as prime numbers to ensure uniform distribution. A critical part of the implementation involves bitwise XOR and bit manipulation, enhancing the avalanche effect, where small input changes cause significant hash variations. Since MurmurHash naturally produces numeric output, I designed a Base62 encoding function to convert it into a 10-character alphanumeric string, ensuring uniqueness and readability.
For producing avalanche effect, ie a small change in input should contribute to a larger change in output, i have made different function which is generally not a part of murmurhash algorithm whose code is as:
```golang
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
```

[Click for more info](https://en.wikipedia.org/wiki/MurmurHash)