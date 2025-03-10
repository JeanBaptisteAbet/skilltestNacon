# skilltestNacon

Technical test conducted for a job interview with these specifications :

-  **Exposes an HTTP API** to **retrieve live events for players**.
-  **Exposes a gRPC API** to **create, update, and delete live events** (internal usage).
-  **Runs both HTTP and gRPC servers on the same port**.
-  **Uses API keys** to distinguish between **public users (HTTP access)** and **internal users (gRPC access)**.
-  **Persists data using SQLite**.
-  **Uses Protobuf definitions** for gRPC.
-  **Includes unit tests** where relevant.

Done in 12h30 hours :

- Installing environment and writing database package : 1h30
- Writing http api: 0h45
- Installing environment and writing grpc protobuf, server and client : 2h15
- Grpc and http on same port : 4h (had to try multiple idea)
- Adding some test, didn't finish : 2h30
- Code cleaning : 0h45
- Writing dockerfile : 0h30
- Writing README : 0h15

 
## How to install the project

### Requirement 
  - go
  - docker
  - docker-compose

 ### Step
+ Create a .env file from the .env.template 
+ Complete the .env file
+ Build and launch the project with docker-compose
	+  `docker-compose build && docker-compose up -d`
 
 ### Test the project 
 + http server:
	 + `curl localhost:$PORT/events --header 'Authorization: $HTTP_API_KEY'` 
	 +  `curl localhost:$PORT/events/$ID --header 'Authorization: $HTTP_API_KEY'` 
 + grpc server: 
	 + use `api/grpcserver/client/main.go` with `go run ./api/grpcserver/client/main.go`
	 
### **API Definitions**

#### **HTTP Endpoints (Public)**

| Method | Endpoint       | Description                  | Authentication               |
|--------|----------------|------------------------------|------------------------------|
| `GET`  | `/events`      | Retrieve all active events   | Requires API Key (HTTP role) |
| `GET`  | `/events/{id}` | Retrieve details of an event | Requires API Key (HTTP role) |

#### **gRPC Endpoints (Internal)**

| Method        | RPC Name      | Description                             | Authentication               |
|---------------|---------------|-----------------------------------------|------------------------------|
| `CreateEvent` | `CreateEvent` | Create a new live event                 | Requires API Key (gRPC role) |
| `UpdateEvent` | `UpdateEvent` | Update an existing event                | Requires API Key (gRPC role) |
| `DeleteEvent` | `DeleteEvent` | Delete an event                         | Requires API Key (gRPC role) |
| `ListEvents`  | `ListEvents`  | Retrieve all events (incl. past events) | Requires API Key (gRPC role) |

## Area of ​​improvement
+ Missing test, and better writing of test.
+ Err wrapping 
+ Improve log with more details
+ Better usage of sqlite (type)
+ Better authentification system (jwt)


