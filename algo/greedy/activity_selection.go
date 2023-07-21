package greedy

import "sort"

// ActivitySelection takes a list of activities an 2-element array of start time and finish time
// selects the maximum number of activities that can be performed by a single person, assuming that
// a person can only work on a single activity at a time. If activities are not in the correct
// format (0 <= start time < finish time), nil is returned.
func ActivitySelection(activities [][2]int, verifyActivities bool) [][2]int {
	if len(activities) == 0 {
		return [][2]int{}
	}

	// Verify activities.
	if verifyActivities {
		for _, activity := range activities {
			if activity[0] < 0 || activity[0] >= activity[1] {
				return nil
			}
		}
	}

	// Sort the activities based on the finish times.
	sort.Slice(activities, func(i, j int) bool {
		if activities[i][1] <= activities[j][1] {
			return true
		}

		return false
	})

	activity := activities[0]
	selection := [][2]int{activities[0]}
	for i := 1; i < len(activities); i++ {
		if activities[i][0] >= activity[1] {
			activity = activities[i]
			selection = append(selection, activity)
		}
	}

	return selection
}
