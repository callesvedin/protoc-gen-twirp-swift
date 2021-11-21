# Twirp Swift Plugin

A protoc plugin for generating a twirp client suitable for iOS and macOS projects. Build from the [Swift Twirp plugin](https://github.com/CrazyHulk/protoc-gen-swiftwirp).

## Setup

The protobuf v3 compiler is required. You can get the latest precompiled binary for your system here:

https://github.com/google/protobuf/releases

### Twirp Go Server (optional)

While not required for generating the client code, it is required to run the server component of the example.

    go get github.com/twitchtv/twirp/protoc-gen-twirp
    go get -u github.com/golang/protobuf/protoc-gen-go

    # generate go files
    protoc -I . --twirp_out=. --go_out=:. example/cmd/haberdasher/service.proto

    # start server
    cd example/cmd
    go run main.go
 
## Usage

    # install 
    go get -u github.com/callesvedin/twirp-swift
    
    protoc  --swift_out=:. --twirp-swift_out=:. ./example/cmd/haberdasher/service.proto
    
All generated files will be placed relative to the specified output directory for the plugin.  
Note that there can not same files in module, the generate files will named with namespace.

Using the Twirp hashberdasher proto:
    
```swift
var c = Haberdasher_Haberdasher()
c.hostname = "http://127.0.0.1:9000"

var req = Haberdasher_Size()
req.inches = 10
c.makeHat(size: req) { resp in
    print(resp)
}
```
    
### Parameters

The plugin parameters should be added in the same manner as other protoc plugins. 
Key/value pairs separated by a single equal sign, and multiple parameters comma separated.

## Using the Example

Run the server:

    cd example/cmd
    go run main.go
     
In a new terminal run the client:
 
    cd example/swift/service
    swift build
    swift run
