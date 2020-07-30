#! /bin/bash
DIR="$(cd "$(dirname "$0")" && pwd)"
Target="auxx"

# go build -v -a -ldflags="-w -s " -trimpath .
# go build -v -a -ldflags="-w -s "  .

cd ${DIR}
mkdir -p bin/

go build -o "${DIR}/bin/${Target}" .
