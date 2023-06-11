# artillery_game
A game where you try to detect and destroy the opposing artillery.  Destroy them before they destroy you!


# Honoring the Golang Rituals

When this repo was created, the venerated modules commands:

    $ go mod init github.com/joshuamhtsang/artillery_game

    $ go mod tidy

were used to create the 'go.mod' file.  It's no longer necessary to put develop
Go code under $GOPATH.  For more information see:
https://medium.com/@vingarcia00/golang-why-not-use-the-gopath-87521259663a


### Running the code

# Setting up protobuf binaries
This section follows the steps here:
https://protobuf.dev/getting-started/gotutorial/

Download protobuf binary from:
https://github.com/protocolbuffers/protobuf/releases

Place the downloaded zip file into the directory "protobuf_binary" and unzip it:

    $ unzip protoc-23.2-linux-x86_64.zip 

'go install' the Go protocol buffers plugin.

    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

By defult, it will install the binaries into GOPATH:

    $HOME/go/bin/

which you can check but running:

    $ go env
    ...
    GOPATH="/home/joshua/go"
    ...



