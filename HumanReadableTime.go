package kata

import "fmt"

func HumanReadableTime(seconds int) string {
	// your code here
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	seconds = seconds % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
