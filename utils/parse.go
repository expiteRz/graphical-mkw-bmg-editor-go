package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
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
)

func init() {
	H = Header{}
	I = IndexTable{MsgEntries: []MsgEntry{
		{Offset: 2, FontType: 0x01},
	}}
	M = MsgId{
		Kind:    [4]byte{0x4d, 0x49, 0x44, 0x31},
		Size:    16,
		Entries: 1,
		Format:  0x10,
		Info:    0,
		Ids:     []uint32{0},
	}
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

	fmt.Println(M.Ids)

	return nil
}

func lenText(p uint32, n uint32) uint {
	return uint(n - p)
}
