Protocol buffers over JSON
- gRPC uses Protocol Buffers for communitcation
- payload size is smaller, JSON is CPU intensive because it's human readable
-


gRPC
- generates codes (C++ JAVA, Python, Go, Ruby.. etc.)
- GRPC-GO Pure implementation of gRPC in Go
- scalable (asyncronous by default, load balancing, )

Protocol Buffer defines an object in the order of (datatype, attributes, data)


HTTP 1.1
- opens a new TCP connection to a server at each request
- does not compress headers (which are plaintext)
- only works with Request/Response mechanism(no server push)
- text

TCP

HTTP/2
- released in 2015
- supports multiplexing, can push messages in parrallel over the same TCP connection
- reduces latancy greatly
- supports server push, server can push streams(multiple messages) for one request from the client
- header compression
- binary, not text (big difference)
- by default SSL required (secure!)

 4 types of API in gRPC
 - Unary (classic request and response)
 - Server Streaming (enabled throu HTTP/2, multiple response from the server! from one request)
 - Client Streaming (the opposite of Server Streaming)
 - Bi directional streaming (Server Streaming + Client, n:n request&response)
 - stream keyword!

Errors
- Plugin failed with status code 1.
=> add $GOPATH/bin to PATH then install things again



