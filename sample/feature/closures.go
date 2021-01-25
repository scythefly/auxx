package feature

import "github.com/spf13/cobra"

var confBuf = []byte(`{
  "Version": "init version",
  "ServerId": {
    "publish.scythefly.top": [
      {
        "path": "app1/.*",
        "config": {
          "slice": "on",
          "record": "on"
        }
      },
      {
        "path": ".*",
        "config": {
          "slice": "on",
          "record": "off"
        }
      }
    ]
  }
}`)

func newClosuresCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "closures",
		Short: "run closures examples",
		Run:   runClosures,
	}

	return cmd
}

func runClosures(_ *cobra.Command, _ []string) {
}
