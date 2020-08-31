#! /bin/bash

Target="auxx"

ChangeLog=""
Version="1.0.1"
BuildTime=$(date +'%Y.%m.%d %H:%M:%S')
subVersion=`git rev-parse --short HEAD`

DIR="$(cd "$(dirname "$0")" && pwd)"
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

if [ a"$1" = "a" ]; then
  funcBuild
fi

if [ a"$1" = "amake" ]; then
  funcBuild
  ${DIR}/bin/${Target}
fi

if [ a"$1" = "arun" ]; then
  ${DIR}/bin/${Target}
fi

if [ a"$1" = "alinux" ]; then
  funcBuildLinuxAmd64
  rm /usr/local/share/nginx/build/${Target}
  mv ${DIR}/bin/${Target} /usr/local/share/nginx/build/
fi
