package app

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// Strings provides compact, one-line examples of common operations you might
// want as a quick reference. Values are assigned to the blank identifier
// when not used so the file stays runnable and builds cleanly.
func Strings() {
	s := "  Welcome to the jungle  "

	// --- basic queries ---
	_ = len(s)                        // byte length
	_ = utf8.RuneCountInString(s)     // rune (character) count
	_ = strings.Index(s, "the")       // first index of substring or -1
	_ = strings.IndexByte(s, 'j')     // index of byte
	_ = strings.LastIndex(s, "e")     // last index
	_ = strings.HasPrefix(s, "  We")  // prefix check
	_ = strings.HasSuffix(s, "le  ")  // suffix check
	_ = strings.Contains(s, "jungle") // containment

	// --- case and trimming ---
	_ = strings.ToLower(s)
	_ = strings.ToUpper(s)
	_ = strings.TrimSpace(s) // remove leading/trailing Unicode spaces
	_ = strings.TrimLeft(s, " \t\n")
	_ = strings.TrimRight(s, " \t\n")

	// --- splitting and joining ---
	_ = strings.Fields(s)         // split on any whitespace
	_ = strings.Split(s, " ")     // split on literal space
	_ = strings.SplitN(s, " ", 3) // limit splits
	_ = strings.Join([]string{"a", "b"}, "-")

	// --- replacement and regex ---
	_ = strings.ReplaceAll(s, " ", "_") // replace all non-regex
	// regex replace (use regexp for patterns)
	_ = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")

	// --- substring and slicing ---
	_ = s[2:7] // substring by byte indices (careful with runes)
	_ = s[:7]
	_ = s[7:]

	// --- conversion helpers ---
	_ = strings.TrimPrefix(s, "  ")
	_ = strings.TrimSuffix(s, "  ")

	// --- builders (mutable accumulation) ---
	var sb strings.Builder
	sb.WriteString("Work on your system")
	sb.WriteByte('!')
	sb.WriteRune('✓')
	_ = sb.String() // get final string

	// --- utility / counts ---
	_ = strings.Count(s, "e")  // count non-overlapping occurrences
	_ = strings.Repeat("x", 3) // "xxx"
	_ = strings.IndexFunc(s, func(r rune) bool { return r == 'j' })

	// --- notes ---
	// Use strings.Builder when performing many concatenations to avoid
	// repeated allocations. Use regexp for pattern-based matching and
	// replacements. For UTF-8-aware slicing and length, use the utf8
	// package or convert to []rune when you need rune indexing.
}

// main intentionally does not print. Run `Strings()` or open this file to
// read the examples; the file is kept runnable so `go build` succeeds.
func main() {}
