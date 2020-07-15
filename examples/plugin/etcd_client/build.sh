#! /bin/bash

Version="1.0.5"
BuildTime=$(date +'%Y.%m.%d %H:%M:%S')
subVersion=$(cat .version)
((subVersion++))
echo $subVersion >.version

DIR="$(cd "$(dirname "$0")" && pwd)"
cd "${DIR}"

BIN="../../../bin"

LDFLAGS="
  -X 'auxx/examples/plugin/etcd_client/version.Built=${BuildTime}'
  -X 'auxx/examples/plugin/etcd_client/version.Version=${Version}.${subVersion}'
  -X 'auxx/examples/plugin/etcd_client/version.SubVersion=${subVersion}'
  -X 'auxx/examples/plugin/etcd_client/version.ChangeLog=${ChangeLog}'
  -s
  -w
"

go build -ldflags "${LDFLAGS}" -buildmode=plugin -o plugin_etcd_client.so plugin_etcd_client.go

rm ${BIN}/plugin_etcd_client.so
mv plugin_etcd_client.so ${BIN}
