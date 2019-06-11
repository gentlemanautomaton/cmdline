package cmdlinewindows

// SplitCommand is like Split, but the first argument is returned separately.
func SplitCommand(cl string) (name string, args []string) {
	a := Split(cl)
	switch len(a) {
	case 0:
		return
	case 1:
		name = a[0]
	default:
		name = a[0]
		args = a[1:]
	}
	return
}

// SplitLiteralCommand is like SplitLiteral, but the first argument is
// returned separately.
func SplitLiteralCommand(cl string) (name string, args []string) {
	a := SplitLiteral(cl)
	switch len(a) {
	case 0:
		return
	case 1:
		name = a[0]
	default:
		name = a[0]
		args = a[1:]
	}
	return
}

// Split breaks the given command line into arguments. The split is performed
// according to the standard windows shell parsing rules as implemented by the
// Microsoft C compiler.
//
// For details on command line parsing in Windows see:
// https://docs.microsoft.com/en-us/cpp/c-language/parsing-c-command-line-arguments
func Split(cl string) (args []string) {
	var (
		buffer  = make([]rune, len(cl)) // Buffer for current argument
		inQuote bool                    // Are we in a section escaped by quotes?
		inArg   bool                    // Have we read some portion of an argument?
		slashes int                     // Current number of contiguous backslashes
		b       int                     // Current position within buffer
	)

	for _, runeValue := range cl {
		switch runeValue {
		case '\\':
			inArg = true
			slashes++
		case '"':
			inArg = true
			switch slashes % 2 {
			case 0:
				// A quote preceded by an even number of backslashes
				b, slashes = writeSlashes(buffer, b, slashes/2)
				inQuote = !inQuote
			case 1:
				// A quote preceded by an odd number of backslashes
				b, slashes = writeSlashes(buffer, b, (slashes-1)/2)
				b = write(buffer, b, '"')
			}
		case ' ', '\t':
			b, slashes = writeSlashes(buffer, b, slashes)
			if inQuote {
				// A whitespace character within a quoted section
				b = write(buffer, b, runeValue)
				break
			}
			if inArg {
				// A whitespace character terminating an argument
				b = flushArg(&args, buffer, b)
				inArg = false
			}
		default:
			inArg = true
			b, slashes = writeSlashes(buffer, b, slashes)
			b = write(buffer, b, runeValue)
		}
	}

	b, _ = writeSlashes(buffer, b, slashes)
	if inArg {
		flushArg(&args, buffer, b)
	}

	return
}

// SplitLiteral is like Split, but it preserves all quotes and escape
// characters within the returned arguments.
func SplitLiteral(cl string) (args []string) {
	var (
		buffer  = make([]rune, len(cl)) // Buffer for current argument
		inQuote bool                    // Are we in a section escaped by quotes?
		inArg   bool                    // Have we read some portion of an argument?
		slashes int                     // Current number of contiguous backslashes
		b       int                     // Current position within buffer
	)

	for _, runeValue := range cl {
		switch runeValue {
		case '\\':
			b = write(buffer, b, runeValue)
			inArg = true
			slashes++
		case '"':
			b = write(buffer, b, runeValue)
			inArg = true
			switch slashes % 2 {
			case 0:
				// A quote preceded by an even number of backslashes
				inQuote = !inQuote
			case 1:
				// A quote preceded by an odd number of backslashes
			}
			slashes = 0
		case ' ', '\t':
			slashes = 0
			if inQuote {
				// A whitespace character within a quoted section
				b = write(buffer, b, runeValue)
				break
			}
			if inArg {
				// A whitespace character terminating an argument
				b = flushArg(&args, buffer, b)
				inArg = false
			}
		default:
			inArg = true
			b = write(buffer, b, runeValue)
			slashes = 0
		}
	}

	if inArg {
		flushArg(&args, buffer, b)
	}

	return
}

func writeSlashes(buffer []rune, b, slashes int) (int, int) {
	for i := 0; i < slashes; i++ {
		buffer[b+i] = '\\'
	}
	return b + slashes, 0
}

func write(buffer []rune, b int, runeValue rune) int {
	buffer[b] = runeValue
	return b + 1
}

func flushArg(args *[]string, buffer []rune, b int) int {
	*args = append(*args, string(buffer[0:b]))
	return 0
}
