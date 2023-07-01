#!/usr/bin/env bash
set -euo pipefail

http_server() {
  echo -n "Starting python http server..."
  python -m http.server >>http.log 2>&1 &
  RETVAL=$?
  PID=$!
  echo $RETVAL
  echo $PID
  [ $RETVAL -eq 0 ] && echo $PID >http.lock && echo success || echo fail
  echo
  return $RETVAL
}

http_server
