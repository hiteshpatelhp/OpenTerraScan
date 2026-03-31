#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

export OPENTERRASCAN_BIN_PATH=${PWD}/bin/openterrascan

go test -p 1 -v ./test/e2e/validatingwebhook/...