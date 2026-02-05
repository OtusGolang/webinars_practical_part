#!/bin/bash
# Quick start script for the caching demo

echo "Building and running caching eviction mechanisms demo..."
echo ""

cd "$(dirname "$0")"

# Run the demo
go run .

echo ""
echo "Demo completed! See README.md for detailed analysis."
