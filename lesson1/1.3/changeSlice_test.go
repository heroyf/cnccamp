package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestChangeSlice(t *testing.T) {
	expected := []string{"I", "am", "smart", "and", "strong"}
	t.Log("start test change slice")
	tempList := []string{"I", "am", "stupid", "and", "weak"}
	finalList := replace(tempList)
	equal := reflect.DeepEqual(finalList, expected)
	if !equal {
		err := errors.New("change slice test is not passed")
		if err != nil {
			return
		}
	}
}
