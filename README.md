# GRAFT

Implement a simple distributed transaction manager using `RAFT` protocol. Nodes are implemented using Golang and for communicating, they use gRPC.
The idea is to replicate transactions between multiple nodes, using `RAFT`.

## Transactions

Each transaction has a `Sender`, `Receiver`, and `Amount`. These transactions should be replicated within all system nodes to provide a safe and consistant
system. Transactions will be executed on each machine to update the client's balances.

## Clients

Each client has an initalized balance which should be same on all system nodes. Only the leader node accepts transactions and replicates them among others.
Therefore, replicated nodes should redirect the client's requests to the leader node.

## Nodes

Each node stores transactions into a `MongoDB` database. It also stores the value of clients' balance into a `Redis` cache.
