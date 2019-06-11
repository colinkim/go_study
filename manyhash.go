package main

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"golang.org/x/crypto/sha3"
)

func main() {
	const LoopNum = 10000
	var seed string
	var out []byte = make([]byte, 32)
	var bigIntSeed = new(big.Int)

	const populationNum = 100
	var population [populationNum]uint32

	file1, _ := os.Create("data.txt")
	//fmt.Fprint(file1, 1, 1.1, "Hello, world!")
	file1.Close()

	for i := 0; i < populationNum; i++ {
		population[i] = uint32(i)
	}

	//1. 주어진 갯수 만큼 해시 만들기
	for i := 0; i < LoopNum; i++ {
		seed += strconv.Itoa(i)
		msg := []byte(seed)
		c1 := sha3.NewShake256()
		c1.Write(msg)
		c1.Read(out) //out에 넣어서 읽어야 되나 보넹..
		bigIntSeed.SetBytes(out)

		//ranNum := Uint64(bigIntSeed)
		ranNum := binary.BigEndian.Uint64(out)

		//10개중에 1개(당첨자) 뽑는다고 하면
		win := ranNum % uint64(populationNum)
		fmt.Println(population[int(win)])
		stringWin := strconv.FormatUint(uint64(win), 10)

		fread, _ := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY, 0644)
		fmt.Fprintln(fread, stringWin)

		defer fread.Close()

		//fmt.Println(bigIntSeed)
		//fmt.Println(hex.EncodeToString(out))
		//fmt.Println(out)

	}

}
