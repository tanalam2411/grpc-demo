

## Four types of GRPC

### Unary

```proto
service YourService {
    rpc YourMethod(MessageForRequest) returns (MessageForResponse) {}
}
```

- Client sends one request and server sends one response


### Server Streaming

```proto
service YourService {
    rpc YourMethod(MessageForRequest) returns (stream MessageForResponse) {}
}
```

- Client sends one request and Server sends multiple response

### Client Streaming

```proto
service YourService {
    rpc YourMethod(stream MessageForRequest) returns (MessageForResponse) {}
}
```

- Client sends multiple requests and Server sends one response

### Bi-Directional Streaming

```proto
service YourService {
    rpc YourMethod(stream MessageForRequest) returns (stream MessageForResponse) {}
}
```

- Both Client and Server can send and receive multiple requests and response

---

gRPC

- Exactly one request and one response
- Message ordering is guaranteed for each individual call  of gRPC method
  - Request part 1, 2, 3, ...
  - Response part 1, 2, 3, ...