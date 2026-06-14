#!/bin/bash
# Usage: ./runner.sh <code_file> <input_file>

CODE_FILE=$1
INPUT_FILE=$2

if [ ! -f "$CODE_FILE" ]; then
  echo "Error: Code file not found." >&2
  exit 1
fi

if [ ! -f "$INPUT_FILE" ]; then
  echo "Error: Input file not found." >&2
  exit 1
fi

# Run lua script
lua5.4 "$CODE_FILE" < "$INPUT_FILE"
