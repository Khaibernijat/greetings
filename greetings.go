package greetings

import "fmt"

// Hello returns a greeting for the name person.
func Hello(Khaiber string) string{

//return a greeting that embeds the name in a message.
message :=fmt.Sprintf("Hi, %v. Welcome!", Khaiber)
return message

}
