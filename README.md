# README

## Database

### run
```bash
docker-compose up -d
```

### Credentials
```
jdbc:postgresql://localhost:5432/postgres 
username: postgres
password: admin
```
## Protocol Buffers

### Install

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Generate

#### Student

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/studentpb/*.proto
```

#### Test

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/testpb/*.proto
```

#### Question

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/questionpb/*.proto
```

#### Enrollment

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/enrollmentpb/*.proto
```
