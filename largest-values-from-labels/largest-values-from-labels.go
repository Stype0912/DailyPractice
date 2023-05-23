package largest_values_from_labels

import "sort"

type ValueLabel struct {
	values []int
	labels []int
}

func (vl ValueLabel) Len() int {
	return len(vl.values)
}

func (vl ValueLabel) Less(i, j int) bool {
	return vl.values[i] > vl.values[j]
}

func (vl ValueLabel) Swap(i, j int) {
	vl.values[i], vl.values[j] = vl.values[j], vl.values[i]
	vl.labels[i], vl.labels[j] = vl.labels[j], vl.labels[i]
}

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	valueLabel := ValueLabel{
		values: values,
		labels: labels,
	}
	sort.Sort(valueLabel)
	values = valueLabel.values
	labels = valueLabel.labels
	labelCountMap := map[int]int{}
	ans := 0
	count := 0
	for i := 0; i < len(values); i++ {
		if count >= numWanted {
			return ans
		}
		if labelCountMap[labels[i]] < useLimit {
			ans += values[i]
			labelCountMap[labels[i]]++
			count++
		}
	}
	return ans
}
