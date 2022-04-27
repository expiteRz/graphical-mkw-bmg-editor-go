package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/expiteRz/graphical-mkw-bmg-editor-go/delve"
	"os"
	"unicode/utf16"
	"unsafe"
)

var (
	N        bool = true // Initialize it as true when launching up the app
	C        bool
	Filepath string

	H Header
	I IndexTable
	P StringPool
	M MsgId
)

var (
	S             = CharsetString[H.Charset]
	CurrentOffset uint32
	NextOffset    uint32
	DecodedPool   []string
)

func init() {
	InitBmg()
}

func InitBmg() {
	Filepath = ""

	H = Header{
		Magic:     [8]byte{0x4D, 0x45, 0x53, 0x47, 0x62, 0x6D, 0x67, 0x31}, // MESGbmg1
		NumBlocks: 3,
		Charset:   2,
	}
	I = IndexTable{
		Kind:       [4]byte{0x49, 0x4E, 0x46, 0x31}, // INF1
		Entries:    1,
		EntrySize:  8,
		Group:      0,
		DefColor:   0,
		Reserved:   0,
		MsgEntries: []MsgEntry{{Offset: 2, FontType: 0x01}},
	}
	P = StringPool{
		Magic: [4]byte{0x44, 0x41, 0x54, 0x31}, // DAT1
		Pool:  nil,
	}
	M = MsgId{
		Kind:    [4]byte{0x4d, 0x49, 0x44, 0x31},
		Entries: 1,
		Format:  0x10,
		Info:    0,
		Ids:     []uint32{0},
	}

	I.Size = uint32(unsafe.Sizeof(I))
	P.Size = uint32(unsafe.Sizeof(P))
	M.Size = uint32(unsafe.Sizeof(M))
	H.DataSize =
		uint32(unsafe.Sizeof(H)) + I.Size + P.Size + M.Size

	return
}

func ParseBmg() error {
	// Open the file
	file, err := os.Open(Filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Start to parse BMG
	// Parse main header at first
	binary.Read(file, binary.BigEndian, &H)

	// Parse first 8 bytes INF1
	binary.Read(file, binary.BigEndian, &I.Kind)
	binary.Read(file, binary.BigEndian, &I.Size)

	// Make temporarily byte array
	reserving := make([]byte, I.Size-8)
	// Read bytes with length of the above one
	binary.Read(file, binary.BigEndian, reserving)
	// Make new reader
	r := bytes.NewReader(reserving)
	// Continue parsing INF1
	binary.Read(r, binary.BigEndian, &I.Entries)
	binary.Read(r, binary.BigEndian, &I.EntrySize)
	binary.Read(r, binary.BigEndian, &I.Group)
	binary.Read(r, binary.BigEndian, &I.DefColor)
	binary.Read(r, binary.BigEndian, &I.Reserved)

	entries := make([]MsgEntry, I.Entries)

	binary.Read(r, binary.BigEndian, &entries)
	I.MsgEntries = entries

	// WIP: Parse DAT1
	binary.Read(file, binary.BigEndian, &P.Magic)
	binary.Read(file, binary.BigEndian, &P.Size)

	pool := make([]uint16, (P.Size-8)/2)
	binary.Read(file, binary.BigEndian, &pool)
	P.Pool = pool

	// WIP: Parse MID1
	binary.Read(file, binary.BigEndian, &M.Kind)
	binary.Read(file, binary.BigEndian, &M.Size)
	binary.Read(file, binary.BigEndian, &M.Entries)
	binary.Read(file, binary.BigEndian, &M.Format)
	binary.Read(file, binary.BigEndian, &M.Info)
	binary.Read(file, binary.BigEndian, &M.Reserved)

	ids := make([]uint32, M.Entries)
	binary.Read(file, binary.BigEndian, &ids)
	M.Ids = ids

	if delve.Enabled {
		fmt.Println(M.Ids)
	}

	return nil
}

// CombineBmg combines all structs and convert to byte array
func CombineBmg() (bytes.Buffer, error) {
	var (
		buf = bytes.Buffer{}
		err error
	)

	/// Check if any length dividable with 16
	// If not, append data
	var tmp bytes.Buffer // Initially define buffer

	// Text Index Table.messageEntries
	if err = binary.Write(&tmp, binary.BigEndian, I.MsgEntries); err != nil {
		return bytes.Buffer{}, err
	}
	if len(tmp.Bytes())%16 != 0 {
		I.MsgEntries = append(I.MsgEntries, *new(MsgEntry))
	}

	// String Pool
	tmp = bytes.Buffer{}
	if err = binary.Write(&tmp, binary.BigEndian, P.Pool); err != nil {
		return bytes.Buffer{}, err
	}
	surp := len(tmp.Bytes()) % 16
	if surp != 0 {
		remind := surp / 2
		P.Pool = append(P.Pool, make([]uint16, remind)...)
	}

	// Message IDs
	tmp = bytes.Buffer{}
	if err = binary.Write(&tmp, binary.BigEndian, M.Ids); err != nil {
		return bytes.Buffer{}, err
	}
	surp = len(tmp.Bytes()) % 16
	if surp != 0 {
		remind := surp / 4
		M.Ids = append(M.Ids, make([]uint32, remind)...)
	}

	/// Start to write structs into a byte array below

	// Header
	err = binary.Write(&buf, binary.BigEndian, H)
	if err != nil {
		return bytes.Buffer{}, err
	}
	if delve.Enabled {
		fmt.Println(buf.Bytes())
	}

	// Text Index Table
	err = binary.Write(&buf, binary.BigEndian, I.Kind)
	err = binary.Write(&buf, binary.BigEndian, I.Size)
	err = binary.Write(&buf, binary.BigEndian, I.Entries)
	err = binary.Write(&buf, binary.BigEndian, I.EntrySize)
	err = binary.Write(&buf, binary.BigEndian, I.Group)
	err = binary.Write(&buf, binary.BigEndian, I.DefColor)
	err = binary.Write(&buf, binary.BigEndian, I.Reserved)
	err = binary.Write(&buf, binary.BigEndian, I.MsgEntries)
	if err != nil {
		return bytes.Buffer{}, err
	}
	if delve.Enabled {
		fmt.Println(buf.Bytes())
	}

	// String Pool
	err = binary.Write(&buf, binary.BigEndian, P.Magic)
	err = binary.Write(&buf, binary.BigEndian, P.Size)
	err = binary.Write(&buf, binary.BigEndian, P.Pool)
	if err != nil {
		return bytes.Buffer{}, err
	}
	if delve.Enabled {
		fmt.Println(buf.Bytes())
	}

	// Message IDs
	err = binary.Write(&buf, binary.BigEndian, M.Kind)
	err = binary.Write(&buf, binary.BigEndian, M.Size)
	err = binary.Write(&buf, binary.BigEndian, M.Entries)
	err = binary.Write(&buf, binary.BigEndian, M.Format)
	err = binary.Write(&buf, binary.BigEndian, M.Info)
	err = binary.Write(&buf, binary.BigEndian, M.Reserved)
	err = binary.Write(&buf, binary.BigEndian, M.Ids)
	if err != nil {
		return bytes.Buffer{}, err
	}
	if delve.Enabled {
		fmt.Println(buf.Bytes())
	}

	return buf, nil
}

// getMessage tries to get texts from string pool bytes between the defined address and the address in INF1
func getMessage(p uint32, n uint32) string { // TODO: How to get next offset if it's end?
	//* Technically the string is null then no assignment *//
	if p <= 0 {
		CurrentOffset = p
	}

	if n <= 0 {
		NextOffset = n
	}
	//* Until here *//

	if CurrentOffset == NextOffset { // This parsing is obviously meaningless. We can replace here condition with simply returning an empty string
		textBytes := P.Pool[0:2]
		DecodedPool := utf16.Decode(textBytes)

		return string(DecodedPool)
	}

	textBytes := P.Pool[CurrentOffset:NextOffset] // Get bytes
	decodedBytes := utf16.Decode(textBytes)       // Decode bytes to string

	return string(decodedBytes)
}
