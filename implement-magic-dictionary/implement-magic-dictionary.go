package implement_magic_dictionary

type MagicDictionary []string

func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	*this = dictionary
}

func (this *MagicDictionary) Search(searchWord string) bool {
	for _, word := range *this {
		if len(word) != len(searchWord) {
			continue
		}
		diff := 0
		for i := 0; i < len(word); i++ {
			if word[i] != searchWord[i] {
				diff++
			}
			if diff > 1 {
				break
			}
		}
		if diff == 1 {
			return true
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
