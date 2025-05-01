# grpc-demo


## Server Streaming gRPC

- Client send one request
- Server send stream of responses
- Client can process each response as it arrives
- Like watching video streaming


## Client Streaming gRPC

- Client send many requests
- Server receive and process each request
- Client finish sending request, then server send one response back
- Response business logic based on all received requests

### Sample Use Cases

- Multiple image processing
- Real time analytics
- Sensor data collection

## Bi-Directional Streaming gRPC

- Client send many requests
- Server receive & process each request
- Server send many responses

### Sample Use Cases

- Real time analytics, take action based on certain data
- Real time collaboration
- Data collection from multiple IoT devices



---

```shell
wget https://raw.githubusercontent.com/googleapis/googleapis/refs/heads/master/google/type/date.proto

```

Install Postgresql - https://squaredup.com/blog/running-postgres-in-docker/
Migration tool - https://github.com/golang-migrate/migrate

---

### Bank gRPC

- Functionalities:
  - Get current balance (unary)
  - Get exchange rates (server stream)
  - Summarize transactions (client stream)
  - Transfer to multiple accounts (bi-directional stream)

---

## Status

- REST API has HTTP response status codes: 1xx(Informational),
  2xx(Success),3xx(Redirect),4xx(Client error), 5xx(Server error)

- gRPC status codes
  - Success
  - Error
  - Stream termination

- gRPC status message

- Status:
  - Successful call should return status code 0
  - Build status using go packages
    - google.golang.org/grpc/status
    - google.golang.org/grpc/codes

- https://grpc.github.io/grpc/core/md_doc_statuscodes.html



| HTTP Status Code       | gRPC Status Code          | Meaning                                                                                   |
|------------------------|---------------------------|-------------------------------------------------------------------------------------------|
| 200 OK                 | 0 OK                      | Success, no error.                                                                        |
| 400 Client error       | 3 INVALID_ARGUMENT        | Client specified as invalid argument. Check error message and error details for more information. |
| 400 Client error       | 9 FAILED_PRECONDITION     | Request can not be executed in the current system state, such as deleting a non-empty directory. |
| 400 Client error       | 11 OUT_OF_RANGE           | Client specified an invalid range.                                                        |
| 401 Unauthorized       | 16 UNAUTHENTICATED        | Request not authenticated due to missing, invalid, or expired OAuth token.                 |
| 403 Forbidden          | 7 PERMISSION_DENIED       | Client does not have sufficient permission. This can happen because the OAuth token does not have the right scopes, the client doesn't have permission, or the API has not been enabled. |
| 404 Not Found          | 5 NOT_FOUND               | A specified resource is not found.                                                        |
| 409 Conflict           | 10 ABORTED                | Concurrency conflict, such as read-modify-write conflict.                                  |
| 409 Conflict           | 6 ALREADY_EXISTS          | The resource that a client tried to create already exists.                                 |
| 429 Too Many Requests  | 8 RESOURCE_EXHAUSTED      | Either out of resource quota or reaching rate limiting. The client should look for `google.rpc.QuotaFailure` error detail for more information. |
| 499 Client Closed Request | 1 CANCELLED            | Request cancelled by the client.                                                          |
| 500 Internal Server Error | 15 DATA_LOSS           | Unrecoverable data loss or data corruption. The client should report the error to the user. |
| 500 Internal Server Error | 2 UNKNOWN              | Unknown server error. Typically a server bug.                                             |
| 500 Internal Server Error | 13 INTERNAL            | Internally server error. Typically a server bug.                                          |
| 501 Not Implemented    | 12 UNIMPLEMENTED          | API method not implemented by the server.                                                 |
| 503 Service Unavailable| 14 UNAVAILABLE            | Service unavailable. Typically the server is down.                                        |
| 504 Gateway Timeout    | 4 DEADLINE_EXCEEDED       | Request deadline exceeded. This will happen only if the caller sets a deadline that is shorter than the method's default deadline (i.e., requested deadline is not enough for the server to process the request) and the request did not finish within the deadline. |



| HTTP Status Code | gRPC Status Code          | Meaning                                |
|------------------|---------------------------|----------------------------------------|
| 200 OK           | `OK` (0)                  | Request was successful                 |
| 400 Bad Request  | `INVALID_ARGUMENT` (3)    | Client provided invalid request data   |
| 401 Unauthorized | `UNAUTHENTICATED` (16)    | Authentication failed or missing       |
| 403 Forbidden    | `PERMISSION_DENIED` (7)   | Client does not have access rights     |
| 404 Not Found    | `NOT_FOUND` (5)           | Resource not found                     |
| 409 Conflict     | `ALREADY_EXISTS` (6)      | Resource already exists                |
| 429 Too Many Requests | `RESOURCE_EXHAUSTED` (8) | Rate limit exceeded                    |
| 500 Internal Server Error | `INTERNAL` (13)     | Server encountered an internal error   |
| 501 Not Implemented | `UNIMPLEMENTED` (12)     | Method not implemented                 |
| 503 Service Unavailable | `UNAVAILABLE` (14)    | Service is currently unavailable       |
| 504 Gateway Timeout    | `DEADLINE_EXCEEDED` (4) | Request timeout exceeded               |


---

### Retry Pattern

- https://github.com/grpc-ecosystem/go-grpc-middleware/tree/main/interceptors/retry

gRPC Interceptor:

- Intercepts gRPC traffic & add functionality
- Interceptor configuration: WithCodes, WithMax, WithBackoff
- Configuration has default values(can be overridden)

`WithCodes`: when to retry(based on received response code)
`WithMax`: Maximum limit to retry
  - keep retry until reach maximum limit, or until get response code not included on WithCodes
`WithBackOff`:
  - Requires `grpc_retry` function
  - `BackoffLinear`: fixed time interval
  - `BackoffExponential`: current interval * 2
  - `BackoffLinearWithJitter / BackoffExponentialWithJitter`

```go
import grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/interceptors/retry"
...
...

var opts []grpc.DialOption

opts = append(opts,
      grpc.WithUnaryInterceptor(
        grpc_retry.UnaryClientInterceptor(   # UnaryInterceptor on dial options
          grpc_retry.WithCodes(...), grpc_retry.WithMax(...),
          grpc_retry.WithBackoff(grpc_retry.BackoffExponential(...)),
        )
      )
    )


opts = append(opts,
      grpc.WithStreamInterceptor(
        grpc_retry.StreamClientInterceptor(   // StreamInterceptor on dial options (Note: only works for server streaming gRPC API)
          grpc_retry.WithCodes(...), grpc_retry.WithMax(...),
          grpc_retry.WithBackoff(grpc_retry.BackoffExponential(...)),
        )
      )
    )

conn, err := grpc.Dial("localhost:9090", opts...)
```

```shell
go get github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry
go: downloading github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.3.1
go: added github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.3.1
```

---

### Circuit Breaker Pattern

- Circuit: connection between two services
- Circuit open when error rates reach threshold, intentionally disabling connection
- Request immediately fail when circuit opens
- Circuit reset(half-open) after configured time
- No error: close circuit & resume traffic
- Repeat process (open/close circuit) based on error

Terms:

- MaxRequests ````
- Interval
- Timeout
- ReadyToTrip
- OnStateChange

- Can define several circuit breakers with different configurations
- Use: https://github.com/sony/gobreaker
- Ony works for unary gRPC API call

---

### gRPC Metadata

- Similar to REST API: HTTP Header (request / response)
- Adding/reading additional information
- gRPC metadata (request / response)
  - key-value pairs

Use Cases

- Authz/n
- Routing
- Tracing & monitoring
- Rate limiting

```go
import (
  "google.golang.org/grpc/metadata"
)

// read request metadata
if requestMetadata, ok := metadata.FromIncomingContext(ctx); ok{
  // ...
}

//  add response metadata
md := map[string]string{
  "response-metadata-key-1": "response-metadata-value-1",
  ...
}

responseMetadata := metadata.New(md)    // can only called once, multiple calls will raise error
err := grpc.SendHeader(ctx, responseMetadata)
```

#### gRPC Metadata

- Client to send request metadata or read response metadata
- Streaming response metadata will has one set of metadata for each opened stream

---

## Interceptor

If we need to a functionality that must be added to all grpc request:

On Client side: 
- Log outgoing client call
- Adding authentication token to be validated by server
- Modify the request message
- Adding some default value if user does not provide the value at the original request message
- Add metadata

On ServerSide:
-  Log incoming request
- Validate authentication token
- Reject the request if authentication token if not valid
- Modify response message, (e.g., masking credit card number)

Interceptor:
- Components that can be added to gRPC client or server
- Intercept & process messages
- Enable common functionality implementation without cluttering main business logic

Interceptor Types:
- UnaryServerInterceptor
- StreamServerInterceptor
- UnaryClientInterceptor
- StreamClientInterceptor

---

### Server Interceptor

Server - Unary interceptor, basic

```go
func MyUnaryServerInterceptor() grpc.UnaryServerInterceptor {
  return func(ctx contexnt.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
    // ... interceptor logic

    return handler(ctx, req)
  }
}
```

Server - Stream interceptor, basic

```go
func MyStreamServerInterceptor() grpc.StreamServerInterceptor {
  return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
    // ... interceptor logic

    return handler(srv, stream)
  }
}
```

Server - Unary interceptor, modify response message

```go
func MyUnaryServerInterceptor() grpc.UnaryServerInterceptor {
  return func(ctx contexnt.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
    
    res, _ := handler(ctx, req)

    if response, ok := res.(*MyResponseMessage); ok {
      // ... do something with response message
    }

    return handler(ctx, req)
  }
}
```

Server - Stream interceptor, modify response message

```go

type InterceptedServerStream struct {
  grpc.ServerStream
}

func MyStreamServerInterceptor() grpc.StreamServerInterceptor {
  return func(srv interface{}, serverStream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
    interceptedServerStream := &InterceptedServerStream{
      ServerStream: serverStream
    }
    return handler(srv, interceptedServerStream)
  }
}

func (s *InterceptedServerStream) SendMsg(msg interface{}) error {
  switch response := msg.(type){
    case *MyResponseMessage:
      // ...modify response
      return s.ServerStream.SendMsg(response)
    default:
      // Forward the original message to the original stream
      return s.ServerStream.SendMsg(msg)
  }
 }
```


Server - Unary interceptor, modify request message

```go

func MyUnaryServerInterceptor() grpc.UnaryServerInterceptor {
    return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
      // modify request
      switch request := req.(type) {
      case *hello_proto.HelloRequest:
        // modify request
        request.Name = "[MODIFIED BY SERVER INTERCEPTOR - 1]" + request.Name
    }
    return handler(ctx, req)
  }
}
```


Server - Stream interceptor, modify request message

```go
func (s *InterceptedServerStream) RecvMsg(msg interface{}) error {
  err := s.ServerStream.RecvMsg(msg)
  if err !=  nil{
    return err
  }

  switch request := msg.(type) {
    case *MyRequestMessage:
      // ... modify request
  }
  return nil
}
```

Server - Unary interceptor, modify response metadata

```go
func MyUnaryServerInterceptor() grpc.UnaryServerInterceptor {
  return func(ctx contexnt.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

    responseMetadata, ok := metadata.FromOutgoingContext(ctx)
    if !ok{
      responseMetadata = metadata.New(nil)
    }

    responseMedata.Append("keyToAdd", "valueToAdd")
    responseMedata.Set("keyToUpdate", "newValue")
    responseMedata.Delete("keyToDelete")

    ctx = metadata.NewOutgoingContext(ctx, responseMetadata)
    grpc.SetHeader(ctx, responseMetadata)

    return handler(ctx, req)
  }
}
```

---

Server - Unary interceptor(single)

```go
grpc.UnaryInterceptor(clientUnaryInterceptor_1)
```

Server - Unary interceptor(multiple)
```go
grpc.ChainUnaryInterceptor(
  clientUnaryInterceptor_1,
  clientUnaryInterceptor_2,
  clientUnaryInterceptor_3,
),
```

Create interceptor on server
```go
grpcServer := grpc.NewServer(
  // ...  interceptor here (single / multiple)
)
```

---

Client - Unary interceptor, basic

```go
func MyUnaryClientInterceptor() grpc.UnaryClientInterceptor{
  return func(ctx context.Context, method string, req, reply interface{}, 
              cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
      // ... interceptor logic

      return invoker(ctx, method, req, reply, cc, opts...)

    }
}
```

Client - Stream interceptor, basic

```go
func MyStreamClientInterceptor() grpc.StreamClientInterceptor {
  return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, 
              streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
    // .. interceptor logic

    return streamer(ctx, desc, cc, method, opts...)
  }
}
```

Client - Unary interceptor, modify request message
```go
func MyUnaryClientInterceptor() grpc.UnaryClientInterceptor{
  return func() error {
    switch request := req.(type) {
      case *MyRequestMessage:
        // ... modify request
    }
    return invoker(ctx, method, req, reply, cc, opts...)
  }
}
```

Client - Stream interceptor, modify request message
```go
type InterceptedClientStream struct {
  grpc.ClientStream
}

func MyStreamClientInterceptor() grpc.StreamClientInterceptor {
  return func() (grpc.ClientStream, error) {
    clientStream, err := streamer(ctx, desc, cc, method, opts...)

    interceptedClientStream := &InterceptedClientStream{
      ClientStream: clientStream,
    }
    return interceptedClientStream, err
  }
}

func (s *InterceptedClientStream) SendMsg(msg interface{}) error {
  switch request := msg.(type) {
    case *MyRequestMessage:
      // ... modify request
  }

  return s.ClientStream.SendMsg(msg)
}
```

Client - Unary interceptor, modify response message

```go
func MyUnaryClientInterceptor() grpc.UnaryClientInterceptor{
  return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker,
              opts ...grpc.CallOption) error {
    
    err := invoker(ctx, method, req, reply, cc, opts...)
    if err != nil{
      return err
    }

    swtich response := reply.(type) {
      case *MyResponseMessage:
        // ... modify response
    }

    return err
  }
}
```

Client - Stream interceptor, modify response message

```go
func (s *InterceptedClientStream) recvMsg(msg interface{}) error {
  err := s.ClientStream.RecvMsg(msg)
  if err != nil {
    return err
  }

  switch response := msg.(type) {
    case *MyResponseMessage:
      // ... modify response
  }
  return nil
}
```


Client - Unary interceptor, modify request metadata

```go
func MyUnaryClientInterceptor() grpc.UnaryClientInterceptor {
  return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker,
              opts ...grpc.CallOption) error {
    md, ok := metadata.FromOutgoingContext(ctx)
    if !ok{
      md = metadata.New(nil)
    }

    md.Append("keyToAdd", "valueToAdd")
    // Add new metadata can also use this
    metadata.AppendToOutgoingContext(ctx, "keyToAdd", )

  }
}
```