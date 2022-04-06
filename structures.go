// Structures are defined based on CT Wiiki (https://wiki.tockdom.com/wiki/BMG_(File_Format))

package main

type Charset byte

const (
	Undefined Charset = iota
	CP1252
	UTF16 // For MKW
	ShiftJIS
	UTF8
)

type EscapeSeq []byte

type Header struct {
	magic     [8]byte // MESGbmg1 in ASCII
	dataSize  uint32
	numBlocks uint32 // Usually 3, sometimes 2 outside from MKW
	charset   Charset
	reserved  [15]uint8
	userWork  int
}

type IndexTable struct {
	kind       [4]byte // INF1 in ASCII
	size       uint32
	entries    uint16 // Number of messages
	entrySize  uint16 // Always 8 in MKW
	group      uint16 // BMG file ID, usually 0
	defColor   uint8
	reserved   int8
	msgEntries []MsgEntry
}

type MsgEntry struct {
	offset uint32
	// **0x00** for texts related countdown and final race strings
	//
	// **0x01** for standard (default)
	//
	// **0x04** for red color for battle and team match
	//
	// **0x05** for blue color for battle and team match
	fontType uint8
	padding  [3]uint8
}

type StringPool struct {
	magic [4]byte // DAT1 in ASCII
	size  uint32
	pool  []rune
}

type MsgId struct {
	kind     [4]byte // MID1 in ASCII
	size     uint32
	entries  uint16 // Number of messages
	format   uint8
	info     uint8
	reserved int32    // Usually 0
	ids      []uint16 // Decode strings with `unicode/utf16.Decode()`
}
