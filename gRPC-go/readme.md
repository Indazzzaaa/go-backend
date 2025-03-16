# gRPC

gRPC (gRPC Remote Procedure Calls) is an open-source remote procedure call (RPC) framework developed by Google. It uses HTTP/2 for transport, Protocol Buffers (protobuf) as the interface description language, and provides features such as authentication, load balancing, and more.

Key features of gRPC include:

* **Language-agnostic** : Supports multiple programming languages.
* **Efficient** : Uses HTTP/2 for multiplexing and binary framing, which makes it more efficient than traditional HTTP/1.x.
* **Strongly typed** : Uses Protocol Buffers for defining service methods and message types, ensuring type safety.
* **Bi-directional streaming** : Supports client-side, server-side, and bi-directional streaming.
* **Pluggable** : Supports pluggable authentication, load balancing, and more.

gRPC is commonly used for microservices architectures, where efficient communication between services is crucial.

### Key aspects of binary framing in HTTP/2 (and by extension, gRPC) include:

1. **Efficiency** : Binary framing allows for more efficient parsing and processing of data compared to text-based protocols like HTTP/1.x. This is because binary data is more compact and can be processed faster by computers.
2. **Multiplexing** : HTTP/2 uses binary framing to enable multiplexing, which allows multiple streams of data to be sent over a single TCP connection simultaneously. Each stream is divided into frames, and frames from different streams can be interleaved and reassembled at the receiving end.
3. **Header Compression** : HTTP/2 uses HPACK, a header compression algorithm, to reduce the overhead of HTTP headers. This is particularly useful for gRPC, where headers can be repetitive and large.
4. **Flow Control** : Binary framing in HTTP/2 includes flow control mechanisms to manage the rate of data transmission and prevent congestion.

### Dependencies

1. Proto compiler from : https://github.com/protocolbuffers/protobuf/releases
2. Add it's bin path to system variable.
3. Install dependencies

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

4. Use this command to generate the code
   protoc --proto_path=e:/WorkSpace/golang/go-backend/gRPC-go/proto --go_out=. --go-grpc_out=. e:/WorkSpace/golang/go-backend/gRPC-go/proto/helloworld.proto
