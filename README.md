# Protocol Buffers

### What is this project?
Implementing high performance API in **Golang** using _gRPC_ and _Protocol Buffers_.<br />
In this project, I created a customer storage service in which you can add users and list them.
My base idea was to create a simple API that works with **protobuf**.

This project consists of a server and a client. You can set up the server and then run
as much as clients you want to communicate with the server.

#### Directories
- client
  - the system client 
- internal
  - main packages that I build for this project
- pkg
  - external packages that I got from other places
- proto
  - my protocol buffer codes
- server
  - API main server

#### Requirements
- go 1.17
- google protobuf 1.17.1
- google gRPC 1.43
- faker 7.3

### How to use?
After cloning the project, go into the root directory:
```shell
git clone https://github.com/amirhnajafiz/Protocol-Buffers.git
```

```shell
cd Protocol-Buffers
```

Now run the server:
```shell
make serve
```

Result should be:
```shell
Attemp to listen on: 8080
```

After than you can create clients:
```shell
make use
```

And you can communicate with terminal UI:
```shell
> create (create a new customer)
....
> list   (returns a list of users)
....
> exit   (terminate the client)
```

### How does it work?
### What is proto?
### What is gRPC?