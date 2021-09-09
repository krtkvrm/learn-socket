#!/bin/bash
a=0

while [ $a -lt 10 ]
do
    a=`expr $a + 1`
    docker run -d --network=host -e BENCH_IP='0.0.0.0:8080' -e BENCH_MAX_CON='10' krtkvrm/websocket-hitter
done