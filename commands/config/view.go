package config

import (
	"fmt"
	"sort"
)

func ShowCommandConfig(cfg CommandConfig, nameOfShowedEntry string, verbose bool) {
	if len(cfg) == 0 {
		fmt.Printf("No %s to show\n\n", nameOfShowedEntry)
		return
	}

	keys := make([]string, 0, len(cfg))
	for k := range cfg {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	fmt.Printf("Available %s\n\n", nameOfShowedEntry)

	counter := 1
	for _, k := range keys {
		if verbose {
			v := cfg[k]
			fmt.Println(counter, k, v)
		} else {
			fmt.Println(counter, k)
		}
		counter++
	}
}
