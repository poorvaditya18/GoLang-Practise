package welcome

import (
	// "fmt"
	"strings"
)

// unlike C,C++ , in go parameter syntax <param_name> <data_type>
func SayWelcome(names []string) string{
	if len(names) == 0 {
		names = []string{"Poorva"}
	}
	return "Welcome, " + strings.Join(names,", ") + "!"
	// return fmt.Sprintf("Welcome, %s!", names)
}