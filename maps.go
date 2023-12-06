package main

func GetKey(m map[int]int) int {
	for k := range m {
		return k
	}
	return 0
}
