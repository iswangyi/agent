package utils

import (
	"fmt"
	"sort"
	"strings"
)

func SortedTags(tag map[string]string) string {
	if tag == nil {
		return ""
	}
	size := len(tag)
	if size == 1 {
		for k, v := range tag {
			return fmt.Sprintf("%s=%s", k, v)
		}
	}
	var key []string

	for k := range tag {
		fmt.Println(k)
		key = append(key, k)
	}

	sort.Strings(key)

	var ret []string

	for _, v := range key {
		ret = append(ret, fmt.Sprintf("%s=%s", v, tag[v]))
	}
	return strings.Join(ret, ",")
}
