# Steps to configure gRPC

### 1. Install dependencies

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/grpc@latest

```

### 2. Install Protoc from below link and add bin dir to system environment path

```plaintext
https://github.com/protocolbuffers/protobuf/releases
```

### 3. Define gRPC service (in .proto file)

```go
syntax = "proto3";

package greet; // This proto file belongs to greet package
option go_package = "./greetpb"; // Specifies go package where the generated code should be placed.

/*
Define gRPC sercie named Greeter.
This service has one RPC method called SayHello that:
Accepts a HelloRequest message
Returns a HelloResponse message
*/
service Greeter {
  rpc SayHello (HelloRequest) returns (HelloResponse);
}

// Defines a message (data structure) named HelloRequest contains one filed, it's explore in detials below.
message HelloRequest {
  string name = 1;
}

message HelloResponse {
  string message = 1;
}

```

### 4. Generate go code from .proto file

```
protoc --go_out=. --go-grpc_out=. greet.proto

```

This will generate:

* `greet.pb.go` (message definitions)
* `greet_grpc.pb.go` (service definitions)

### 5. Implement the gRPC server and Client. 

- Look at the code how it's implemented and how the created files have been referenced.

---

### What is this syntex explain about it

```go
message HelloRequest {
  string name = 1;
}

```


* Defines a **message** (data structure) named `HelloRequest`.
* Contains  **one field** :
  * `string name = 1;`
    * `string` → Data type.
    * `name` → Field name.
    * `1` → Field number (used for serialization).

### **Why Do We Need Field Numbers?**

Field numbers are used **internally by Protocol Buffers** for **efficient serialization** and  **backward compatibility** .

Instead of storing field names like `"name"` in serialized data, Protocol Buffers store just the field **number** (`1`). This makes data smaller and faster to process.


### **Rules for Field Numbers**

1. **Must be unique within a message**

   * Every field in a message must have a unique number.

     ```go
     message Person {
       string name = 1;
       int32 age = 2;
       string email = 3;
     }

     ```
   * Here, `name`, `age`, and `email` have unique field numbers.

* **Valid Number Ranges**

  * Allowed: `1` to `2^29 - 1` (`1` to `536,870,911`).
  * Recommended:
    * `1 - 15` → **Smaller & efficient (1 byte when serialized).**
    * `16 - 2047` → **Uses more bytes (2+ bytes).**
  * Numbers `19000 - 19999` are **reserved** and should not be used.
* **Backward Compatibility**

  * If you later  **add a new field** , old clients **won't break** because they ignore unknown fields.

    ```go
    message HelloRequest {
      string name = 1;
      int32 age = 2; // New field added later
    }

    ```

    Old clients will just ignore `age` if they don’t understand it.
* **How Field Numbers Affect Serialization**

  When `HelloRequest` is serialized, it gets converted into a  **compact binary format** , where:

  * `1` (the field number) is stored instead of `"name"`.
  * The actual string value (`"John"`, `"Alice"`, etc.) follows.

  This makes gRPC **faster and smaller** compared to formats like JSON.


### Example: JSON vs ProtoBuf Serialization

```json
{
  "name": "John"
}

```

```go
// protoBuf( Compact, Machine-Friendly)
08 04 4A 6F 68 6E

```

* `08` → Field number `1`, data type `string`.
* `04` → Length of string (`4` bytes for `"John"`).
* `4A 6F 68 6E` → `"John"` in bytes.

This is much smaller and faster to process!
