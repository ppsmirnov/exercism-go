package robotname

import "fmt"

var counter = 0

// Robot type
type Robot struct {
	name string
}

// New returns new robot
func New() *Robot {
	r := new(Robot)
	r.name = fmt.Sprintf("AA%03d", counter)
	counter++
	return r
}

// Name returns robot's name
func (r *Robot) Name() string {
	return r.name
}

// Reset resets robot's name
func (r *Robot) Reset() {
	r.name = fmt.Sprintf("AA%03d", counter)
	counter++
}
