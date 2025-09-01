## Go Maps

Maps are used to store data values in key:value pairs. Each element in a map is a key:value pair. A map is an unordered and changeable collection that does not allow duplicates.

The length of a map is the number of its elements. You can find it using the `len()` function. The default value of a map is nil. 

Maps hold references to an underlying hash table. Go has multiple ways for creating maps.

## Create Maps Using `var` and `:=`

```go
var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
b := map[KeyType]ValueType{key1:value1, key2:value2,...}
```

This example shows how to create maps in Go. Notice the order in the code and in the output:

```go
package main  
import ("fmt")  
  
func main() {  
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}  
  b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}  
  
  fmt.Printf("a\t%v\n", a)  
  fmt.Printf("b\t%v\n", b)  
}
```

## Create Maps Using the `make()` Function:

```go
var _a_ = make(map[KeyType]ValueType)  
_b_ := make(map[KeyType]ValueType)
```

This example shows how to create maps in Go using the `make()`function:

```go
package main  
import ("fmt")  
  
func main() {  
  var a = make(map[string]string) _// The map is empty now_  
  a["brand"] = "Ford"  
  a["model"] = "Mustang"  
  a["year"] = "1964"  
                                 _// a is no longer empty_  
  b := make(map[string]int)  
  b["Oslo"] = 1  
  b["Bergen"] = 2  
  b["Trondheim"] = 3  
  b["Stavanger"] = 4  
  
  fmt.Printf("a\t%v\n", a)  
  fmt.Printf("b\t%v\n", b)  
}
```

## Create an Empty Map

There are two ways to create an empty map. One is by using the `make()`function and the other is by using the following syntax.

```go
var a map[KeyType]ValueType
```

**Note:** The `make()`function is the right way to create an empty map. If you make an empty map in a different way and write to it, it will causes a runtime panic.
This example shows the difference between declaring an empty map using with the `make()`function and without it.

```go
package main  
import ("fmt")  
  
func main() {  
  var a = make(map[string]string)  
  var b map[string]string  
  
  fmt.Println(a == nil)  
  fmt.Println(b == nil)  
}
```

## Remove Element from Map

Removing elements is done using the `delete()` function:

```go
delete(map_name, key)
```

```go
package main  
import ("fmt")  
  
func main() {  
  var a = make(map[string]string)  
  a["brand"] = "Ford"  
  a["model"] = "Mustang"  
  a["year"] = "1964"  
  
  fmt.Println(a)  
  
  delete(a,"year")  
  
  fmt.Println(a)  
}
```

## Check For Specific Elements in a Map

You can check if a certain key exists in a map using:

```go
_val_, _ok_ :=_map_name_[key]
```

If you only want to check the existence of a certain key, you can use the blank identifier (`_`) in place of val: 

```go
package main  
import ("fmt")  
  
func main() {  
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day":""}  
  
  val1, ok1 := a["brand"] _// Checking for existing key and its value_  
  val2, ok2 := a["color"] _// Checking for non-existing key and its value_  
  val3, ok3 := a["day"]   _// Checking for existing key and its value_  
  _, ok4 := a["model"]    _// Only checking for existing key and not its value_  
  
  fmt.Println(val1, ok1)  
  fmt.Println(val2, ok2)  
  fmt.Println(val3, ok3)  
  fmt.Println(ok4)  
}
```

## Iterate Over Maps

You can use `range` to iterate over maps:

```go
package main  
import ("fmt")  
  
func main() {  
  a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}  
  
  for k, v := range a {  
    fmt.Printf("%v : %v, ", k, v)  
  }  
}
```

## Iterate Over Maps in a Specific Order

Maps are unordered data structures. If you need to iterate over a map in a specific order, you must have a separate data structure that specifies that order: 

```go
package main  
import ("fmt")  
  
func main() {  
  a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}  
  
  var b []string             _// defining the order_  
  b = append(b, "one", "two", "three", "four")  
  
  for k, v := range a {        _// loop with no order_  
    fmt.Printf("%v : %v, ", k, v)  
  }  
  
  fmt.Println()  
  
  for _, element := range b {  _// loop with the defined order_  
    fmt.Printf("%v : %v, ", element, a[element])  
  }  
}
```