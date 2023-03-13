package minimum_hours_of_training_to_win_a_competition

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	ans := 0
	energySum := 0
	for i := 0; i < len(energy); i++ {
		energySum += energy[i]
	}
	if initialEnergy <= energySum {
		ans += energySum - initialEnergy + 1
	}
	currentExperience := initialExperience
	for i := 0; i < len(experience); i++ {
		if currentExperience <= experience[i] {
			ans += experience[i] - currentExperience + 1
			currentExperience += experience[i] - currentExperience + 1
		}
		currentExperience += experience[i]
	}
	return ans
}
