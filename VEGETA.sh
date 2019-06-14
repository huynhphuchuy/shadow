#!/bin/bash
echo "GET http://localhost:6969" | ./cmd/cli/stress-test/vegeta attack -duration=0 -connections=1000000