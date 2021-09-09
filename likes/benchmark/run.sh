#!/bin/sh

go run main.go -ip $(echo $BENCH_IP) -conn $(echo $BENCH_MAX_CON)