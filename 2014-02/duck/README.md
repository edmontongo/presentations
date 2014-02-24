## Duck Hunt: Interfaces in Go

* [Laughing dog](http://www.youtube.com/watch?v=g1QCbXCezNc) (video)
* [Interfaces: A new leaf in an old book](http://video.fosdem.org/2014/K4601/Sunday/Interfaces_a_new_leaf_for_an_old_book.webm) (video)
* [Effective Go](http://golang.org/doc/effective_go.html#interfaces)
* [Built In Interfaces](http://jmoiron.net/blog/built-in-interfaces/)
* [Go Data Structures: Interfaces](http://research.swtch.com/interfaces)
* [Stunned by Go](http://how-bazaar.blogspot.co.nz/2013/07/stunned-by-go.html)

## Notes

* types & methods
* duck typing, care about behaviour

`first.go` 

* A duck and a dog that `Speak()`.

`second.go`

* Implementing `IsDuck()`. 
* (In a dynamic language, just pass in any object) A Speaker interface.
* Actually a good thing: don't need to hunt for your duck types.

`third.go`

* ClayPigeon does not implement Speaker.

`fourth.go`

* Rather silly example. Type assertions.

`fifth.go`

* Empty interfaces.

`stack.go`

Empty interfaces in a slightly more realistic example.

`second-stack.go`

What happens when we put a duck on the stack? Type switches.


