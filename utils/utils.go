package utils

//TotalSum receive array and ruturn sum of all elements
func TotalSum(array []int) (ts int) {
	for _, v := range array {
		ts += v
	}
	return
}

//CopyArray return copy of received array
func CopyArray(arr []int) []int {
	tmp := []int{}
	for _, v := range arr {
		tmp = append(tmp, v)
	}
	return tmp
}
