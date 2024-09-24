#!/bin/bash

export $SRC_DIR=proto/
export DST_DIR=pkg/rpc/

protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/log.proto
