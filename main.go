package main

import (
	"fmt"
	"net"
)

func main() {
	opts := parseOptions()
	if opts == nil {
		return
	}

	if len(opts.codecKey) > 0 {
		err := initAes(opts.codecKey)
		if err != nil {
			panic(err)
		}
	}

	ln, err := net.Listen("tcp", opts.laddr)
	if err != nil {
		panic(err)
	}
	lt := ln.(*net.TCPListener)

	for {
		conn, err := lt.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go relay(conn, opts.taddr, opts.codecMode)
	}
}
