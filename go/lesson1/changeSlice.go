package main

import "fmt"

// src list
var tempList = [...]string{"I", "am", "stupid", "and", "weak"}

// replace map: for replace word from src list
var replaceMap = map[string]string{
	"stupid": "smart",
	"weak":   "strong",
}

func main() {
	for i := 0; i < len(tempList); i++ {
		if valReplaceTo, isExist := replaceMap[tempList[i]]; isExist {
			tempList[i] = valReplaceTo
		}
	}
	fmt.Printf("changed list: %v", tempList)
}
