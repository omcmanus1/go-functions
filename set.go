package main

// implement set data structure in Go
// (not available natively)
func Set() map[string]struct{} {
	duplicatesArr := []string{"hello", "hello", "hi", "hi", "not hi", "still not hi"}
	set := map[string]struct{}{}
	for _, val := range duplicatesArr {
		set[val] = struct{}{}
	}
	return set
}
