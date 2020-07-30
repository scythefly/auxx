package gin

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gin",
		Short: "Run gin examples",
		RunE:  ginRun,
	}
	return cmd
}

const (
	rewriteKey   = "inner_rewriteKey"
	rewriteValue = "inner_rewriteValue"
)

var (
	r *gin.Engine
)

func ginRun(_ *cobra.Command, _ []string) error {
	r = gin.Default()

	r.GET("/usr/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "handleUsr[%s] - %s", name, c.Request.URL.Path)
	})
	r.GET("/flv/*action", handleFlv)
	r.GET("/m3u8/*action", handleM3u8)

	r.NoRoute(dispatch)

	return r.Run(":30001")
}

func dispatch(c *gin.Context) {
	fmt.Println(">>>>> dispatch")
	parts := strings.Split(c.Request.URL.Path, ".")
	// strings.HasSuffix()
	fmt.Println(parts)
	if len(parts) > 0 {
		switch parts[len(parts)-1] {
		case "flv":
			fmt.Println(">>>> to /flv/xxx")
			c.Request.URL.Path = "/flv" + c.Request.URL.Path
		case "m3u8":
			fmt.Println(">>>> to /m3u8/xxx")
			c.Request.URL.Path = "/m3u8" + c.Request.URL.Path
		default:
			handleRedirect(c)
			return
		}
	}
	r.HandleContext(c)
}

func response(_ *gin.Context) {
	fmt.Println(">>>>> response")
}

func handleFlv(c *gin.Context) {
	fmt.Println(">>>>> handleFlv")
	c.String(200, "handleFlv - %s!!!\n", c.Request.URL.Path)
}

func handleM3u8(c *gin.Context) {
	fmt.Println(">>>>> handleM3u8")
	c.String(200, "handleM3u8 - %s!!!\n", c.Request.URL.Path)
}

func handleRedirect(c *gin.Context) {
	fmt.Println(">>>>> handleRedirect")
	c.String(200, "handleRedirect - %s!!!\n", c.Request.URL.Path)
}
