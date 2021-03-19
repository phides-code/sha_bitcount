/*write a function that counts the number of bits that are different in two SHA256 hashes. see PopCount. */
package main
import (
    "fmt"
    "crypto/sha256"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var i int
	for i = 0; x != 0; i++ {
		x = x&(x-1)
	}
	
	return i
}

func main() {
    //get the 2 hashes (arrays of bytes) [32]byte
    h1 := sha256.Sum256([]byte("x"))
    h2 := sha256.Sum256([]byte("X"))

    fmt.Printf("Number of different bits: %d\n", countDiffBits(h1[:], h2[:]))
}

func countDiffBits(h1, h2 []byte) int {
    var totalDiffBits int

    // for each BYTE
    for i := range h1 {
        
        xorResult := (h1[i])^(h2[i]) // find different bits

        fmt.Printf("Compare %08b with %08b ... xorResult is %08b. ", h1[i], h2[i], xorResult)
        
        diffBits := PopCount(uint64(xorResult)) // count different bits
        fmt.Printf("diffBits are %d\n", diffBits)
        
        totalDiffBits += diffBits // add different bits to running total
    }

    return totalDiffBits
}
