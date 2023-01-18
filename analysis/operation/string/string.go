package main

import (
	"fmt"
)

/**
Step 1: Define a function that accepts a string, i.e., str.
Step 2: Convert str to byte string.
Step 3: Start iterating the byte string.
Step 4: Swap the first element with the last element of converted byte string.
Step 5: Convert the byte string to string and return it.
*/

func reverseStringManual(strn string) string {
	_bytestr := []byte(strn)
	strLength := len(_bytestr)
	for i, j := 0, (strLength - 1); i < j; i, j = i+1, j-1 {
		fmt.Printf("\nCharacter = %c Bytes = %v", _bytestr[i], _bytestr[i])
		_bytestr[i], _bytestr[j] = _bytestr[j], _bytestr[i]
	}
	return string(_bytestr)
}

// Main function
func main() {

	// Creating and initializing a string
	// str := "Welcome to GeeksforGeeks"
	// _bytestr := []byte(str)
	// // Accessing the bytes of the given string
	// for _, c := range _bytestr {
	// 	// fmt.Println(string(_bytestr[i]), c)
	// 	fmt.Printf("\nCharacter = %c Bytes = %v", c, c)
	// }
	// fmt.Println(reverseStringManual("Welcome to GeeksforGeeks"))
	// fmt.Println(lengthOfLongestSubstring("a"))
	fmt.Println(reverseSubStr("abcdefg", 3))
	// fmt.Println(canConstruct("bhjdigif", "dbjdhdehgbcdjjgadeegdbegehjffie"))
	//fmt.Println(lengthOfLongestSubstring("bbbbb"))
}

func reverseSubStr(s string, k int) string {
	byte_str := []byte(s)

	for i := 0; i < len(byte_str); {

		revLengthLimit := (i + k - 1)
		fmt.Println(i, len(byte_str), revLengthLimit)
		if revLengthLimit >= len(byte_str) {
			revLengthLimit = len(byte_str) - 1
		}

		for j, l := i, revLengthLimit; j < l; j, l = j+1, l-1 {
			byte_str[j], byte_str[l] = byte_str[l], byte_str[j]
		}
		i += (2 * k)

	}

	return string(byte_str)
}

// Q: Given a string s, find the length of the longest substring without repeating characters.
func lengthOfLongestSubstring(str string) int {
	maxChar := make(map[uint8]int)
	max, left := 0, 0
	_bytestr := []byte(str)
	for idx, c := range _bytestr {
		if _, okay := maxChar[c]; okay == true && maxChar[c] >= left {
			if idx-left > max {
				max = idx - left
			}
			left = maxChar[c] + 1
		}
		maxChar[c] = idx
	}
	if len(str)-left > max {
		max = len(str) - left
	}
	return max
}

func canConstruct(ransomNote string, magazine string) bool {
	// _byteRansomNote := []rune(ransomNote)
	// // _byteMagazine := []byte(ransomNote)
	// if len(ransomNote) == 1 && len(magazine) == 1 {
	// 	if ransomNote == magazine {
	// 		return true
	// 	}
	// 	return false
	// }

	// if len(ransomNote) == 1 && strings.Contains(magazine, ransomNote) {
	// 	return true
	// }

	// rStr := magazine
	// fmt.Println(ransomNote)
	// for i := 0; i <= len(_byteRansomNote)-1; i = i + 1 {
	// 	fmt.Println(i, rStr, (i+1 < len(_byteRansomNote)))
	// 	if strings.Contains(rStr, string(_byteRansomNote[i])) && (i+1 < len(_byteRansomNote)) {
	// 		if string(_byteRansomNote[i]) == string(_byteRansomNote[i+1]) {
	// 			sIndex := strings.Index(rStr, string(_byteRansomNote[i]))
	// 			rStr = rStr[sIndex+1:]
	// 			fmt.Println(rStr)
	// 		} else {
	// 			rStr = magazine
	// 		}
	// 	} else {
	// 		if (i+1 == len(_byteRansomNote)) && strings.Contains(rStr, string(_byteRansomNote[i])) {
	// 			return true
	// 		}
	// 		return false
	// 	}
	// }
	// return true

	charSlice := make([]int, 26, 26)

	for i := 0; i < len(magazine); i++ {
		fmt.Println(int(magazine[i])-int('a'), int(magazine[i]), int('a'))
		charSlice[int(magazine[i])-int('a')] += 1
	}

	for j := 0; j < len(ransomNote); j++ {
		charSlice[int(ransomNote[j])-int('a')] -= 1
		if charSlice[int(ransomNote[j])-int('a')] < 0 {
			return false
		}
	}

	return true
}
