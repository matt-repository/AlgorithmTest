package simple

import "fmt"

//Leed Code  383. 赎金信
//给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
//如果可以，返回 true ；否则返回 false 。
//magazine 中的每个字符只能在 ransomNote 中使用一次
//示例 1：
//
//输入：ransomNote = "a", magazine = "b"
//输出：false

//输入：ransomNote = "aa", magazine = "ab"
//输出：false

//输入：ransomNote = "aa", magazine = "aab"
//输出：true

func CanConstruct(ransomNote string, magazine string) bool {
	mapCache1 := [26]int{}
	fmt.Println(len(mapCache1))
	mapCache := map[byte]int{}
	if len(ransomNote) > len(magazine) {
		return false
	}
	for i := 0; i < len(magazine); i++ {
		_, ok := mapCache[magazine[i]]
		if !ok {
			mapCache[magazine[i]] = 1
		} else {
			mapCache[magazine[i]]++
		}
	}

	for i := 0; i < len(ransomNote); i++ {
		_, ok := mapCache[ransomNote[i]]
		if !ok {
			return false
		} else {
			mapCache[ransomNote[i]]--
			if mapCache[ransomNote[i]] == 0 {
				delete(mapCache, ransomNote[i])
			}
		}
	}
	return true
}
