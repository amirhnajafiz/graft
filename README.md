<p align="center">
    <img src="assets/go-proto.png" alt="logo" />
</p>

<h1 align="center">
Protocol Buffers
</h1>

Implementing high performance API in **Golang** using _gRPC_ and _Protocol Buffers_.
In this project, I created a customer storage service in which you can add users and list them.
My base idea was to create a simple API that works with **protobuf**.

## What do you learn in this project?
- Protocol buffers
- gRPC server in Golang

## What is Protocol Buffer?
Protocol Buffers is Google's language-neutral, platform-neutral, extensible
mechanism for serializing structured data.

Protobuf is smaller, faster and simpler that provides high performance
than other standards such as XML and JSON.

By using protocol buffers, you can define your structured data,
then you generate source code for your choice of programming language using the
protocol buffer compiler named **protoc**.

For this project, our structured data is:
```protobuf
// The customer service definition
service Customer {
  rpc GetCustomers(CustomerFilter) returns (stream CustomerRequest) {}
  rpc CreateCustomer(CustomerRequest) returns (CustomerResponse) {}
}
```

Which are two handlers. You can access the whole file in _proto_ directory.

If you want to compile the proto file, you can use the following command: (make sure to install protoc)
```shell
protoc -I customer/ customer/customer.proto --go_out=plugins=grpc:customer
```

This will run the proto compiler and generates your Golang codes.

Read more about **Protocol Buffers** at [Google](https://developers.google.com/protocol-buffers).

## What is gRPC?
gRPC is a high performance, open-source remote procedure call framework
that can run anywhere. It enables client and server applications to communicate transparently.

The gPRC framework is developed by Google. It follows HTTP semantics
over HTTP/2. It also allows you to build services with both
synchronous and asynchronous communication model. It supports traditional Request/Response
model and bidirectional streams.

Read more about **gRPC** at [gRPC website](https://grpc.io/).

### How to use this project?
After cloning the project, go into the root directory:
```shell
git clone https://github.com/amirhnajafiz/Protocol-Buffers.git
````
```shell
cd Protocol-Buffers
```

Now run the server:
```shell
go run internal/server/main.go
```

Result should be:
```shell
Attemp to listen on: 8080
```

After than you can create clients:
```shell
go run internal/client/main.go
```

And you can communicate with terminal UI:
```shell
> create (create a new customer)
....
> list   (returns a list of users)
....
> exit   (terminate the client)
```
