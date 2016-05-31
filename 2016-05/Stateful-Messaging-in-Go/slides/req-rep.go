context, _ := zmq.NewContext()
socket, _ := context.NewSocket(zmq.REQ)

socket.Connect("tcp://localhost:5555")

socket.Send([]byte("Hello"), 0)

reply, _ := socket.Recv(0)
println("Received ", string(reply))

