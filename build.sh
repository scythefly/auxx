#! /bin/bash

Target="auxx"

ChangeLog=""
Version="1.0.1"
BuildTime=$(date +'%Y.%m.%d %H:%M:%S')
subVersion=`git rev-parse --short HEAD`

DIR="$(cd "$(dirname "$0")" && pwd)"
BINARY="${DIR}/bin/${Target}"
cd "${DIR}"
mkdir -p bin/

funcBuild() {
  LDFLAGS="
  -X '${Target}/version.Built=${BuildTime}'
  -X '${Target}/version.Version=${Version}.${subVersion}'
  -X '${Target}/version.ChangeLog=${ChangeLog}'
  -s
  -w
"
  echo "build ${Target} ..."
  go build -ldflags "${LDFLAGS}" -o "${DIR}/bin/${Target}" .
}

funcBuildLinuxAmd64() {
  LDFLAGS="
  -X '${Target}/version.Built=${BuildTime}'
  -X '${Target}/version.Version=${Version}.${subVersion}'
  -X '${Target}/version.ChangeLog=${ChangeLog}'
  -s
  -w
"
  echo "build ${Target} on linux amd64..."
  GOOS=linux GOARCH=amd64 go build -ldflags "${LDFLAGS}" -o "${DIR}/bin/${Target}" .

}

case "$1" in
make)
  funcBuild
  if [ $? -eq 0 ];then
    ${BINARY} ${@:2}
  fi
  ;;
run)
  ${BINARY} ${@:2}
  ;;
*)
  funcBuild
  if [ $? -eq 0 ];then
    ${BINARY} version
  fi
  ;;
esac

exit 0
