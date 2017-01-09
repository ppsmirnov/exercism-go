package gigasecond

import "time"

const testVersion = 4
const gigasecond = time.Second * 1000000000

// AddGigasecond adds a gigasecond to birthDate
func AddGigasecond(birthDate time.Time) time.Time {
	return birthDate.Add(gigasecond)
}
