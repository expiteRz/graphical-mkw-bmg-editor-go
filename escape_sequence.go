package main

import (
	"errors"
	"github.com/expiteRz/graphical-mkw-bmg-editor-go/utils"
)

func getCurrentPlayer() utils.EscapeSeq {
	return []byte{0, 0x1a, 0, 0x6, 0, 0x2, 0, 0, 0, 0}
}

func getFontSize(v ...byte) (utils.EscapeSeq, error) {
	base := []byte{0, 0x1a, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0}

	if len(v) < 2 {
		base = append(base, 0, v[0])
		return base, nil
	} else if len(v) < 1 {
		base = append(base, 0, 0)
		return base, nil
	} else if len(v) > 2 {
		return nil, errors.New("values are too much")
	}

	base = append(base, v...)
	return base, nil
}

type FontColor int

const (
	gray        FontColor = 0
	white       FontColor = 2
	red         FontColor = 40 + 20 + 32
	yellow      FontColor = 30
	green       FontColor = 33
	blue        FontColor = 21 + 31
	transparent FontColor = 8
)

func getFontColor(v FontColor) utils.EscapeSeq {
	return []byte{0, 0x1a, 0, 8, 0, 0, 0, 0, 0, 1, 0, 0, interface{}(v).(uint8)}
}
