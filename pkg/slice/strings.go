package slice

// StringsVertical ...
func StringsVertical(in []string, index int) (out string) {
	out = ""
	for i := 0; i < len(in); i++ {
		if in[i][index:index+1] == "1" {
			out += "1"
		} else {
			out += "0"
		}
	}
	return
}
