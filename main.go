package main

import (
	"log"
	"sort"
)

type SortArray []int

func (s SortArray) Len() int           { return len(s) }
func (s SortArray) Less(i, j int) bool { return s[i] < s[j] }
func (s SortArray) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	var arr SortArray
	for i := 100; i > 0; i-- {
		arr = append(arr, i)
	}
	sort.Sort(arr)
	log.Println(arr)
}
