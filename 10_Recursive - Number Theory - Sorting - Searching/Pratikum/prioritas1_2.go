package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Item  string
	Count int
}

func CountItem(items []string) []Pair {
	itemCounts := make(map[string]int)
	for _, item := range items {
		if _, ok := itemCounts[item]; ok {
			itemCounts[item]++
		} else {
			itemCounts[item] = 1
		}
	}

	pairs := make([]Pair, len(itemCounts))
	i := 0
	for k, v := range itemCounts {
		pairs[i] = Pair{Item: k, Count: v}
		i++
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Count > pairs[j].Count
	})

	return pairs

}

func main() {
	fmt.Println(CountItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// [{js 4} {ruby 2} {golang 1}]
	fmt.Println(CountItem([]string{"B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// [{A 4} {B 3} {D 2} {C 1}]
	fmt.Println(CountItem([]string{"football", "basketball", "tenis"}))
	// [{football 1} {basketball 1} {tenis 1}]
}
