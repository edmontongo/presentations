package main

import "fmt"

type Request struct { // HL
	key, val string     // used for get/add
	response chan Entry // nil for add
} // HL

type Entry struct {
	Val string // value corresponding to @key
	Ok  bool   // if @key is in map
}

// GoHash is a hash managed by a monitor goroutine // HL
type GoHash struct {
	m        map[string]string // internal
	requests chan Request      // external
}

// ENDDEF OMIT

func (gh *GoHash) Add(key, val string) {
	gh.requests <- Request{key: key, val: val} // HL
}

func (gh *GoHash) Get(key string) Entry {
	var responseChan = make(chan Entry)
	gh.requests <- Request{key: key, response: responseChan} // HL
	return <-responseChan
}

func NewGoHash() *GoHash {
	var g = GoHash{m: make(map[string]string), requests: make(chan Request)}

	go g.loop() // run until closed // HL
	return &g
}

// ENDMETH OMIT

func (gh *GoHash) loop() { // HL
	var response chan Entry // activated on demand // HL
	var e Entry
	for {
		select {
		case r, ok := <-gh.requests: // HL
			if !ok {
				return
			}
			if r.response == nil {
				gh.m[r.key] = r.val
			} else {
				e.Val, e.Ok = gh.m[r.key]
				response = r.response
			}
		case response <- e: // HL
			response = nil
		}
	}
} // HL

func (gh *GoHash) Close() { close(gh.requests) } // HL

// ENDLOOP OMIT

func main() {
	var gh = NewGoHash() // HL

	gh.Add("uno", "1")
	gh.Add("due", "2")
	gh.Add("tre", "3")

	for _, k := range []string{"tre", "uno", "due", "quattro"} {
		fmt.Printf("%s: %s\n", k, gh.Get(k))
	}
	// ENDMAIN OMIT
}

func (e Entry) String() string {
	if !e.Ok {
		return "(entry not found)"
	} else {
		return e.Val
	}
}
