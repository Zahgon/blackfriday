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
// LaTeX rendering backend
//
//

package blackfriday

import (
	"bytes"
)

// Latex is a type that implements the Renderer interface for LaTeX output.
//
// Do not create this directly, instead use the LatexRenderer function.
type Latex struct {
}

// LatexRenderer creates and configures a Latex object, which
// satisfies the Renderer interface.
//
// flags is a set of LATEX_* options ORed together (currently no such options
// are defined).
func LatexRenderer(flags int) Renderer { _ = "STUB: not implemented"; return *new(Renderer) }

func (options *Latex) GetFlags() int {
	_ = "STUB: not implemented"

	// render code chunks using verbatim, or listings if we have a language
	return 0
}

func (options *Latex) BlockCode(out *bytes.Buffer, text []byte, info string) {
	_ = "STUB: not implemented"
	return
}

func (options *Latex) TitleBlock(out *bytes.Buffer, text []byte) { _ = "STUB: not implemented"; return }

func (options *Latex) BlockQuote(out *bytes.Buffer, text []byte) { _ = "STUB: not implemented"; return }

func (options *Latex) BlockHtml(out *bytes.Buffer, text []byte) {
	_ = "STUB: not implemented"
	// a pretty lame thing to do...
	return
}

func (options *Latex) Header(out *bytes.Buffer, text func() bool, level int, id string) {
	_ = "STUB: not implemented"
	return
}

func (options *Latex) HRule(out *bytes.Buffer) { _ = "STUB: not implemented"; return }

func (options *Latex) List(out *bytes.Buffer, text func() bool, flags int) {
	_ = "STUB: not implemented"
	return
}

func (options *Latex) ListItem(out *bytes.Buffer, text []byte, flags int) {
	_ = "STUB: not implemented"
	return
}

func (options *Latex) Paragraph(out *bytes.Buffer, text func() bool) {
	_ = "STUB: not implemented"
	return
}

func (options *Latex) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {
	_ = "STUB: not implemented"
	return
}

func (options *Latex) TableRow(out *bytes.Buffer, text []byte) { _ = "STUB: not implemented"; return }

func (options *Latex) TableHeaderCell(out *bytes.Buffer, text []byte, align int) {
	_ = "STUB: not implemented"
	return
}

func (options *Latex) TableCell(out *bytes.Buffer, text []byte, align int) {
	_ = "STUB: not implemented"
	return
}

// TODO: this
func (options *Latex) Footnotes(out *bytes.Buffer, text func() bool) {
	_ = "STUB: not implemented"
	return
}

func (options *Latex) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {
	_ = "STUB: not implemented"
	return
}

func (options *Latex) AutoLink(out *bytes.Buffer, link []byte, kind int) {
	_ = "STUB: not implemented"
	return
}

func (options *Latex) CodeSpan(out *bytes.Buffer, text []byte) { _ = "STUB: not implemented"; return }

func (options *Latex) DoubleEmphasis(out *bytes.Buffer, text []byte) {
	_ = "STUB: not implemented"
	return
}

func (options *Latex) Emphasis(out *bytes.Buffer, text []byte) { _ = "STUB: not implemented"; return }

func (options *Latex) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
	_ = "STUB: not implemented"
	return
}

// treat it like a link

func (options *Latex) LineBreak(out *bytes.Buffer) { _ = "STUB: not implemented"; return }

func (options *Latex) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {
	_ = "STUB: not implemented"
	return
}

func (options *Latex) RawHtmlTag(out *bytes.Buffer, tag []byte) { _ = "STUB: not implemented"; return }

func (options *Latex) TripleEmphasis(out *bytes.Buffer, text []byte) {
	_ = "STUB: not implemented"
	return
}

func (options *Latex) StrikeThrough(out *bytes.Buffer, text []byte) {
	_ = "STUB: not implemented"
	return
}

// TODO: this
func (options *Latex) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {
	_ = "STUB: not implemented"
	return
}

func needsBackslash(c byte) bool { _ = "STUB: not implemented"; return false }

func escapeSpecialChars(out *bytes.Buffer, text []byte) { _ = "STUB: not implemented"; return }

// directly copy normal characters

// escape a character

func (options *Latex) Entity(out *bytes.Buffer, entity []byte) {
	_ = "STUB: not implemented"
	// TODO: convert this into a unicode character or something
	return
}

func (options *Latex) NormalText(out *bytes.Buffer, text []byte) { _ = "STUB: not implemented"; return }

// header and footer
func (options *Latex) DocumentHeader(out *bytes.Buffer) { _ = "STUB: not implemented"; return }

func (options *Latex) DocumentFooter(out *bytes.Buffer) { _ = "STUB: not implemented"; return }
