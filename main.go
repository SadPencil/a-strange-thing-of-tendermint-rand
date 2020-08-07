package main

import (
	"fmt"
	tmRand "github.com/tendermint/tendermint/libs/rand"
	mathRand "math/rand"
	"runtime"
	"runtime/debug"
)

func f1() {
	for i := 0; ; i++ {
		str := tmRand.Str(tmRand.Intn(32768))
		fmt.Printf("tick-tock %d %p\n", i, &str)

		runtime.GC()
		debug.FreeOSMemory()
	}
}

func f2() {
	for i := 0; ; i++ {
		str := tmRand.Str(32768)
		fmt.Printf("tick-tock %d %p\n", i, &str)

		runtime.GC()
		debug.FreeOSMemory()
	}
}

func f3() {
	for i := 0; ; i++ {
		randBytes := make([]byte, mathRand.Intn(32768))
		_, err := mathRand.Read(randBytes)
		if err != nil {
			panic(err)
		}
		str := string(randBytes)

		fmt.Printf("tick-tock %d %p\n", i, &str)

		runtime.GC()
		debug.FreeOSMemory()
	}
}

func main() {
	f1()
}
