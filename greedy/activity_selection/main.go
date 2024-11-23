package main

import (
	"fmt"
)

// Activity represents an activity with a start and end time
type Activity struct {
	start, finish int
}

// activitySelection selects the maximum number of non-overlapping activities
func activitySelection(activities []Activity) []Activity {
	var selectedActivities []Activity

	// TODO: implement

	return selectedActivities
}

func main() {
	// Input: List of activities with start and finish times
	activities := []Activity{
		{1, 3},
		{2, 5},
		{4, 6},
		{6, 7},
		{5, 8},
		{8, 9},
	}

	// Get the selected activities
	selectedActivities := activitySelection(activities)

	// Print the result
	fmt.Println("Selected activities:")
	for _, activity := range selectedActivities {
		fmt.Printf("Start: %d, Finish: %d\n", activity.start, activity.finish)
	}
}
