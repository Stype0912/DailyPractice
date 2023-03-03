package making_file_names_unique

import (
	"fmt"
)

func getFolderNames(names []string) []string {
	nameMap := map[string]int{}
	ans := []string{}
	for _, name := range names {
		if k, ok := nameMap[name]; !ok {
			ans = append(ans, name)
			nameMap[name] = 1
		} else {
			for i := k; ; i++ {
				newName := name + fmt.Sprintf("(%v)", i)
				if _, newOk := nameMap[newName]; !newOk {
					ans = append(ans, newName)
					nameMap[newName] = 1
					nameMap[name] = i + 1
					break
				}
			}
		}
	}
	return ans
}
