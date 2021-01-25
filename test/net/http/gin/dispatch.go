package gin

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/kakami/gocron"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"auxx/types"
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

	go schedule()

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

func schedule() {
	// for i := 0; i < 3; i++ {
	types.G.Go(func() error {
		s := gocron.NewScheduler()
		// s.Every(5).Seconds().Do(scheduleDo)
		// s.Every(5).Seconds().StartImmediately().Do(scheduleDoImmediately)
		s.Every(100 * time.Second).From(time.Now().Add(10 * time.Second)).Do(scheduleDo)

		s.StartAsync()
		// fmt.Println(tt)
		// s.Every(5).Seconds().Do(scheduleDoAfterStart)
		// s.Every(5).Seconds().StartImmediately().Do(scheduleDoAfterStartImmediately)
		s.Start()
		return nil
		// t := time.NewTicker(time.Second)
		// for {
		// 	select {
		// 	case <-t.C:
		// 		go scheduleDo()
		// 	}
		// }
	})
	// }
	types.G.Wait()
}

func scheduleDo() {
	fmt.Println("--- do ---", time.Now().Second())
}

func scheduleDoImmediately() {
	fmt.Println("--- do immediately ---", time.Now().Second())
}

func scheduleDoAfterStart() {
	fmt.Println("--- do after start ---", time.Now().Second())
}

func scheduleDoAfterStartImmediately() {
	fmt.Println("--- do after start immediately ---", time.Now().Second())
}
