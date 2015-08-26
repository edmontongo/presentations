package main

import (
	"flag"
	"fmt"
	"os"
	"net"
	"github.com/fatih/color"
	"github.com/getwe/figlet4go"
)


func main() {
	flag_str := flag.String("str", "Edmonton  Go", "input string")
	flag.Usage = func() {
		fmt.Println("usage: figlet_nc [-str <string>]  <hostname> <port>")
		os.Exit(0)
	}
	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
	}

	conn, err := net.Dial("tcp", flag.Arg(0) + ":" + flag.Arg(1))
	if err != nil {
		fmt.Printf("dial error: %s\n", err)
		os.Exit(-1)
	}

	str := *flag_str
	ascii := figlet4go.NewAsciiRender()

	// change the font color
	colors := [...]color.Attribute{
		color.FgMagenta,
		color.FgYellow,
		color.FgBlue,
		color.FgCyan,
		color.FgRed,
		color.FgWhite,
	}
	options := figlet4go.NewRenderOptions()
	options.FontColor = make([]color.Attribute, len(str))
	for i := range options.FontColor {
		options.FontColor[i] = colors[i%len(colors)]
	}
	renderStr, _ := ascii.RenderOpts(str, options)
	conn.Write([]byte(renderStr))
	fmt.Println("Wrote '" + str + "' to ", conn.RemoteAddr())
}
