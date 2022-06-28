package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalBirdCount := 0
	for _, i := range birdsPerDay {
		totalBirdCount += i
	}
	return totalBirdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	totalBirdCount := 0
	for w, i := range birdsPerDay {
		if week == (w/7)+1 {
			totalBirdCount += i
		}
	}
	return totalBirdCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for w, i := range birdsPerDay {
		if w%2 == 0 {
			birdsPerDay[w] = i + 1
		}
	}
	return birdsPerDay
}
