package commands

import (
	"fmt"
	"sort"
)

func GetSortedMapKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

func Show(array []string, nameOfShowed string) {
	if len(array) == 0 {
		fmt.Printf("No %s to show\n\n", nameOfShowed)
		return
	}

	fmt.Printf("Available %s\n\n", nameOfShowed)
	counter := 1

	for _, url := range array {
		fmt.Println(counter, url)
		counter++
	}
}
