// Structures are defined based on CT Wiiki (https://wiki.tockdom.com/wiki/BMG_(File_Format))

package utils

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
	Magic     [8]byte // MESGbmg1 in ASCII
	DataSize  uint32
	NumBlocks uint32 // Usually 3, sometimes 2 outside from MKW
	Charset   Charset
	Reserved  [15]uint8
}

type IndexTable struct {
	Kind       [4]byte // INF1 in ASCII
	Size       uint32
	Entries    uint16 // Number of messages
	EntrySize  uint16 // Always 8 in MKW
	Group      uint16 // BMG file ID, usually 0
	DefColor   uint8
	Reserved   int8
	MsgEntries []MsgEntry
}

type MsgEntry struct {
	Offset uint32
	// **0x00** for texts related countdown and final race strings
	//
	// **0x01** for standard (default)
	//
	// **0x04** for red color for battle and team match
	//
	// **0x05** for blue color for battle and team match
	FontType uint8
	Padding  [3]uint8
}

type StringPool struct {
	Magic [4]byte // DAT1 in ASCII
	Size  uint32
	Pool  []rune
}

type MsgId struct {
	Kind     [4]byte // MID1 in ASCII
	Size     uint32
	Entries  uint16 // Number of messages
	Format   uint8
	Info     uint8
	Reserved int32    // Usually 0
	Ids      []uint16 // Decode strings with `unicode/utf16.Decode()`
}
