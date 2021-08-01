#!/usr/bin/env bash

# Some files should succeed in parsing.
RED='\033[0;31m'
LIGHTGREEN='\033[1;32m'
NC='\033[0m'

exit_status=0

for file in examples/pos/*; do
  result=$(./linter -file=${file} 2>&1)

  if [[ ! -z "${result}" ]]; then
    echo -e "Test ${file} -- ${RED}err${NC}"
    exit_status=2
  else
    echo -e "Test ${file} -- ${LIGHTGREEN}ok${NC}"
  fi

done

exit ${exit_status}
