package main

import "io"

type codecMode int

const (
	codecModeNone codecMode = iota
	codecModeEncode
	codecModeDecode
)

type copyCodec func(io.Writer, io.Reader) error

func (m codecMode) copyFuns() (copyCodec, copyCodec) {
	switch m {
	case codecModeEncode:
		return copyEncode, copyDecode
	case codecModeDecode:
		return copyDecode, copyEncode
	default:
		f := func(dst io.Writer, src io.Reader) error {
			_, err := io.Copy(dst, src)
			return err
		}
		return f, f
	}
}
