#!/bin/sh

export WAIT_HOSTS=localhost:6379
export WAIT_HOSTS_TIMEOUT=300
export WAIT_SLEEP_INTERVAL=30
export WAIT_HOST_CONNECT_TIMEOUT=30

wait && echo 'SET "cache-msg" "Hello World"' | redis-cli