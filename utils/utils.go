package utils

// MergeSlices - function to merged slice of bytes recieved from multiple api
func MergeSlices(s1, s2 []byte) []byte {
	s1 = s1[0 : len(s1)-1]
	s1 = append(s1, ',')
	s2 = s2[1:]
	res := append(s1, s2...)

	return res
}
