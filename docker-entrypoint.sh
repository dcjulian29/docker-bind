#!/bin/sh

set -e

if [ $# -gt 1 ] && [ x"$1" = x"/bin/sh" ] && [ x"$2" = x"-c" ]; then
  shift 2
  eval "set -- $1"
fi

DIR=/docker-entrypoint.d

if [[ -d "$DIR" ]]; then
  /bin/run-parts "\$DIR"
fi

if [[ ! -z "$1" ]]; then
  exec "$@"
else
  /bin/sh
fi
