package twelve

const testVersion = 1

var (
	dayNumbers = [...]string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
	ammount    = [...]string{"", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve"}
	presents   = [...]string{"Drummers Drumming", "Pipers Piping", "Lords-a-Leaping", "Ladies Dancing", "Maids-a-Milking", "Swans-a-Swimming", "Geese-a-Laying", "Gold Rings", "Calling Birds", "French Hens", "Turtle Doves", "Partridge in a Pear Tree"}
)

// Verse prints one line of the song
func Verse(day int) string {
	result := "On the " + dayNumbers[day-1] + " day of Christmas my true love gave to me, "

	sep := " "
	end := ", "
	var lastAnd string

	for i := day; i > 0; i-- {
		if i <= 1 {
			sep = ""
			end = "."

			if day > 1 {
				lastAnd = "and a "
			} else {
				lastAnd = "a "
			}
		}
		result += ammount[i-1] + sep + lastAnd + presents[12-i] + end
	}

	return result
}

// Song prints whole song
func Song() string {
	var result string
	for i := 1; i < 13; i++ {
		result += Verse(i) + "\n"
	}
	return result
}
