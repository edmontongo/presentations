## Duck Hunt: Interfaces in Go

### Notes

* types & methods, implemented Stringer interface
* duck typing, behaviour, flexibility

`first.go` 

* A duck and a dog that `Speak()`.

`second.go`

* Implementing `IsDuck()`. 
* (In a dynamic language, just pass in any object) A Speaker interface.
* Actually a good thing: don't need to hunt for your duck types.

`third.go`

* ClayPigeon does not implement Speaker.
* Worth noting: interfaces apply across package boundaries.

`fourth.go`

* Rather silly example. Type assertions.

`fifth.go`

* Empty interfaces.

`stack.go`

* Empty interfaces in a slightly more realistic example.

`second-stack.go`

* What happens when we put a duck on the stack? Type switches.

### Resources

* [Laughing dog](http://www.youtube.com/watch?v=g1QCbXCezNc) (video)
* [Interfaces: A new leaf in an old book](http://video.fosdem.org/2014/K4601/Sunday/Interfaces_a_new_leaf_for_an_old_book.webm) (video)
* [Built In Interfaces](http://jmoiron.net/blog/built-in-interfaces/)
* [Effective Go](http://golang.org/doc/effective_go.html#interfaces)
* [Go Data Structures: Interfaces](http://research.swtch.com/interfaces)
* [Stunned by Go](http://how-bazaar.blogspot.co.nz/2013/07/stunned-by-go.html)
* [Top Ten Reasons Cocoa is Great Because of Objective-C](http://www.informit.com/articles/article.aspx?p=1353396) (#5 is duck typing)
