package command

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func newHttpCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "http",
		Short: `run http server`,
		Run: func(cmd *cobra.Command, args []string) {
			r := gin.Default()
			r.Any("/", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "ok ok",
				})
			})
			r.Run("localhost:8080")
		},
	}

	return cmd
}
