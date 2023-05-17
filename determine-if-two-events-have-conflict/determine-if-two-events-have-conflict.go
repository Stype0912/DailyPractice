package determine_if_two_events_have_conflict

func haveConflict(event1 []string, event2 []string) bool {
	return !(event2[0] > event1[1] || event1[0] > event2[1])
}
