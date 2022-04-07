package main

import (
	"encoding/binary"
	"github.com/expiteRz/graphical-mkw-bmg-editor-go/utils"
	"os"
)

var (
	newFile  bool = true // Initialize it as true when launching up the app
	changed  bool
	filepath string

	header     utils.Header
	indexTable utils.IndexTable
)

var (
	s = utils.CharsetString[header.Charset]
)

func initVars() {
	header = utils.Header{}
	indexTable = utils.IndexTable{}
}

func parseBmg() error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err = binary.Read(file, binary.BigEndian, &header); err != nil {
		return err
	}
	binary.Read(file, binary.BigEndian, &indexTable.Kind)
	binary.Read(file, binary.BigEndian, &indexTable.Size)
	binary.Read(file, binary.BigEndian, &indexTable.Entries)
	binary.Read(file, binary.BigEndian, &indexTable.EntrySize)
	binary.Read(file, binary.BigEndian, &indexTable.Group)
	binary.Read(file, binary.BigEndian, &indexTable.DefColor)
	binary.Read(file, binary.BigEndian, &indexTable.Reserved)

	entries := make([]utils.MsgEntry, indexTable.Entries)

	binary.Read(file, binary.BigEndian, &entries)
	indexTable.MsgEntries = entries

	return nil
}
