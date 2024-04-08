package main

func defangIPaddr(address string) string {
	newStr := ""
	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			newStr += "[.]"
		} else {
			newStr += string(address[i])
		}
	}
	return newStr
}

func main() {
	println(defangIPaddr("1.1.1.1"))
	println(defangIPaddr("255.100.50.0"))
}
