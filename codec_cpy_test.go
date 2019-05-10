package main

import (
	"bytes"
	"testing"
)

func TestCopyEncodeDecode(t *testing.T) {
	var err error

	err = initAes("abc")
	if err != nil {
		t.Fail()
	}

	src := bytes.NewBuffer(nil)
	dst := bytes.NewBuffer(nil)
	dst2 := bytes.NewBuffer(nil)

	src.WriteByte(0)
	src.WriteByte(1)
	src.WriteByte(2)

	err = copyEncode(dst, src)
	if err != nil {
		t.Fail()
	}

	err = copyDecode(dst2, dst)
	if err != nil {
		t.Fail()
	}

	bs := dst2.Bytes()
	t.Logf("len(bs) = %d\n", len(bs))

	if len(bs) != 3 {
		t.Fail()
	}

	for i := 0; i < len(bs); i++ {
		t.Log(bs[i])
		if int(bs[i]) != i {
			t.Fail()
		}
	}
}
