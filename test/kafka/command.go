package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	clients  int
	group    string
	version  string
	brokers  string
	assignor string
	topics   string
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kafka",
		Short: "Run kafka examples",
		RunE:  kafkaRunE,
	}
	flags := cmd.Flags()
	flags.IntVarP(&clients, "clients", "c", 3, "number of clients")
	flags.StringVar(&group, "group", "", "Kafka consumer group definition")
	flags.StringVar(&version, "version", "2.5.0", "Kafka cluster version")
	flags.StringVar(&brokers, "brokers", "work.scythefly.top", "Kafka bootstrap brokers to connect to, as a comma separated list")
	flags.StringVar(&assignor, "assignor", "range", "Consumer group partition assignment strategy (range, roundrobin, sticky)")
	flags.StringVar(&topics, "topics", "", "Kafka topics to be consumed, as a comma separated list")

	return cmd
}

func kafkaRunE(cmd *cobra.Command, args []string) error {
	if len(topics) == 0 {
		return errors.New("no topics given to be consumed")
	}
	if len(group) == 0 {
		return errors.New("no Kafka consumer group defined")
	}

	version, err := sarama.ParseKafkaVersion(version)
	if err != nil {
		return errors.WithMessage(err, "parse kafka version")
	}
	config := sarama.NewConfig()
	config.Version = version

	switch assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "roundrobin":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	case "range":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	default:
		return errors.Errorf("Unrecognized consumer group partition assignor: %s", assignor)
	}
	return nil
}
