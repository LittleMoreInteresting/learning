package main

func main() {

}

func strLen(str []byte) int {
	n := len(str)
	for i := n - 1; i >= 0; i-- {
		if i != n-1 && string(str[i]) == " " {
			return n - 1 - i
		}
	}
	return n
}
