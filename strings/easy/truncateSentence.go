package main

func truncateSentence(s string, k int) string {
	words := []string{}
	i, j := 0, 0
	for j < len(s) && len(words) < k {
		if string(s[j]) == string(" ") {
			words = append(words, s[i:j])
			i = j
		}
		if j == len(s)-1 {
			words = append(words, s[i:j+1])
		}
		j++
	}
	resp := ""
	for _, w := range words {
		resp += w
	}
	return resp
}
func truncateSentence2(s string, k int) string {
	wTtrack, wti := 0, 0
	j := 0
	for j < len(s) && wTtrack < k {
		if string(s[j]) == string(" ") {
			wTtrack++
			wti = j
		}
		if j == len(s)-1 {
			wTtrack++
			wti = j + 1
		}
		j++
	}
	return s[:wti]
}
func main() {
	println(truncateSentence2("Hello how are you Contestant", 4))         // "Hello how are you"
	println(truncateSentence2("What is the solution to this problem", 4)) // "What is the solution"
	println(truncateSentence2("chopper is not a tanuki", 5))              // "chopper is not a tanuki"
}
