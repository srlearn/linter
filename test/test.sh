#!/usr/bin/env bash

# Test that the "positive" and "negative" examples each succeed.
# Parsing should succeed on all the positive cases and none of the
# negative cases.

# Rebuild the binary
make clean
make

bash test/test_pos.sh
bash test/test_neg.sh
