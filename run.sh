#!/usr/bin/env bash

go build
if type /usr/bin/time &> /dev/null; then
    (/usr/bin/time -v ./latinsquares $1 results/$1.txt) &> results/$1.time
else
    (time ./latinsquares $1 results/$1.txt) &> results/$1.time
fi

echo "Results output to results/$1.time"
