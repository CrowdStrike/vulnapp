#!/bin/sh

set -e -o pipefail

cd /home/eval

# Mark as CS testcontainer
sh -c echo CS_testcontainer starting

exec /shell2http -show-errors -include-stderr -commands-file /detections.json
