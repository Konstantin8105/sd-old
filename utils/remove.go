package utils

// Remove - formula a = a - b
func Remove(a []int, b []int) []int {
AGAIN:
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				// remove
				a = append(a[:i], a[i+1:]...)
				goto AGAIN
			}
		}
	}
	return a
}
