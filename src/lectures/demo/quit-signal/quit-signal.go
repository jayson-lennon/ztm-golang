package main

import (
	"fmt"
	"hash/crc32"
)

func hashStrings(quit chan struct{}, data ...string) <-chan uint32 {
	out := make(chan uint32, 1)

	go func() {
		defer close(out)

		for _, d := range data {
			bytes := []byte(d)
			hash := crc32.ChecksumIEEE((bytes))
			select {
			case out <- hash:
			case <-quit:
				return
			}
		}
	}()

	return out
}

func main() {
	data := []string{"hello", "goodbye", "hash", "go", "ztm", "a", "b", "c", "d"}

	quit := make(chan struct{})

	hashes := hashStrings(quit, data...)

	fmt.Println("hashes:")
	i := 0
	for hash := range hashes {
		fmt.Printf("0x%08x\n", hash)
		i += 1
		// abort after printing 2 hashes
		if i == 2 {
			// close the `quit` channel so the `hashStrings` goroutine
			// stops generating hashes
			close(quit)
			break
		}
	}

	// There may still be data in the channel which was added in the time
	// between processing and the `quit` signal
	fmt.Println("\nleftover hashes:")
	for hash := range hashes {
		fmt.Printf("0x%08x\n", hash)
	}

}
