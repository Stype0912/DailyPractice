package reward_top_k_students

import (
	"sort"
	"strings"
)

type Stu struct {
	Id    int
	Grade int
}

func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	Students := make([]Stu, len(student_id))
	for i := 0; i < len(student_id); i++ {
		Students[i].Id = student_id[i]
	}
	positive_feedbackMap := map[string]int{}
	for _, single_positive_feedback := range positive_feedback {
		positive_feedbackMap[single_positive_feedback]++
	}
	negative_feedbackMap := map[string]int{}
	for _, single_negative_feedback := range negative_feedback {
		negative_feedbackMap[single_negative_feedback]++
	}
	for i, singleReport := range report {
		words := strings.Split(singleReport, " ")
		for _, word := range words {
			if _, ok := positive_feedbackMap[word]; ok {
				Students[i].Grade += 3
			}
			if _, ok := negative_feedbackMap[word]; ok {
				Students[i].Grade--
			}
		}
	}
	sort.Slice(Students, func(i, j int) bool {
		if Students[i].Grade == Students[j].Grade {
			return Students[i].Id < Students[j].Id
		}
		return Students[i].Grade > Students[j].Grade
	})
	ans := []int{}
	for i := 0; i < k; i++ {
		ans = append(ans, Students[i].Id)
	}
	return ans
}
