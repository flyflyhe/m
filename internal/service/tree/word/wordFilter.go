package word

import "log"

type DictTree struct {
	Head  *[26]*DictNode
	Head2 *[26]*DictNode
}

type DictNode struct {
	V     byte
	End   bool
	Index []int //用下标标识 存在的字符串
	Child *[26]*DictNode
}

type WordFilter struct {
	Tree *DictTree
}

func ConstructorWord(words []string) WordFilter {
	wf := WordFilter{Tree: &DictTree{Head: &[26]*DictNode{}, Head2: &[26]*DictNode{}}}
	for index, word := range words {
		node := wf.Tree.Head
		build(node, word, index)

		node2 := wf.Tree.Head2
		build(node2, string(arrayReverse([]byte(word))), index)
	}
	return wf
}

func build(node *[26]*DictNode, word string, index int) {
	var v byte
	for i := 0; i < len(word); i++ { // 正向树
		v = word[i]
		if node[v-'a'] == nil {
			node[v-'a'] = &DictNode{V: v, Child: &[26]*DictNode{}, Index: make([]int, 0)}
		}
		node[v-'a'].Index = append(node[v-'a'].Index, index)
		if i == len(word)-1 {
			node[v-'a'].End = true
		}
		node = node[v-'a'].Child
	}
}

func (this *WordFilter) Dfs(node *[26]*DictNode, bs []byte, result *[]string) {
	//log.Println(node)
	for _, v := range node {
		if v == nil {
			continue
		}

		log.Println(v.V, v.Index)
		bs = append(bs, v.V)
		if v.End {
			*result = append(*result, string(bs))
		}
		this.Dfs(v.Child, bs, result)
		bs = bs[:len(bs)-1]
	}
}

func arrayReverse(arr []byte) []byte {
	rArr := make([]byte, len(arr))
	for i := 0; i < len(arr); i++ {
		rArr[len(arr)-i-1] = arr[i]
	}
	return rArr
}

func query(node *[26]*DictNode, word string) []int {
	var list []int
	var v byte
	var next *DictNode
	for i := 0; i < len(word); i++ {
		v = word[i]
		next = node[v-'a']
		if next == nil {
			return list
		}
		if i == len(word)-1 {
			list = next.Index
			break
		}
		node = next.Child
	}

	return list
}

func (this *WordFilter) F(pref string, suff string) int {

	node := this.Tree.Head
	list1 := query(node, pref)
	node = this.Tree.Head2
	list2 := query(node, string(arrayReverse([]byte(suff))))

	log.Println(list1)
	log.Println(list2)
	ans := -1
	i := len(list1) - 1
	j := len(list2) - 1
	for i >= 0 && j >= 0 {
		if list1[i] > list2[j] {
			i--
		} else if list1[i] < list2[j] {
			j--
		} else {
			return list1[i]
		}
	}

	return ans
}
