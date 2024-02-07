# Grpc Golang Course

This is a companion repository for my [GRPC Golang course](http://bit.ly/grpc-golang-github)

[![course logo](https://img-c.udemycdn.com/course/480x270/1792960_19b1_6.jpg)](http://bit.ly/grpc-golang-github)

# Content

- Greeting Service
- Calculator Service
- Unary, Server Streaming, Client Streaming, BiDi Streaming
- Error Handling, Deadlines, SSL Encryption
- Blog API CRUD w/ MongoDB

# How to run:

1. Install Go and ProtoBuff
2. Then run:
  go mod init PACKAGE_NAME
3. go mod tidy
4. protoc -Igreet/proto --go_out=. --go_opt=module=PACKAGE_NAME --go-grpc_out=.
   --go-grpc_opt=module=PACKAGE_NAME greet/proto/dummy.proto
5. Alternative to 4, use Makefile:
    make greet
6. Makefile link: https://github.com/Clement-Jean/grpc-go-course/blob/master/Makefile
7. If you have any problem with import run: go get -u google.golang.org/grpc