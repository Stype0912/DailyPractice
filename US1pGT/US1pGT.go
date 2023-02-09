package US1pGT

type MagicDictionary struct {
	Freq map[int][]string
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	this.Freq = make(map[int][]string)
	for _, word := range dictionary {
		//lengthMatchedWords, ok := this.Freq[len(word)]
		//if !ok {
		//	this.Freq[len(word)] = []string{word}
		//} else {
		//	lengthMatchedWords = append(lengthMatchedWords, word)
		//	this.Freq[len(word)] = lengthMatchedWords
		//}
		this.Freq[len(word)] = append(this.Freq[len(word)], word)
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	lengthMatchedWords, ok := this.Freq[len(searchWord)]
	if !ok {
		return false
	}
	for _, lengthMatchedWord := range lengthMatchedWords {
		cnt := 0
		for i, _ := range searchWord {
			if searchWord[i] != lengthMatchedWord[i] {
				cnt++
			}
		}
		if cnt == 1 {
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
