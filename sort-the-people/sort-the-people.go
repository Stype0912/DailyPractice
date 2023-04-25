package sort_the_people

import "sort"

type People struct {
	Names   []string
	Heights []int
}

func (p People) Len() int {
	return len(p.Names)
}

func (p People) Less(i, j int) bool {
	return p.Heights[i] > p.Heights[j]
}

func (p People) Swap(i, j int) {
	p.Names[i], p.Names[j] = p.Names[j], p.Names[i]
	p.Heights[i], p.Heights[j] = p.Heights[j], p.Heights[i]
}
func sortPeople(names []string, heights []int) []string {
	people := &People{
		Names:   names,
		Heights: heights,
	}
	sort.Sort(people)
	return people.Names
}
