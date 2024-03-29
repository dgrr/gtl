# DEPRECATION!

## Use TL (https://github.com/dgrr/tl)

# GTL: Golang Template Library (WIP)

GTL is a template library written in pure Go(2).
It is intended to hold common data structures that might be missing in the standard library
or they are cumbersome if generics are not present (using interface{} as replacement, type casting, etc...).

You can learn more about Golang's generics [here](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md).

# Table of Contents
1. [Result](#result)
2. [Optional](#optional)
3. [Iterator](#iterator)
4. [Vector](#vector)
5. [Bytes](#bytes)
6. [Numeric](#numeric)
7. [Pair](#pair)

## Result

Result tries to emulate [Rust's result](https://doc.rust-lang.org/std/result/).
But being honest, it is more similar to [std::expected](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2017/p0323r3.pdf) from C++.

As Golang already has a type for defining errors, the second type is replaced with the error.

After getting the returned value,
we can manage the [Result](https://github.com/dgrr/gtl/blob/master/result.go2#L4) in different ways.

A valid usage would be:
```go
func Do() (r gtl.Result[string]) {
  if rand.Intn(100) & 1 == 0 { // is even?
    return r.Err(
      errors.New("unexpected result"))
  }
  
  return r.Ok("Success")
}

func onSuccess(v string) {
	fmt.Printf("Got result: %s\n", v)
}

func onError(err error) {
	fmt.Printf("Got error: %s\n", err)
}

func main() {
  Do().Then(onSuccess).Else(onError)
}
```

`Then` is executed if `Do` returned the call to `Ok`. `Ok` will store the string `"Success"`
into the Result's expected value. In the other hand, `Else` will be executed if `Err` has been called.

If we don't want to handle errors but we just want to get the value, we can do the following:
```go
func Do() (r gtl.Result[string]) {
  // assume Do() didn't change
}

func main() {
  fmt.Printf("Got %s\n", Do().Or("Failed"))
}
```

## Optional

Optional represents an optional value. In C++ we have the [std::optional](https://en.cppreference.com/w/cpp/utility/optional)
which might be similar.

A valid usage would be:
```go
func myFunc() (o gtl.Optional[int]) {
	if n := rand.Int(); n % 2 == 0 { // is even
		o.Set(n)
	}
	
	return o
}

func main() {
	value := myFunc()
	if value.Has() {
		fmt.Printf("Got: %d\n", value.V())
	}
}
```

## Iterator

Iterator tries to emulate a [C++'s iterator](https://en.cppreference.com/w/cpp/iterator/iterator).
It is defined as follows:
```go
// Iterator defines an interface for iterative objects.
type Iterator[T any] interface {
	// Next increments the iterator.
	Next() bool
	// Advance advances the cursor `n` steps. Returns false if `n` overflows.
	Advance(n int) bool
	// Get returns the value held in the iterator.
	Get() T
	// Ptr returns a pointer to T.
	Ptr() *T
}
```

## Vector

Vec tries to emulate a [C++'s vector](https://en.cppreference.com/w/cpp/container/vector) (somehow).
It doesn't try to emulate it exactly, but it just works as a C++ vector in a way that internally is just
a slice with some helper functions, in this case functions like `Append`, `Push`, `PopBack`, `PopFront` or `Len`.

A valid usage would be:
```go
package main

import (
        "os"
        "sort"
        "fmt"

        "github.com/dgrr/gtl"
)

func main() {
        vec := gtl.NewVec(os.Args[1:]...)

        sort.Slice(vec, func(i, j int) bool {
                return vec[i] < vec[j]
        })

        fmt.Println(vec)
}
```

## Bytes

Bytes is a helper for working with byte slices. It can be used as
a normal slice for reading data from a net.Conn.

You can see an example of how to use Bytes [here](https://github.com/dgrr/gtl/blob/2642e2ac98bd8a8fbfbc3e9789d4b87bf6e6e317/examples/echo_tcp/main.go2#L73).

### Numeric

There are global numeric helper functions like [Max](https://github.com/dgrr/gtl/blob/b5b6ba36de904e757d00f78351c577a6ad0547e1/numeric.go2#L9)
and [Min](https://github.com/dgrr/gtl/blob/b5b6ba36de904e757d00f78351c577a6ad0547e1/numeric.go2#L19).

### Pair

Pair is a data structure that holds 2 values [T, U].
```
pair := gtl.MakePair("hello", 1234)
```
