// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

// Given a string s, return true if it is a palindrome, or false otherwise.

// Example 1:

// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.
// Example 2:

// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.
// Example 3:

// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.

func isPalindrome(s string) bool {
	var cleanStr []rune
	for _, i := range s {
		if unicode.IsLetter(i) || unicode.IsDigit(i) {
			cleanStr = append(cleanStr, unicode.ToLower(i))
		}
	}
	left, right := 0, len(cleanStr)-1
	for left < right {
		if cleanStr[left] != cleanStr[right] {
			return false
		}
		left++
		right--
	}
	return true
}