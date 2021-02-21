// Package secret converts a decimal number to an appropriate sequence of events
package secret

// Handshake returns the appropriate handshake sequence for a given integer
func Handshake(n uint) []string {

	var result []string

	operators := []uint{1, 2, 4, 8, 16}

	actions := map[uint]string{
		16: "reverse",
		8:  "jump",
		4:  "close your eyes",
		2:  "double blink",
		1:  "wink",
	}

	for _, thisOp := range operators {
		action, ok := actions[thisOp&n]
		if ok {
			if action == "reverse" {
				result = reverse(result)
			} else {
				result = append(result, action)
			}
		}
	}

	return result

}

func reverse(ss []string) []string {
	var result []string

	for i := len(ss) - 1; i >= 0; i-- {
		result = append(result, ss[i])
	}

	return result
}
