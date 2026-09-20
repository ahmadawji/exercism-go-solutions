package lasagna


const OvenTime = 40
const PreparationTimeMinutes = 2

func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime	- actualMinutesInOven
}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * PreparationTimeMinutes
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return actualMinutesInOven + PreparationTime(numberOfLayers)
}
