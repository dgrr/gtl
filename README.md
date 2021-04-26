# GTL: Golang Template Library

GTL is a template library written in pure Go(2).
It is intended to hold common data structures that might be missing in the standard library
or they are cumbersome if generics are not present (using interface{} as replacement, type casting, etc...).

You can learn more about Golang's generics [here](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md).

## Data Structures

### Result

Result tries to emulate [Rust's result](https://doc.rust-lang.org/std/result/).

After getting the returned value,
we can manage the [Result](https://github.com/dgrr/gtl/blob/master/result.go2#L4) in different ways.

A valid usage would be:
```go
func Do() (r gtl.Result[string, error]) {
  if rand.Intn(100) & 1 == 0 { // is even?
    return r.Err(
      errors.New("unexpected result"))
  }
  
  return r.Ok("Success")
}

func main() {
  Do().Then(func(v string) {
    fmt.Printf("Got result: %s\n", v)
  }).Else(func(err error) {
    fmt.Printf("Got error: %s\n", err)
  })
}
```

`Then` is executed if `Do` returned the call to `Ok`. `Ok` will store the string `"Success"`
into the Result's expected value. In the other hand, `Else` will be executed if `Err` has been called.

If we don't want to handle errors but we just want to get the value, we can do the following:
```go
func Do() (r gtl.Result[string, error]) {
  // assume Do() didn't change
}

func main() {
  fmt.Printf("Got %s\n", Do().Or("Failed"))
}
```


### Vector

Vector tries to emulate [C++'s vector](https://en.cppreference.com/w/cpp/container/vector) (somehow).
It doesn't try to emulate it exactly, but it just works as a C++ vector in a way that internally is just
a slice with some helper functions, in this case functions like `PushBack`, `PushFront`, `PopBack`, `PopFront` or `Len`.

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
        var vec gtl.Vector[string]

        vec.PushBack(os.Args[1:]...)

        sort.Slice(vec, func(i, j int) bool {
                return vec[i] < vec[j]
        })

        fmt.Println(vec)
}
```
