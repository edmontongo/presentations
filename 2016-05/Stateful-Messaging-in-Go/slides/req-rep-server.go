context, _ := zmq.NewContext()
socket, _ := context.NewSocket(zmq.REP)

socket.Bind("tcp://*:5555")

msg, _ := socket.Recv(0)
socket.Send([]byte("World"), 0)
