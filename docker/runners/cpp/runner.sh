#!/bin/sh

# The code will be mounted at /sandbox/Solution.cpp
# The input will be piped in via stdin or mounted as input.txt

if ! g++ -O2 -Wall -std=c++17 Solution.cpp -o solution; then
    echo "Compilation Error"
    exit 1
fi

# Run the compiled binary
./solution
