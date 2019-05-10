package main

import (
	"net"
)

const (
	donePart = iota
	doneAll
)

func relay(c1 *net.TCPConn, taddr string, mode codecMode) {
	var c2 *net.TCPConn
	defer func() {
		c1.Close()
		if c2 != nil {
			c2.Close()
		}
	}()

	c, err := net.Dial("tcp", taddr)
	if err != nil {
		return
	}

	c2 = c.(*net.TCPConn)
	done := make(chan int, 2)

	fe, fd := mode.copyFuns()

	go cpy(c1, c2, done, fe)
	go cpy(c2, c1, done, fd)

	for i := 0; i < 2; i++ {
		v := <-done
		if v == doneAll {
			break
		}
	}
}

func cpy(c1, c2 *net.TCPConn, done chan<- int, doCpy copyCodec) {
	err := doCpy(c2, c1)
	if err != nil {
		done <- doneAll
		return
	}

	err = c2.CloseWrite()
	if err != nil {
		done <- doneAll
		return
	}
	done <- donePart
}
