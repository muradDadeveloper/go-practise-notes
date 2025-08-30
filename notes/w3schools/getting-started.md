# Getting Started

### 1. Create a new directory and initialize a module

```go
go mod init example.com/hello
```
### 2. Create hello.go in the hello directory with the following contents

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println(reverse.String("Hello"))
}
```
### 3. Run the helloworld program: 

```go
$ go run helloworld.go 
```

### 4. If you want to save the program as an executable: 

```go
$ go build helloworld.go
```

### 5. You can run the executable using:

```go
$ ./helloworld.go
```