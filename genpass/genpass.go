package genpass

import (
	"time"

	"../rnd"
	"../service"
)

var allowedChars = []rune("QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm0987654321")
var specialChars = []rune("~!@#$%^&*()_+-=|{}[]\\/';:")

// generate password
// strongPassword - include to password special characters
func Generate(passwordLength int, strongPassword bool) string {

	rnd.New(time.Now().UTC().UnixNano())
	charsArray := make([]rune, passwordLength)
	for i := range charsArray {
		if strongPassword && service.Itob(i%rnd.Next(3, passwordLength)) {
			charsArray[i] = specialChars[rnd.Next(0, len(specialChars))]
		} else {
			charsArray[i] = allowedChars[rnd.Next(0, len(allowedChars))]
		}
	}
	return string(charsArray)
}