// Package for normalizing string
package strNorm

// Normalize return string with one space between words
func Normalize (s string) string {
	var res string
	spaceFlag := false
	for _, char := range s {
		if char == 32 {
			if !spaceFlag {
				res += string(char)
				spaceFlag = true
			}
		} else {
			res += string(char)
			spaceFlag = false
		}
	}
	return res
}