package main

// 13/19 passed
import "math/rand"

type RandomizedSet struct {
	Data map[int]int
	Keys []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		Data: make(map[int]int),
		Keys: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.Data[val]; ok {
		return false
	}
	this.Keys = append(this.Keys, val)
	this.Data[val] = len(this.Keys) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.Data[val]; ok {
		index := this.Data[val]
		oldArr := this.Keys
		this.Keys = oldArr[0:index]
		this.Keys = append(this.Keys, oldArr[index:]...)
		delete(this.Data, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	random := rand.Intn(len(this.Keys))
	return this.Keys[random]
}

func main() {

}
