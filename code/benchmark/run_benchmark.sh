#!/bin/bash

echo "Go"
time go run code/benchmark/go_bench.go -i 100000000

echo ""
echo "Python"
time python code/benchmark/python_bench.py 10000000