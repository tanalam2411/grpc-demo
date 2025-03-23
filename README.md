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