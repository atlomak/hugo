# HUGO
(HUG)e files over http written in G(O) (HUGO hehe).

Very basic tool, which I used more then once
so assumed its better to put it in a repo.

Main use case was to test some embedded platforms how they behave
when downloading arbitrary large files over http via `lwip` or some
other TCP/IP stack.

I do not plan to expand this already large tool (over 69 LOC) until I am really
bored.

# Build
Just `go build .`

## Usage
```
./hugo -p <port> -i <interface>
```
and then you can send request from your favourite http client
`<address>:<port>/<size>` where size can be any value with suffix:
 - `K`
 - `M`
 - `G`

example: `wget 127.0.0.1/100M` will download guess what... `10M` file.
