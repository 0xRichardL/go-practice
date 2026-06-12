package main

func canConstruct(ransomNote string, magazine string) bool {
	check := [26]int16{}
	for _, v := range magazine {
		check[v-'a']++
	}
	for _, v := range ransomNote {
		check[v-'a']--
		if check[v-'a'] < 0 {
			return false
		}
	}
	return true
}
