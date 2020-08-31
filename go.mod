module auxx

go 1.14

replace (
	code.google.com/p/graphics-go/graphics => github.com/kakami/graphics-go/graphics v0.0.0-20200817091218-b0f49abb9380
	fyne.io/fyne => github.com/fyne-io/fyne v1.3.3
	github.com/valyala/fasthttp => github.com/kakami/fasthttp v1.16.1
)

require (
	code.google.com/p/graphics-go/graphics v0.0.0-00010101000000-000000000000
	fyne.io/fyne v0.0.0-00010101000000-000000000000
	github.com/Messier78/gocron v0.1.2
	github.com/Shopify/sarama v1.27.0
	github.com/baidu/go-lib v0.0.0-20200321100322-ccd61749c524
	github.com/coreos/etcd v3.3.22+incompatible
	github.com/deckarep/golang-set v1.7.1
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.1.1 // indirect
	github.com/inconshreveable/go-update v0.0.0-20160112193335-8152e7eb6ccf
	github.com/jasonlvhit/gocron v0.0.1
	github.com/jolestar/go-commons-pool/v2 v2.1.1
	github.com/lucas-clemente/quic-go v0.17.3
	github.com/mattn/go-sqlite3 v1.14.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.0.0
	github.com/valyala/fasthttp v0.0.0-00010101000000-000000000000
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	sigs.k8s.io/yaml v1.2.0 // indirect
)
