module auxx

go 1.14

require (
	code.google.com/p/graphics-go/graphics v0.0.0-00010101000000-000000000000
	github.com/Shopify/sarama v1.26.4
	github.com/coreos/etcd v3.3.22+incompatible
	github.com/deckarep/golang-set v1.7.1
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.1.1 // indirect
	github.com/jasonlvhit/gocron v0.0.0-20200423141508-ab84337f7963
	github.com/jolestar/go-commons-pool/v2 v2.1.1
	github.com/lucas-clemente/quic-go v0.17.3
	github.com/mattn/go-sqlite3 v1.14.0
	github.com/pkg/errors v0.9.1
	github.com/scythefly/orb v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.0.0
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace (
	code.google.com/p/graphics-go/graphics => ../pkg/graphics-go/graphics
	github.com/scythefly/orb => ../pkg/orb
)
