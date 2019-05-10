package main

import "flag"

type options struct {
	laddr     string
	taddr     string
	codecMode codecMode
	codecKey  string
}

func parseOptions() *options {
	var (
		laddr        = flag.String("laddr", ":7000", "listen address")
		taddr        = flag.String("taddr", ":7001", "target address")
		codecModeInt = flag.Int("mode", 0, "codec mode {none: 0,  encode: 1, decode: 2}")
		codeckey     = flag.String("key", "", "codec key")
	)

	flag.Parse()

	if len(*laddr) == 0 || len(*taddr) == 0 {
		flag.Usage()
		return nil
	}

	mode := codecMode(*codecModeInt)

	switch mode {
	case codecModeNone:
		// if len(*codeckey) > 0 {
		// 	flag.Usage()
		// 	return nil
		// }

	case codecModeEncode:
		fallthrough
	case codecModeDecode:
		if len(*codeckey) == 0 {
			flag.Usage()
			return nil
		}
	default:
		flag.Usage()
		return nil
	}

	return &options{
		*laddr, *taddr, mode, *codeckey,
	}
}
