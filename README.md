# Twirp Swift Plugin

A protoc plugin for generating a twirp **client** suitable for iOS and macOS projects.

## Setup
The protobuf v3 compiler is required. You can get the latest precompiled binary for your system here:
https://github.com/google/protobuf/releases  \
It will also need the swift protobuf plugin from Apple (https://github.com/apple/swift-protobuf) 
 
## Usage

    # install 
    go get -u github.com/callesvedin/twirp-swift
        
    # generate
    protoc  --swift_out=:. --twirp-swift_out=:. ./service.proto
    
All generated files will be placed relative to the specified output directory for the plugin.  
Note that there can not same files in module, the generate files will named with namespace.
    