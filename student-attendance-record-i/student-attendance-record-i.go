package student_attendance_record_if

import "strings"

func checkRecord(s string) bool {
	if strings.Count(s, "A") < 2 && !strings.Contains(s, "LLL") {
		return true
	} else {
		return false
	}
}
