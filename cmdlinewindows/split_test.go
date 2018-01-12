package cmdlinewindows_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gentlemanautomaton/cmdline/cmdlinewindows"
)

var splitTests = []struct {
	Name     string
	CL       string
	Expected []string
}{
	{"quoted-space", `"a b c" d e`, []string{`a b c`, `d`, `e`}},
	{"escaped-quote", `"ab\"c" "\\" d`, []string{`ab"c`, `\`, `d`}},
	{"unescaped-backslash", `a\\\b d"e f"g h`, []string{`a\\\b`, `de fg`, `h`}},
	{"escaped-backslash-odd", `a\\\"b c d`, []string{`a\"b`, `c`, `d`}},
	{"escaped-backslash-even", `a\\\\"b c" d e`, []string{`a\\b c`, `d`, `e`}},
	{"empty-arg-1", `a "" c`, []string{`a`, ``, `c`}},
	{"empty-arg-2", `a "" c ""`, []string{`a`, ``, `c`, ``}},
	{"empty-arg-3", `a b "" ""`, []string{`a`, `b`, ``, ``}},
	{"uneven-quote-1", `a"bc`, []string{`abc`}},
	{"uneven-quote-2", `a"b c d`, []string{`ab c d`}},
	{"uneven-quote-3", `ab c d"`, []string{`ab`, `c`, `d`}},
	{"uneven-quote-4", `a "b c d`, []string{`a`, `b c d`}},
}

func TestSplit(t *testing.T) {
	for _, tc := range splitTests {
		tc := tc // capture range variable
		t.Run(tc.Name, func(t *testing.T) {
			actual := cmdlinewindows.Split(tc.CL)
			if !testEq(actual, tc.Expected) {
				e := strings.TrimPrefix(fmt.Sprintf("%#v", tc.Expected), "[]string")
				a := strings.TrimPrefix(fmt.Sprintf("%#v", actual), "[]string")
				t.Errorf("input: %s   expected: %s   actual: %s", tc.CL, e, a)
			}
		})
	}
}

func testEq(a []string, b []string) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
