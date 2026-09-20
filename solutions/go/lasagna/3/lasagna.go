package lasagna


const OvenTime = 40
const PreparationTimeMinutes = 2

func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime	- actualMinutesInOven
}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * PreparationTimeMinutes
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return actualMinutesInOven + PreparationTime(numberOfLayers)
}
