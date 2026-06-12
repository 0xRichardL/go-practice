package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func compile_vs_must_compile() {
	fmt.Println("=== Compile vs MustCompile ===")

	// Compile returns error if pattern is invalid
	re1, err := regexp.Compile(`\d+`)
	if err != nil {
		fmt.Println("Compile error:", err)
	} else {
		fmt.Printf("Compile success: %v\n", re1.MatchString("123"))
	}

	// MustCompile panics if pattern is invalid (use for known-good patterns)
	re2 := regexp.MustCompile(`[a-z]+`)
	fmt.Printf("MustCompile success: %v\n", re2.MatchString("abc"))

	// Invalid pattern example (commented to avoid panic)
	// re3 := regexp.MustCompile(`[invalid(`) // would panic
	_, err2 := regexp.Compile(`[invalid(`)
	fmt.Printf("Invalid pattern error: %v\n", err2)
	fmt.Println()
}

// match demonstrates regex matching
func match() {
	fmt.Println("=== Match Patterns ===")

	re := regexp.MustCompile(`\d{3}-\d{4}`)

	// MatchString: simple true/false
	fmt.Printf("Match '123-4567': %v\n", re.MatchString("123-4567"))
	fmt.Printf("Match 'abc-defg': %v\n", re.MatchString("abc-defg"))
	fmt.Println()
}

// find demonstrates regex finding
func find() {
	fmt.Println("=== Find Patterns ===")

	text := "Contact: 555-1234 or 555-5678"
	re := regexp.MustCompile(`\d{3}-\d{4}`)

	// FindString: first match
	first := re.FindString(text)
	fmt.Printf("First match: %s\n", first)

	// FindAllString: all matches
	all := re.FindAllString(text, -1)
	fmt.Printf("All matches: %v\n", all)

	// FindStringSubmatch: with capture groups
	re2 := regexp.MustCompile(`(\d{3})-(\d{4})`)
	matches := re2.FindStringSubmatch("Phone: 555-1234")
	fmt.Printf("Submatch: %v\n", matches) // [full, group1, group2]
	fmt.Println()
}

// replace demonstrates regex replacement
func replace() {
	fmt.Println("=== Replace Patterns ===")

	text := "My email is user@example.com and backup is admin@test.org"
	re := regexp.MustCompile(`\w+@\w+\.\w+`)

	// ReplaceAllString: replace all occurrences
	redacted := re.ReplaceAllString(text, "[REDACTED]")
	fmt.Printf("Redacted: %s\n", redacted)

	// ReplaceAllStringFunc: custom replacement logic
	re2 := regexp.MustCompile(`\d+`)
	doubled := re2.ReplaceAllStringFunc("I have 5 apples and 10 oranges", func(s string) string {
		n, _ := strconv.Atoi(s)
		return strconv.Itoa(n * 2)
	})
	fmt.Printf("Doubled: %s\n", doubled)
	fmt.Println()
}

func main() {
	compile_vs_must_compile()
	match()
	find()
	replace()
}
