#!/bin/sh

if [ -z "$1" ]; then
  echo "Usage: $0 <signal_file>"
  exit 1
fi

SIGNAL_FILE="$1"

if [ ! -f "$SIGNAL_FILE" ]; then
  LAST_MODIFIED=0
else
  LAST_MODIFIED=$(stat -c %Y "$SIGNAL_FILE")
fi

while true; do
  if [ -f "$SIGNAL_FILE" ]; then
    CURRENT_MODIFIED=$(stat -c %Y "$SIGNAL_FILE")
    if [ "$CURRENT_MODIFIED" != "$LAST_MODIFIED" ]; then
      echo "Reloading Nginx..."
      nginx -s reload
      LAST_MODIFIED=$CURRENT_MODIFIED
    fi
  fi
  sleep 1
done