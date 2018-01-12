package cmdlineposix

// Split breaks the given command line into arguments. The arguments are split
// according to posix shell parsing rules.
//
// For an introduction to command line parsing in unix shells, see:
// http://www.grymoire.com/Unix/Quote.html
//
// For details on posix shell parsing requirements see:
// http://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html#tag_18
func Split(cl string) (args []string) {
	var (
		buffer        = make([]rune, len(cl)) // Buffer for current argument
		inEscape      bool                    // Are we in a section escaped by weak quotes?
		inSingleQuote bool                    // Are we in a section escaped by strong quotes?
		inDoubleQuote bool                    // Are we in a section escaped by weak quotes?
		inArg         bool                    // Have we read some portion of an argument?
		b             int                     // Current position within buffer
	)

	for _, runeValue := range cl {
		switch runeValue {
		case '\\':
			inArg = true
			switch {
			case inEscape:
				inEscape = false
				fallthrough
			case inSingleQuote:
				b = write(buffer, b, runeValue)
			default:
				inEscape = true
			}
		case '\'':
			inArg = true
			switch {
			case inDoubleQuote && inEscape:
				b = write(buffer, b, '\\')
				fallthrough
			case inEscape:
				inEscape = false
				fallthrough
			case inDoubleQuote:
				b = write(buffer, b, runeValue)
			default:
				inSingleQuote = !inSingleQuote
			}
		case '"':
			inArg = true
			switch {
			case inEscape:
				inEscape = false
				fallthrough
			case inSingleQuote:
				b = write(buffer, b, runeValue)
			default:
				inDoubleQuote = !inDoubleQuote
			}
		case ' ', '\t':
			switch {
			case inEscape:
				// A whitespace character after an escaping backslash
				inEscape = false
				fallthrough
			case inSingleQuote, inDoubleQuote:
				// A whitespace character within a quoted section
				b = write(buffer, b, runeValue)
			default:
				if inArg {
					// A whitespace character terminating an argument
					b = flushArg(&args, buffer, b)
					inArg = false
				}
			}
		case '\n':
			switch {
			case inEscape:
				// A line return after a trailing backslash, indicating continuation
				inEscape = false
			case inArg:
				// A whitespace character terminating an argument
				b = flushArg(&args, buffer, b)
				inArg = false
			}
		default:
			inArg = true
			switch {
			case inEscape:
				inEscape = false
				switch runeValue {
				case '$', '`':
				default:
					b = write(buffer, b, '\\')
				}
			}
			b = write(buffer, b, runeValue)
		}
	}

	if inEscape {
		inArg = true
		b = write(buffer, b, '\\')
	}

	if inArg {
		flushArg(&args, buffer, b)
	}

	return
}

func write(buffer []rune, b int, runeValue rune) int {
	buffer[b] = runeValue
	return b + 1
}

func flushArg(args *[]string, buffer []rune, b int) int {
	*args = append(*args, string(buffer[0:b]))
	return 0
}
