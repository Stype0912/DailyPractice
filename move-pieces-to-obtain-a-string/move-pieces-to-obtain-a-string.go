package move_pieces_to_obtain_a_string

func canChange(start string, target string) bool {
	i, j, n := 0, 0, len(start)
	for i < n && j < n {
		for i < n && start[i] == '_' {
			i++
		}
		for j < n && target[j] == '_' {
			j++
		}
		if i < n && j < n {
			if start[i] != target[j] {
				return false
			}
			if (start[i] == 'L' && i < j) || (start[i] == 'R' && i > j) {
				return false
			}
			i++
			j++
		}
	}
	for i < n {
		if start[i] != '_' {
			return false
		}
		i++
	}
	for j < n {
		if target[i] != '_' {
			return false
		}
		j++
	}
	return true
}
