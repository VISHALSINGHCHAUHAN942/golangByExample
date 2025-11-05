package main

import (
	"fmt"
	"unicode/utf8"
)

/*
------------------------ RUNE vs STRING IN GO ------------------------

A Go string is a sequence of **bytes**, not characters.
UTF-8 encoding uses **variable length encoding**:
- English letters → 1 byte
- Hindi characters → 3 bytes each
- Emojis → 4 bytes

So, the Hindi word: "कबूतर" looks like 5 characters,
but in memory it takes 15 bytes.

क → E0 A4 95 (3 bytes)
ब → E0 A4 AC (3 bytes)
ू → E0 A5 82 (3 bytes)
त → E0 A4 A4 (3 bytes)
र → E0 A4 B0 (3 bytes)

If we loop using ordinary indexing (i++), we read **1 byte at a time** → characters break → WRONG ❌
If we loop using runes, we read **a full Unicode character** → CORRECT ✅

utf8.DecodeRuneInString(s[i:]) helps us:
- runeValue → the actual character
- width → how many bytes to move forward

We then do: i += width   (move to next full character)

This ensures each loop step moves to the **start of the next character**, not the next byte.
----------------------------------------------------------------------
*/

func main() {

	s := "कबूतर"

	fmt.Println("String:", s)
	fmt.Println("Bytes length (len):", len(s))
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// ❌ Method 1: Reading byte-by-byte (WRONG)
	fmt.Println("\nReading byte-by-byte (Incorrect):")
	for i := 0; i < len(s); i++ {
		// Here we print single bytes, so multi-byte characters break.
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()

	// ✅ Method 2: Reading rune-by-rune (Correct)
	fmt.Println("\nReading rune-by-rune (Correct):")
	for _, r := range s {
		// Here r is a rune → one complete Unicode character.
		// %c prints the full character represented by the rune.
		fmt.Printf("%c ", r)
	}
	fmt.Println()

	// ✅ Method 3: Manual rune iteration using utf8.DecodeRuneInString
	fmt.Println("\nUsing DecodeRuneInString (manual rune iteration):")
	for i, w := 0, 0; i < len(s); i += w {
		/*
		   DecodeRuneInString returns:
		   - runeValue → the full Unicode character at position i
		   - width → how many bytes this character occupies
		*/
		runeValue, width := utf8.DecodeRuneInString(s[i:])

		fmt.Printf("%#U at byte index %d (uses %d bytes)\n", runeValue, i, width)

		// Move index forward by width so next iteration jumps correctly
		w = width

		// Check the rune with custom logic
		examineRune(runeValue)
	}
}

// examineRune demonstrates working with full Unicode characters
func examineRune(r rune) {
	if r == 'त' {
		fmt.Println("→ Found character: त")
	}
}
