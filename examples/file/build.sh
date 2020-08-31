#! /bin/bash

Target="file"
DIR="$(cd "$(dirname "$0")" && pwd)"
BINARY="${DIR}/bin/${Target}"

mkdir -p ${DIR}/bin

funcBuildLinuxAmd64() {
  echo "build ${Target} on linux amd64..."
  return `GOOS=linux GOARCH=amd64 go build -o ${BINARY} .`
}

funcBuildLinuxAmd64
if [ $? -eq 0 ]; then
  rm -rf /usr/local/share/nginx/download/binary/${Target}
  mv ${BINARY} /usr/local/share/nginx/download/binary
fi