package hello

import "fmt"

const testVersion = 2

func HelloWorld(str string) string {
	if len(str) > 0 {
		return fmt.Sprintf("Hello, %s!", str)
	}
	return "Hello, World!"
}
