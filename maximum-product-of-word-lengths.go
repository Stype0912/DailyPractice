package main

func maxProduct(words []string) int {
	var bin []int64
	ans := 0
	for _, item := range words {
		bin = append(bin, toBin(item))
	}
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if bin[i]&bin[j] == 0 && len(words[i])*len(words[j]) > ans {
				ans = len(words[i]) * len(words[j])
			}
		}
	}
	return ans
}

func toBin(s string) int64 {
	var ans int64 = 0
	for _, item := range s {
		ans |= 1<<(int64(item) - 97)
	}
	return ans
}