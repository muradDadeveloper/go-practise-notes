# Go by Example: Structs

Go’s _structs_ are typed collections of fields. They’re useful for grouping data together to form records. This `person` struct type has `name` and `age` fields. `newPerson` constructs a new person struct with the given name. Go is a garbage collected language; you can safely return a pointer to a local variable - it will only be cleaned up by the garbage collector when there are no active references to it. 

This syntax creates a new struct. You can name the fields when initializing a struct.
Omitted fields will be zero-valued. An `&` prefix yields a pointer to the struct. It’s idiomatic to encapsulate new struct creation in constructor functions. Access struct fields with a dot.

You can also use dots with struct pointers - the pointers are automatically dereferenced. Structs are mutable. If a struct type is only used for a single value, we don’t have to give it a name. The value can have an anonymous struct type. This technique is commonly used for table-driven tests.

 
```go
package main
import "fmt"

type person struct {
    name string
    age  int
}

func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42
    return &p
}

func main() {
    fmt.Println(person{"Bob", 20})
    fmt.Println(person{name: "Alice", age: 30})
    fmt.Println(person{name: "Fred"})
    fmt.Println(&person{name: "Ann", age: 40})
    fmt.Println(newPerson("Jon"))

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    sp := &s
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(sp.age)

    dog := struct {
        name   string
        isGood bool
    }{
        "Rex",
        true,
    }
    fmt.Println(dog)
}
```

Output:

```Bash
$ go run structs.go
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
&{Jon 42}
Sean
50
51
{Rex true}
```