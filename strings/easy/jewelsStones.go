package main

func numJewelsInStones(jewels string, stones string) int {
	jm := make(map[rune]struct{})
	for _, j := range jewels {
		jm[j] = struct{}{}
	}
	jewelsCount := 0
	for _, s := range stones {
		if _, ok := jm[s]; ok {
			jewelsCount++
		}
	}
	return jewelsCount
}

func main() {
	println(numJewelsInStones("aA", "aAAbbbb"))
	println(numJewelsInStones("z", "ZZ"))
	println(numJewelsInStones("z", "ZZz"))
}
