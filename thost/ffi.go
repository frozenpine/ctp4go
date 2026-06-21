package thost

import (
	"bytes"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func BytesToGBK(bs []byte) string {
	msg, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(bytes.Split(bs, []byte("\x00"))[0])
	return string(msg)
}

func BytesToString(bs []byte) string {
	return string(bytes.Split(bs, []byte("\x00"))[0])
}
