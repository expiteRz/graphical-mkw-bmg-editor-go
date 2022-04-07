package main

import (
	"encoding/binary"
	"fmt"
	"github.com/expiteRz/graphical-mkw-bmg-editor-go/utils"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./Common.bmg")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer file.Close()

	h := utils.Header{}

	// Header
	err = binary.Read(file, binary.BigEndian, &h.Magic)
	err = binary.Read(file, binary.BigEndian, &h.DataSize)
	err = binary.Read(file, binary.BigEndian, &h.NumBlocks)
	err = binary.Read(file, binary.BigEndian, &h.Charset)
	err = binary.Read(file, binary.BigEndian, &h.Reserved)

	i := utils.IndexTable{}

	//Text Index Table
	binary.Read(file, binary.BigEndian, &i.Kind)
	binary.Read(file, binary.BigEndian, &i.Size)
	binary.Read(file, binary.BigEndian, &i.Entries)
	binary.Read(file, binary.BigEndian, &i.EntrySize)
	binary.Read(file, binary.BigEndian, &i.Group)
	binary.Read(file, binary.BigEndian, &i.DefColor)
	binary.Read(file, binary.BigEndian, &i.Reserved)

	num := i.Entries
	entries := make([]utils.MsgEntry, num)

	err = binary.Read(file, binary.BigEndian, &entries)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Magic: %s\n", h.Magic)
	fmt.Printf("Size: %d\n", h.DataSize)
	switch h.Charset {
	case 0:
		fmt.Printf("Charset: %s\n", "Undefined")
	case 1:
		fmt.Printf("Charset: %s\n", "CP1252")
	case 2:
		fmt.Printf("Charset: %s\n", "UTF16")
	case 3:
		fmt.Printf("Charset: %s\n", "SJIS")
	case 4:
		fmt.Printf("Charset: %s\n", "UTF8")
	default:
		fmt.Printf("Charset: %s\n", "Unknown")
	}

	fmt.Println(entries)
}
