package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	listen := flag.Bool("l", false, "listen mode")
	flag.Usage = func() {
		fmt.Println(
`usage:	nc hostname port  - connect to hostname:port
	nc -l port        - listen on port`)
		os.Exit(0)
	}
	flag.Parse()

	if *listen {	/* server */
		if flag.NArg() != 1 {
			flag.Usage()
		}
		lis, err := net.Listen("tcp", ":" + flag.Arg(0))
		if err != nil {
			log.Fatalf("listen error: %s", err)
		}
		log.Printf("listening on %s", flag.Arg(0))
		for {
			conn, err := lis.Accept()
			if err != nil {
				log.Fatalf("Accept error: %s", err)
			}
			log.Printf("accept from %s", conn.RemoteAddr())

			io.Copy(os.Stdout, conn)
		}
	} else {	/* client */
		if flag.NArg() != 2 {
			flag.Usage()
		}

		conn, err := net.Dial("tcp", flag.Arg(0) + ":" + flag.Arg(1))
		if err != nil {
			log.Fatalf("dial error: %s", err)
		}
		io.Copy(conn, os.Stdin)
	}
}
