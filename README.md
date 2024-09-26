# GRAFT
In this project, we will use the RAFT protocol to build a distributed system using the Go programming language. The goal is to copy system logs between an odd number of independent nodes.

## Nodes
Each node has its own storage that keeps a series of logs in order. Since we are copying these logs, we need to ensure consistency, availability, and fault tolerance.

## RAFT
When a new request (log entry) is created, the system's leader sends it to the other nodes. Once the majority of nodes have confirmed the entry, it will be saved in each node's storage.

## I/O
As this is a distributed log storage system, user logs will be our input. Users can fetch these logs from our system nodes whenever needed.
