# Websocket

- Client server architecture.

### **Message Flow**

1. **Client initiates a WebSocket connection** → Server registers the client.
2. **Client sends a message** → Server receives and broadcasts to all clients.
3. **Other clients receive and display the message.**
4. **If a client disconnects** → Server removes it from the active connections list.

### **Tech Stack**

* **Backend:** Go (using `github.com/gorilla/websocket` for WebSockets)
* **Frontend:** Any frontend framework (React, Vanilla JS, etc.)
* **Communication Protocol:** JSON messages over WebSockets

### **Alternative Data Formats for WebSockets**

| Format                      | Use Case                                          |
| --------------------------- | ------------------------------------------------- |
| JSON                        | General-purpose structured data                   |
| XML                         | Data exchange with legacy systems                 |
| Protocol Buffers (Protobuf) | High-performance, efficient binary format         |
| MessagePack                 | Binary format like JSON but more compact          |
| BSON                        | Binary JSON used in MongoDB                       |
| Raw Binary Data             | Streaming images, audio, video, or custom formats |

**Deps : go get -u github.com/gorilla/websocket**
