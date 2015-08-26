## Simple netcat clone

The `nc.go` listens on a TCP port, the `figlet_netcat.go` sends a string over a TCP
connection to a remote listener.

This is based on [go netcat package](https://github.com/codeskyblue/netcat) and [go figlet](https://github.com/getwe/figlet4go).

To run the examples in this directory,
```
> source source.me           # sets GOPATH
> go get                     # fetch the github packages
> go run nc.go               # displays help menu
> go run figlet_netcat.go    # displays help menu
```
