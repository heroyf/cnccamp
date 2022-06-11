package main

import (
	"fmt"
)

// src list
var tempList = []string{"I", "am", "stupid", "and", "weak"}

// replace map: for replace word from src list
var replaceMap = map[string]string{
	"stupid": "smart",
	"weak":   "strong",
}

func main() {
	finalList := replace(tempList)
	fmt.Printf("changed list: %v", finalList)
}

func replace(srcList []string) []string {
	for i := 0; i < len(srcList); i++ {
		if valReplaceTo, isExist := replaceMap[srcList[i]]; isExist {
			srcList[i] = valReplaceTo
		}
	}
	return srcList
}
