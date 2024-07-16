package functions

import "fmt"

func FormatXXD(data []byte) string {
	var result string
	for i := 0; i < len(data); i += 16 {
		result += fmt.Sprintf("%08x: ", i)
		for j := 0; j < 16; j++ {
			if i+j < len(data) {
				result += FormatHex(data[i+j])
			} else {
				result += "   "
			}
		}
		result += " "
		for j := 0; j < 16; j++ {
			if i+j < len(data) {
				result += FormatAsciiColored(data[i+j], false)
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

func FormatHex(b byte) string {
	switch b {
	case 0xff, 0x00:
		return fmt.Sprintf("\033[1;37m%02x\033[0m ", b) // White
	default:
		return fmt.Sprintf("\033[1;90m%02x\033[0m ", b) // Gray
	}
}

func FormatAscii(b byte) string {
	if 32 <= b && b < 127 {
		return fmt.Sprintf("%c", b)
	}
	return "."
}

func FormatAsciiColored(b byte, isDifferent bool) string {
	color := "\033[1;90m" // Gray
	if b == 0xff || b == 0x00 {
		color = "\033[1;37m" // White
	} else if isDifferent {
		color = "\033[1;31m" // Red
	}
	return fmt.Sprintf("%s%s\033[0m", color, FormatAscii(b))
}
