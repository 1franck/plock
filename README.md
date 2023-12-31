# Plock!

Spin off an HTTP server to serve static files from the current directory or else.

## Install

> $ go install github.com/1franck/plock@latest

## Usage

 Default command:
 > $ plock 

Will serve files from current directory to localhost:8080

To change the address and/or directory:

 > $ plock --addr localhost:3333 --dir ./foo/bar
 
## Args

| arg | description                | default        |
| --- |----------------------------|----------------|
| addr | address to serve to.       | localhost:8080 |
| dir | directory to serve from.   | ./ |
| ssl | enable ssl. (need openssl) | false |