package hackerrank

func SuperReducedStringSolution(input string) string {
	l := len(input)
	if l <= 1 {
		return input
	}
	var output string
	changed := true
	for changed == true {
		output = ""
		changed = false

		for i := 0; i < len(input); i++ {
			if i == len(input)-1 {
				output += string(input[i])
				break
			}
			if input[i] != input[i+1] {
				output += string(input[i])
			} else {
				changed = true
				i++
			}
		}
		input = output
	}
	if len(output) == 0 {
		return "Empty String"
	}
	return output
}
