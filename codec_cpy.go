package main

import (
	"bufio"
	"encoding/binary"
	"io"
)

func copyEncode(dst io.Writer, src io.Reader) error {
	bf := make([]byte, 32*1024)
	for {
		n, err := src.Read(bf)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		if n == 0 {
			continue
		}

		b1, err := encrypt(bf[:n])
		if err != nil {
			return err
		}

		b2 := make([]byte, binary.MaxVarintLen32+len(b1))

		l := binary.PutUvarint(b2, uint64(len(b1)))
		b2 = append(b2[:l], b1...)

		for len(b2) > 0 {
			n, err := dst.Write(b2)
			if err != nil {
				return err
			}
			b2 = b2[n:]
		}
	}
}

func copyDecode(dst io.Writer, src io.Reader) error {
	rd := bufio.NewReader(src)
	for {
		n, err := binary.ReadUvarint(rd)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		if n == 0 {
			continue
		}

		bf := make([]byte, n)

		_, err = io.ReadFull(rd, bf)
		if err != nil {
			if err == io.EOF {
				err = io.ErrUnexpectedEOF
			}
			return err
		}

		bf, err = decrypt(bf)
		if err != nil {
			return err
		}

		for len(bf) > 0 {
			n, err := dst.Write(bf)
			if err != nil {
				return err
			}
			bf = bf[n:]
		}
	}
}
