// Child error
childError := errortree.Add(nil, "test0", errors.New("child error"))

// Another, on the same level
childError = errortree.Add(childError, "test1", errors.New("another child error"))

// Top-level error
err := errortree.Add(nil, "child", childError) // HL

fmt.Printf("%s\n", err)
