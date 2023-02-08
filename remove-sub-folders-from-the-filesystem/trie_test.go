package remove_sub_folders_from_the_filesystem

import "testing"

func TestTrie(t *testing.T) {
	t.Log(removeSubfolders([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}))
}
