package clock

import "fmt"

const testVersion = 4

const (
	day = 24 * 60
)

// Clock type
type Clock struct {
	Hour, Minute int
}

// New initializes new clock
func New(hour, minute int) Clock {
	totalMinutes := checkTotalMinutes(hour*60 + minute)
	return Clock{totalMinutes / 60, totalMinutes % 60}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.Hour, clock.Minute)
}

// Add adds minutes to clock
func (clock Clock) Add(minutes int) Clock {
	oldMinutes := clock.Hour*60 + clock.Minute
	newMinutes := checkTotalMinutes(oldMinutes + minutes)
	return Clock{newMinutes / 60, newMinutes % 60}
}

func checkTotalMinutes(minutes int) int {
	minutes = minutes % day
	if minutes < 0 {
		minutes = day + minutes
	}
	return minutes
}
