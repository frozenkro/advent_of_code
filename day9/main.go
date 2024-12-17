package main

import "strconv"

func main() {

}

func DefragChksum(data []byte) int {
	dataMap := createDataMap(data)
	_ = dataMap
	return 0
}

func createDataMap(src []byte) []byte {
	dataMap := []byte{}
	isEven := true
	for i, v := range src {
		var ch byte
		if isEven {
			ch = '.'
		} else {
			ch = byte(i / 2)
		}

		intV, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}
		frag := createFrag(ch, intV)

		dataMap = append(dataMap, frag...)
		isEven = !isEven
	}
	return dataMap
}

func createFrag(ch byte, ln int) []byte {
	f := make([]byte, ln)
	for i := 0; i < ln; i++ {
		f[i] = ch
	}
	return f
}
