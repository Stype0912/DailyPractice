package count_days_spent_together

import "strconv"

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	arriveAliceDay := dateToDay(arriveAlice)
	leaveAliceDay := dateToDay(leaveAlice)
	arriveBobDay := dateToDay(arriveBob)
	leaveBobDay := dateToDay(leaveBob)
	if arriveBobDay > leaveAliceDay || arriveAliceDay > leaveBobDay {
		return 0
	}
	if arriveAliceDay > arriveBobDay {
		if leaveAliceDay < leaveBobDay {
			return int(leaveAliceDay - arriveAliceDay + 1)
		} else {
			return int(leaveBobDay - arriveAliceDay + 1)
		}
	} else {
		if leaveAliceDay < leaveBobDay {
			return int(leaveAliceDay - arriveBobDay + 1)
		} else {
			return int(leaveBobDay - arriveBobDay + 1)
		}
	}
}

func dateToDay(date string) int64 {
	month, _ := strconv.ParseInt(date[:2], 10, 64)
	day, _ := strconv.ParseInt(date[3:], 10, 64)
	ans := int64(0)
	dayPerMonth := []int64{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for i := int64(1); i < month; i++ {
		ans += dayPerMonth[i-1]
	}
	return ans + day
}
