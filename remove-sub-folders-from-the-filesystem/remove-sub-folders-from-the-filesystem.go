package remove_sub_folders_from_the_filesystem

import "strings"

type Trie struct {
	children map[string]*Trie
	fid      int
}

func newTrie() *Trie {
	return &Trie{children: map[string]*Trie{}, fid: -1}
}

func (trie *Trie) insert(folder string, fid int) {
	node := trie
	folderList := strings.Split(folder, "/")
	for _, folderName := range folderList[1:] {
		if _, ok := node.children[folderName]; !ok {
			node.children[folderName] = newTrie()
		}
		node = node.children[folderName]
	}
	node.fid = fid
}

func (trie *Trie) search() (ans []int) {
	var dfs func(*Trie)
	dfs = func(trie *Trie) {
		if trie.fid != -1 {
			ans = append(ans, trie.fid)
			return
		}
		for _, child := range trie.children {
			dfs(child)
		}
	}
	dfs(trie)
	return
}

func removeSubfolders(folder []string) []string {
	root := newTrie()
	for i, folderName := range folder {
		root.insert(folderName, i)
	}
	var res []string
	for _, fid := range root.search() {
		res = append(res, folder[fid])
	}
	return res
}
