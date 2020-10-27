# Learning GO lang

Learn Go Programming - Golang Tutorial for Beginners: https://www.youtube.com/watch?v=YS4e4q9oBaU

## Setting up the env
- Download go binary
- As root, tar -C /usr/local -xzf go1.15.3.src.tar.gz
- vim ~/.basrc
  - Add go binary path
    - `export GOROOT=/usr/local/go`
    - `export PATH=$PATH:$GOROOT/bin`
  - Add GOLIB for 3rd party libraries
    - Folder will be automatically created, `mkdir /home/shubham_choudhary/golib`
    - `export GOPATH=/home/shubham_choudhary/golib`
    - `export PATH=$PATH:$GOPATH`
    - Check by running: `go get github.com/nsf/gocode` and check home/shubham_choudhary/golib
  - Adding workspace
    - Add workspace location too
    - `export GOPATH=$GOPATH:/LOCATION_OF_WORKSPACE`
  - We didn't mess our WORKSPACE folder with 3rd party libraries via this approach

### Running go code
- go run src/path/to/code/Main.go
- go build path/to/code/ to package under bin/ folder

### Go Features
- Singe binary file o/p
- Can't declare unused vars
- Very strict on operation with different Primitive types

### Primitives
- Bool
- Numberic
  - Integers
    - Signed, Unsigned
    - int, int8, int32
  - Floating
  - Complex
- Text String
  - String
    - UTF-8
    - Immutable
    - Can be concatenated with plus (+)
    - Can be converted to []byte
  - Rune
    - UTF-32
    - Alias for int32
    - Special methods normally requires to process

### Constants
- Constants
  - Immutable
  - Assigned at compile time, must be present at that time
  - Implicit type conversion works
  - Can be defined in block
- Enurated constants
  - ```iota``` starts at 0 in each const block and increment by one

### Collections
- Arrays
  - Collection of items with same types
  - Fixed Size
  - call by value & reference
- Slices
  - Backed by array, it also have same type of array
  - ```a := []int{10, 20, 30}```
  - Size & capacity based
  - Always call by refrence
  - a = append(a, 10)
  - a = append(a, 10, 20, 30)
  - a = append(a, b...) // spreador
  - **DANGEROUS, a will be changed use loops**
    - ```b = append(a[:2], a[4:]...)```
- Make funtion
  - ```make(type, size)```, cap becomes size
  - ```make(type, size, cap)```
- Maps
  - **ALWAYS Pass by refrence***
  - always return something like 0, check with ok
  - ```value, ok = myMap["wrongKey"]```
- Structs
  - Collection of different types of data
  - Keyed by name
  - Can be passed by value & reference
  - No inheritance, but embedding possible
  - Tags can be fields of struct to provide metadata
