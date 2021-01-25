module auxx

go 1.15

replace (
	code.google.com/p/graphics-go/graphics => github.com/kakami/graphics-go/graphics v0.0.0-20200817091218-b0f49abb9380
	fyne.io/fyne => github.com/fyne-io/fyne v1.4.3
	github.com/aws/aws-sdk-go => github.com/kakami/aws-sdk-go v1.34.14
	github.com/valyala/fasthttp => github.com/kakami/fasthttp v1.16.1
)

require (
	code.google.com/p/graphics-go/graphics v0.0.0-00010101000000-000000000000
	fyne.io/fyne v0.0.0-00010101000000-000000000000
	github.com/aws/aws-sdk-go v0.0.0-00010101000000-000000000000
	github.com/baidu/go-lib v0.0.0-20200819072111-21df249f5e6a
	github.com/deckarep/golang-set v1.7.1
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.2
	github.com/google/go-cmp v0.5.2 // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/hashicorp/golang-lru v0.5.1
	github.com/inconshreveable/go-update v0.0.0-20160112193335-8152e7eb6ccf
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kakami/flock v0.9.0
	github.com/kakami/gocron v0.0.0-20201221071540-5e96d754babc
	github.com/kakami/pkg v0.0.0-20201210010425-144611d45889
	github.com/klauspost/compress v1.11.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lucas-clemente/quic-go v0.18.1
	github.com/mattn/go-sqlite3 v1.14.4
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.1
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/valyala/fasthttp v0.0.0-00010101000000-000000000000
	go.uber.org/atomic v1.7.0
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/grpc v1.21.1
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)
