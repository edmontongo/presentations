## Celsius

### Convert

Temperature conversion example.

`first.go` 

* main, float64, declaration, assignment, initialization
* format strings (value, %4.2f width.precision)
* function signature

`second.go`

* shortcut to initialize
* introduce user-defined types and type conversion
* demonstrate mismatched types and explain underlying types

`third.go`

* introduce methods and explicit receivers
* ToCelsius for completeness

`fourth.go`

* satisfying the Stringer interface
* stack overflow with %v

### Almanac

Aggregation functions on [Environment Canada data](http://edmonton.weatherstats.ca/download.html).

`almanac.go`

* structures, slices (variable length arrays), composite literals
* Climate type, methods on a slice
* name/value syntax for composite literals
* for range loops, underscore

## Transcript

Go's approach to object-oriented programming is a little different. In this episode, we will introduce types & methods into a temperature conversion tool. 

It was a balmy -18, so if we run this, we'll see that it was -0.4 degrees Fahrenheit.

The default formatting could be better. Instead of %v for value, let's specify %f for float, with a precision. Okay.

Next, let's add two user-defined types. Celsius and Fahrenheit, both which are floats. With that we can change our function to take a Celsius and return a Fahrenheit.

This doesn't work quite yet. We need to tell Go to convert the return value. If we pass -18 directly into the function, Go will infer the Celsius type for us, but because we are defining it beforehand, we need to convert it as well.

With our types defined, the compiler will prevent us from directly comparing or performing arithmetic on them.

Now let's change this function into a method - hanging off the Celsius type. Notice the explicit receiver 'c'. Our method is now within a namespace and we can use the familiar dot notation to call it.

Let's implement a ToCelsius method for completeness. Go is using static dispatch to call these methods, which means the compiler determines which method to call. At this point there is no distinction between "sending a message" and "calling a method".

Go achieves polymorphism through interfaces. For a *taste* of how this works, let's add a String() method. 

This seems a little magical. How does Printf know to call String?

It turns out that by adding a String method, we are implicitly satisfying the `Stringer` interface. At runtime, the `fmt` package can determine if our type implements String, and if so, call it. 

Here Go is using dynamic dispatch, as is the case with Objective-C. The String() message resolves to the method on our Celsius type at runtime.

Some care must be taken when implementing String(). If we use %v, the Sprintf function will call String(), which will call Sprintf, and so on until there is a stack overflow. One solution is to convert the value to its underlying type.

To wrap up, I want to point out that we can add methods to any concrete user-defined type that we own.

Let's break down what that means:

* User-defined type means that we can't add methods to int or float64 directly. We need to name our own type.
* Concrete just means we can't define methods on an interface type like Stringer.
* "That we own" means that the type and its methods must live in the same package. For example, we can't monkey patch a type defined in the standard library.

In this example we define a type for a slice of structures (A slice is a variable length array in Go). Then we define Warmest and Coldest methods on that type. And sure enough, it works.

