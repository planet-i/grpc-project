package utils

import "sort"

func IsStringInSlice(str string, strSlice []string) bool {
	sort.Strings(strSlice)
	index := sort.SearchStrings(strSlice, str)
	if index < len(strSlice) && strSlice[index] == str {
		return true
	}
	return false
}
