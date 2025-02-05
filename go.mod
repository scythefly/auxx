module auxx

go 1.16

replace (
code.google.com/p/graphics-go/graphics => github.com/kakami/graphics-go/graphics v0.0.0-20200817091218-b0f49abb9380
github.com/kakami/pkg => /Users/iuz/work/kakami/pkg
)
require (
	code.google.com/p/graphics-go/graphics v0.0.0-00010101000000-000000000000
	fyne.io/fyne v1.4.3
	github.com/aws/aws-sdk-go v1.38.0
	github.com/baidu/go-lib v0.0.0-20210316014414-55daa983069e
	github.com/deckarep/golang-set v1.7.1
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.5.1
	github.com/gorilla/websocket v1.4.2
	github.com/hashicorp/golang-lru v0.5.4
	github.com/inconshreveable/go-update v0.0.0-20160112193335-8152e7eb6ccf
	github.com/kakami/flock v0.9.0
	github.com/kakami/gocron v0.0.0-20201221071540-5e96d754babc
	github.com/kakami/pkg v0.0.0-20201210010425-144611d45889
	github.com/lucas-clemente/quic-go v0.19.3
	github.com/mattn/go-sqlite3 v1.14.6
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.3
	github.com/valyala/fasthttp v1.22.0
	go.uber.org/atomic v1.7.0
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20210226101413-39120d07d75e
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210317225723-c4fcb01b228e
	golang.zx2c4.com/wireguard v0.0.20201118
	google.golang.org/grpc v1.36.0
)
