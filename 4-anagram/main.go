package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}

	var allMaps []map[string]int
	var uniqueMaps []map[string]int

	for _, word := range words {
		tempMap := make(map[string]int)

		for _, letter := range strings.ToLower(word) {
			tempMap[string(letter)] += 1
		}

		allMaps = append(allMaps, tempMap)

		if len(uniqueMaps) == 0 {
			uniqueMaps = append(uniqueMaps, tempMap)
			continue
		}

		var unique = true

		for _, mp := range uniqueMaps {
			var matched = true

			for key := range mp {
				if mp[key] != tempMap[key] {
					matched = false
					break
				}
			}

			if matched {
				unique = false
				break
			}
		}

		if unique {
			uniqueMaps = append(uniqueMaps, tempMap)
		}
	}

	wordsGroup := make([][]string, len(uniqueMaps))

	for indWord, notUniqueMap := range allMaps {

		for ind, uniqueMap := range uniqueMaps {

			var matched = true

			for key := range uniqueMap {

				if notUniqueMap[key] != uniqueMap[key] {
					matched = false
					break
				}
			}

			if matched {
				wordsGroup[ind] = append(wordsGroup[ind], words[indWord])
				break
			}
		}
	}

	for _, slc := range wordsGroup {
		fmt.Println(slc)
	}
}
