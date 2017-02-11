#!/usr/bin/env bash
go build && time ./latinsquares $1 results/$1.txt
