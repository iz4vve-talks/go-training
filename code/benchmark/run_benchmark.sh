#!/bin/bash

echo "Go"
time go run code/benchmark/go_bench.go 100000000

echo "Python"
time python code/benchmark/python_bench.py 10000000