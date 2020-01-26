// A solution to the TwoFer exercism.
package twofer

import "fmt"

// Given a name, return a sentence with that name.
// If no name is specified, return a generic one.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
