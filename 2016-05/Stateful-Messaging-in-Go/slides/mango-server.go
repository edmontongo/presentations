sock, _ := rep.NewSocket()
sock.AddTransport(tcp.NewTransport())
sock.Listen("tcp://127.0.0.1:40899")

msg, _ := sock.Recv()
println("Server received ", string(msg))

sock.Send([]byte("World"))
