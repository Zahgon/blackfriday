//
// Blackfriday Markdown Processor
// Available at http://github.com/russross/blackfriday
//
// Copyright © 2011 Russ Ross <russ@russross.com>.
// Distributed under the Simplified BSD License.
// See README.md for details.
//

//
//
// SmartyPants rendering
//
//

package blackfriday

import (
	"bytes"
)

type smartypantsData struct {
	inSingleQuote bool
	inDoubleQuote bool
}

func wordBoundary(c byte) bool { _ = "STUB: not implemented"; return false }

func tolower(c byte) byte { _ = "STUB: not implemented"; return 0 }

func isdigit(c byte) bool { _ = "STUB: not implemented"; return false }

func smartQuoteHelper(out *bytes.Buffer, previousChar byte, nextChar byte, quote byte, isOpen *bool, addNBSP bool) bool {
	_ = "STUB: not implemented"
	// edge of the buffer is likely to be a tag that we don't get to see,
	// so we treat it like text sometimes
	return false
}

// enumerate all sixteen possibilities for (previousChar, nextChar)
// each can be one of {0, space, punct, other}

// context is not any help here, so toggle

// [ "] might be [ "<code>foo...]

// [!"] hmm... could be [Run!"] or [("<code>...]

/* isnormal(previousChar) && */
// [a"] is probably a close

// [" ] might be [...foo</code>" ]

// [ " ] context is not any help here, so toggle

// [!" ] is probably a close

/* isnormal(previousChar) && */
// [a" ] this is one of the easy cases

// ["!] hmm... could be ["$1.95] or [</code>"!...]

// [ "!] looks more like [ "$1.95]

// [!"!] context is not any help here, so toggle

/* isnormal(previousChar) && */
// [a"!] is probably a close

/* && isnormal(nextChar) */
// ["a] is probably an open

/* && isnormal(nextChar) */
// [ "a] this is one of the easy cases

/* && isnormal(nextChar) */
// [!"a] is probably an open

// [a'b] maybe a contraction?

// Note that with the limited lookahead, this non-breaking
// space will also be appended to single double quotes.

func smartSingleQuote(out *bytes.Buffer, smrt *smartypantsData, previousChar byte, text []byte) int {
	_ = "STUB: not implemented"
	return 0
}

func smartParens(out *bytes.Buffer, smrt *smartypantsData, previousChar byte, text []byte) int {
	_ = "STUB: not implemented"
	return 0
}

func smartDash(out *bytes.Buffer, smrt *smartypantsData, previousChar byte, text []byte) int {
	_ = "STUB: not implemented"
	return 0
}

func smartDashLatex(out *bytes.Buffer, smrt *smartypantsData, previousChar byte, text []byte) int {
	_ = "STUB: not implemented"
	return 0
}

func smartAmpVariant(out *bytes.Buffer, smrt *smartypantsData, previousChar byte, text []byte, quote byte, addNBSP bool) int {
	_ = "STUB: not implemented"
	return 0
}

func smartAmp(angledQuotes, addNBSP bool) func(out *bytes.Buffer, smrt *smartypantsData, previousChar byte, text []byte) int {
	_ = "STUB: not implemented"
	return nil
}

func smartPeriod(out *bytes.Buffer, smrt *smartypantsData, previousChar byte, text []byte) int {
	_ = "STUB: not implemented"
	return 0
}

func smartBacktick(out *bytes.Buffer, smrt *smartypantsData, previousChar byte, text []byte) int {
	_ = "STUB: not implemented"
	return 0
}

func smartNumberGeneric(out *bytes.Buffer, smrt *smartypantsData, previousChar byte, text []byte) int {
	_ = "STUB: not implemented"
	return 0
}

// is it of the form digits/digits(word boundary)?, i.e., \d+/\d+\b
// note: check for regular slash (/) or fraction slash (⁄, 0x2044, or 0xe2 81 84 in utf-8)
//       and avoid changing dates like 1/23/2005 into fractions.

func smartNumber(out *bytes.Buffer, smrt *smartypantsData, previousChar byte, text []byte) int {
	_ = "STUB: not implemented"
	return 0
}

func smartDoubleQuoteVariant(out *bytes.Buffer, smrt *smartypantsData, previousChar byte, text []byte, quote byte) int {
	_ = "STUB: not implemented"
	return 0
}

func smartDoubleQuote(out *bytes.Buffer, smrt *smartypantsData, previousChar byte, text []byte) int {
	_ = "STUB: not implemented"
	return 0
}

func smartAngledDoubleQuote(out *bytes.Buffer, smrt *smartypantsData, previousChar byte, text []byte) int {
	_ = "STUB: not implemented"
	return 0
}

func smartLeftAngle(out *bytes.Buffer, smrt *smartypantsData, previousChar byte, text []byte) int {
	_ = "STUB: not implemented"
	return 0
}

type smartCallback func(out *bytes.Buffer, smrt *smartypantsData, previousChar byte, text []byte) int

type smartypantsRenderer [256]smartCallback

var (
	smartAmpAngled      = smartAmp(true, false)
	smartAmpAngledNBSP  = smartAmp(true, true)
	smartAmpRegular     = smartAmp(false, false)
	smartAmpRegularNBSP = smartAmp(false, true)
)

func smartypants(flags int) *smartypantsRenderer { _ = "STUB: not implemented"; return nil }
