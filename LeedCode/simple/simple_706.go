package simple

//设计哈希函数

var base = 1000

type MyHashMap struct {
	Array [1000][]*MyMap
}

type MyMap struct {
	Key   int
	Value int
}

func Constructor_706() MyHashMap {
	result := MyHashMap{
		Array: [1000][]*MyMap{},
	}
	return result
}

func (this *MyHashMap) Put(key int, value int) {

	key_ := key % base
	for _, v := range this.Array[key_] {
		if v.Key == key {
			v.Value = value
			return
		}
	}
	this.Array[key_] = append(this.Array[key_], &MyMap{
		Key:   key,
		Value: value,
	})
}

func (this *MyHashMap) Get(key int) int {
	key_ := key % base

	for _, v := range this.Array[key_] {
		if v.Key == key {
			return v.Value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	key_ := key % base
	for i, v := range this.Array[key_] {
		if v.Key == key {
			this.Array[key_] = append(this.Array[key_][:i], this.Array[key_][i+1:]...)
			return
		}
	}
}
