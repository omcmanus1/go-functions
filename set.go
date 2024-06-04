package main

// implement set data structure in Go
// (not available natively)
func Set(duplicatesArr []string) []string {
	set := make(map[string]struct{})
	for _, item := range duplicatesArr {
		set[item] = struct{}{}
	}
	arr := make([]string, 0, len(set))
	for key := range set {
		arr = append(arr, key)
	}
	return arr
}
