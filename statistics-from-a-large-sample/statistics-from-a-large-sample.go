package statistics_from_a_large_sample

import "math"

func sampleStats(count []int) []float64 {
	sum := 0
	numCount := 0
	min := math.MaxInt
	max := math.MinInt
	mean := float64(0)
	median := float64(0)
	mode := 0
	modeNum := 0
	for i := 0; i < len(count); i++ {
		if count[i] != 0 {
			if i < min {
				min = i
			}
			if i > max {
				max = i
			}
			if count[i] > modeNum {
				mode = i
				modeNum = count[i]
			}
		}
		sum += i * count[i]
		numCount += count[i]
	}
	mean = float64(sum) / float64(numCount)
	counter := 0
	for i := 0; i < len(count); i++ {
		if count[i] == 0 {
			continue
		}
		counter += count[i]
		if counter > numCount/2 {
			median = float64(i)
			break
		} else if counter == numCount/2 {
			if numCount%2 == 0 {
				for j := i + 1; j < len(count); j++ {
					if count[j] != 0 {
						median = float64(i+j) / float64(2)
						break
					}
				}
				break
			}
		}
	}
	return []float64{float64(min), float64(max), mean, median, float64(mode)}
}
