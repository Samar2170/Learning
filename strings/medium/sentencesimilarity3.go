package main

import (
	"strings"
)

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	w1, w2 := strings.Split(sentence1, ""), strings.Split(sentence2, "")
	f1, f2 := true, true
	for _, w := range w1 {
		a := Contains(w2, w)
		if a == false {
			f1 = false
			break
		}
	}
	for _, w := range w2 {
		a := Contains(w1, w)
		if a == false {
			f2 = false
			break
		}
	}
	if f1 == true || f2 == true {
		return true
	}
	return false
}
func areSentencesSimilar2(sentence1 string, sentence2 string) bool {
	if len(sentence1) == len(sentence2) {
		return sentence1 == sentence2
	}
	if len(sentence1) <= 1 || len(sentence2) <= 1 {
		return false
	}
	var bs, ls string
	if len(sentence1) > len(sentence2) {
		bs = sentence1
		ls = sentence2
	} else {
		bs = sentence2
		ls = sentence1
	}
	w1, w2 := strings.Split(bs, " "), strings.Split(ls, " ")
	i, j, k := 0, len(w2)-1, len(w1)-1
	for i < j && i < k {
		if w2[i] != w1[i] || w2[j] != w1[k] {
			return false
		}
		i++
		j--
		k--
	}
	return true
}

func Contains(arr []string, w string) bool {
	for _, a := range arr {
		if a == w {
			return true
		}
	}
	return false
}

func main() {
	// println(areSentencesSimilar("My name is Haley", "My Haley")) // true
	// println(areSentencesSimilar("of", "A lot of words"))         // false
	// println(areSentencesSimilar("Eating right now", "Eating"))   // true

	println(areSentencesSimilar2("My name is Haley", "My Haley")) // true
	println(areSentencesSimilar2("of", "A lot of words"))         // false
	println(areSentencesSimilar2("Eating right now", "Eating"))   // true

}
