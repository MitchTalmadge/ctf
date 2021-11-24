package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math"

	"github.com/x448/float16"
)

func main() {
	encodedData := demodulate("signal.ham")

	/* for i := 2; i < 100; i++ {
		for j := 0; j <= 1; j++ {
			extended := false
			if j == 1 {
				extended = true
			}
			decodedData := decode(encodedData, i, extended)

			fmt.Printf("(%d/%v) '%s'", i, extended, decodedData[:3])
		}
	} */

	decodedData := decode(encodedData, 3, false)

	ioutil.WriteFile("signal.ham.dec", decodedData, 0644)
}

func demodulate(fileName string) (demodulatedData []byte) {
	modulatedData, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	if len(modulatedData)%2 != 0 {
		panic("Expected modulated file to have even number of bytes.")
	}

	numDemodulatedBits := len(modulatedData) / 2
	numDemodulatedBytes := int(math.Ceil(float64(numDemodulatedBits) / 8.0))
	demodulatedData = make([]byte, numDemodulatedBytes, numDemodulatedBytes)
	for i := 0; i < numDemodulatedBytes; i++ {
		demodulatedData[i] = 0
	}
	fmt.Printf("Demodulating to %d bits across %d bytes\n", numDemodulatedBits, numDemodulatedBytes)

	// Two bytes at a time, for float16
	for i := 0; i < len(modulatedData); i += 2 {
		currentBit := i / 2
		currentByte := currentBit / 8
		bitPosition := currentBit % 8

		sampleBytes := modulatedData[i : i+2]
		sampleFloat := float16.Frombits(binary.LittleEndian.Uint16(sampleBytes))

		// Demodulate
		var b byte = 0
		if sampleFloat.Float32() >= 0 {
			b = 1
		}

		//fmt.Printf("%d", b)

		// Store the demodulated bit.
		demodulatedData[currentByte] = storeBit(b, bitPosition, demodulatedData[currentByte])
	}

	return demodulatedData
}

func decode(encodedData []byte, parityBits int, extended bool) (decodedData []byte) {
	totalBits := int(math.Pow(2, float64(parityBits))) - 1
	numDataBits := totalBits - parityBits

	if extended {
		parityBits++
		totalBits++
	}

	numBlocks := int(math.Ceil(float64(len(encodedData)*8) / float64(totalBits)))
	numDecodedBytes := int(math.Ceil(float64(numBlocks*numDataBits) / 8.0))
	decodedData = make([]byte, numDecodedBytes, numDecodedBytes)

	fmt.Printf("Decoding from %d bytes to %d bytes with (%d, %d) Hamming code\n", len(encodedData), numDecodedBytes, totalBits, numDataBits)

	for i := 0; i < numBlocks; i++ {
		startBit := totalBits * i
		startByteIndex := startBit / 8
		startByteOffset := startBit % 8
		endByteIndex := startByteIndex + ((startByteOffset + totalBits) / 8)
		startDstByteIndex := (numDataBits * i) / 8
		startDstByteOffset := (numDataBits * i) % 8

		var byteRange []byte
		if endByteIndex < len(encodedData) {
			byteRange = encodedData[startByteIndex : endByteIndex+1]
		} else {
			// Padding
			byteCount := (endByteIndex - startByteIndex) + 1
			byteRange = make([]byte, byteCount, byteCount)
			for j := 0; j < byteCount; j++ {
				byteRange[j] = 0
			}
			copy(byteRange, encodedData[startByteIndex:])
		}

		for j := 0; j < totalBits; j++ {
			byteRangeIndex := (startByteOffset + j) / 8
			currentByteOffset := (startByteOffset + j) % 8
			dstByteIndex := startDstByteIndex + ((startDstByteOffset + j) / 8)
			dstByteOffset := (startDstByteOffset + j) % 8
			b := (byteRange[byteRangeIndex] >> currentByteOffset) & 1
			//fmt.Printf("%d", b)

			if j < numDataBits {
				decodedData[dstByteIndex] = storeBit(b, dstByteOffset, decodedData[dstByteIndex])
			}
		}
		//fmt.Printf("\n")
	}

	return decodedData
}

func storeBit(val byte, pos int, dst byte) byte {
	if val != 0 && val != 1 {
		panic("Trying to store non-binary bit")
	}
	if pos < 0 || pos > 7 {
		panic("Trying to store bit out of range")
	}
	bitMask := byte(1) << pos
	return (dst & ^bitMask) | ((val << pos) & bitMask)
}
