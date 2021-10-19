package word

type WordDictionary struct {
	Root *Node
	Height int
}

type Node struct {
	CharList [26]*Node
	CharListEnd [26]bool
}

func Constructor() WordDictionary {
	return WordDictionary{Root: &Node{}}
}


func (this *WordDictionary) AddWord(word string)  {
	cur := this.Root
	length := len(word)
	if length < 1 {
		return
	}
	if length > this.Height {
		this.Height = length
	}

	for i := 0; i < length; i++ {
		index := word[i] - 'a'

		if cur.CharList[index] == nil {
			cur.CharList[index] = &Node{}
		}

		if i + 1 == length {
			cur.CharListEnd[index] = true
		}

		cur = cur.CharList[index]
	}

}


func (this *WordDictionary) Search(word string) bool {
	length := len(word)
	if length > this.Height {
		return false
	}

	return search(this.Root, word)
}

func search(root *Node, word string) bool {
	length := len(word)
	cur := root
	for i := 0; i < length; i++ {
		if cur == nil {
			return false
		}
		if word[i] != '.' {
			index := word[i] - 'a'
			if cur.CharList[index] == nil {
				return false
			} else {
				if i + 1 == length {
					return cur.CharListEnd[index]
				} else {
					cur = cur.CharList[index]
				}
			}
		} else {
			for j := 0; j < 26; j++ {
				var s string
				if i + 1 == length {
					s = string(rune('a'+j))
				} else {
					s = string(rune('a'+j))+word[i+1:]
				}
				if search(cur, s) {
					return true
				}
			}

			break
		}
	}
	return false
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
