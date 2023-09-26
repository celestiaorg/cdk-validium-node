#!/usr/bin/env bash

set -e

function wait_up {
  echo -n "Waiting for $1 to come up..."
  i=0
  until curl -s -f -o /dev/null "$1"
  do
    echo -n .
    sleep 0.25

    ((i=i+1))
    if [ "$i" -eq 300 ]; then
      echo " Timeout!" >&2
      exit 1
    fi
  done
  echo "Done!"
}

# Bring up everything else.
(
  echo "Bringing up local celestia devnet..."
  docker-compose -f docker-compose.yml up -d da
  wait_up http://localhost:26659/header/1
  export CELESTIA_NODE_AUTH_TOKEN="$(docker exec test-da-1 celestia bridge auth admin --node.store /bridge)"

  echo "Building CDK Celestium Devnet..."
  make run
)

echo "Devnet ready."