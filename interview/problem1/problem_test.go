package problem1_test

import (
	"fmt"
	"testing"

	. "github.com/KINGMJ/go-algorithm/interview/problem1"
)

func TestProblem(t *testing.T) {
	randName1 := GetRandName(map[string]int{"name1": 10, "name2": 20, "name3": 50})
	randName2 := GetRandName(map[string]int{"name1": 10, "name2": 20, "name3": 50, "name4": 20})

	fmt.Println("Random name 1:", randName1)
	fmt.Println("Random name 2:", randName2)
}
