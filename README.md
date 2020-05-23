# Mac Address Resolver

This program will start a web server and listen on 8080 for HTTP request. 
And if the client is connecting on a local network, 
it will returns the Mac Address of the client.

## Prerequisite

- `arp`
- go >= 1.12

*This program has only been tested on MacOS*

## Run

```bash

go install ./...

go run cmd/main.go

```