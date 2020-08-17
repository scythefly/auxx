package etcd

import "github.com/spf13/cobra"

func newKvCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kv",
		Short: "run etcd kv examples",
		RunE:  etcdKv,
	}
	return cmd
}

func etcdKv(_ *cobra.Command, _ []string) error {
	return nil
}
