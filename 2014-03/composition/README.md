## Composition & Delegation

* manual delegation
* embedding
    * one-way, person type doesn't know about employee (no polymorphism here)
* ambiguous selector (compile time)
    * only if we use Name()
    * resolve by implementing `Name()` on Employee
