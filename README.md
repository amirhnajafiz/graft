# GRAFT

Implement a distributed transaction manager using Golang and __RAFT__ consensus protocol. The idea behind this project is to replicate all transactions between multiple nodes using RAFT consensus protocol to provide data safety and consistant. Also, to provide a fault tolerance system.

## Transactions model

Each transaction has a `Sender`, `Receiver`, and `Amount`. Sender and receiver values are a 16 bits character that has only numbers, like `"1234-5678-8765-4321"`. The amount is a positive number and for each sender 
t should be greater than 10 or equal to it (a client's balance should be at least 10 dollors). 

An example payload of the transaction model:

```json
{
  "sender": "9090-8912-1209-0098",
  "receiver": "2209-6239-0911-8234",
  "amount": 50,
}
```

## Nodes

By default, a node is a follower that keeps following a leader. The system elects a leader in each term. The leader election algorithm works in a way that each node votes for candidates that have higher term, and higher log index compared to them. A log index is a sequential number that increases by each input request. Therefore, each log entry has a term and an index.

An example of log model:

```json
{
  "term": 1,
  "index": 10,
  "payload": {
    "sender": "9090-8912-1209-0098",
  },
}
```

### RPCs

There are only two RPCs for nodes to exchange data, and perform a leader election; `RequestVote` and `AppendEntry`. Servers also have other RPCs for submitting a new transaction and get a list of previous transactions, however they will be used by clients not other nodes.

#### RequestVote

Each node has a timer, and it resets that timer once it gets one request from the leader node. If that timer expires, the node increases the term and calls `RequestVote` RPC to collect others' votes. In each request, the leaders sends its last term and last index. The follower nodes compare these numbers to their own, and call a `Vote` RPC on the candidate node. This leader election process also has a timeout. If it gets expired, the term number increases without anything happenning.

#### AppendEntry

The leader node calls this RPC to replicate its logs onto other nodes. The leader starts at its newest log, and calls `AppendEntry` on each node. The follower nodes, reset their timer on each RPC call, and check the logs index and term. The follower calls a `Sync` RPC on the leader, if its log does not match to the leader. Otherwise, it will append that entry but does not commit it. The leader needs to send a log entry again in order to commit it on all nodes. Once the leader's log is committed by the majority, it will send the response to clients.

### Storage

Each node stores the committed transactions into a `MongoDB` database. It also stores the value of clients' balances into a `Redis` cache.
