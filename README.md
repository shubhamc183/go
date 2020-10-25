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