sock, _ := req.NewSocket()
sock.AddTransport(tcp.NewTransport())

sock.Dial("tcp://127.0.0.1:40899")
sock.Send([]byte("Hello"))

reply, _ := sock.Recv()
println("Client received ", string(reply))
