package log

import (
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/kakami/pkg/zlog"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "log",
		Short: "Run log examples",
		RunE:  runLog,
	}

	return cmd
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func runLog(_ *cobra.Command, _ []string) error {
	lw := zlog.NewTimeFileLogWriter("zzzlog.log", "M", 3)
	if lw == nil {
		return errors.New("create file log failed")
	}

	encoderConf := &zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapcore.DebugLevel)
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(*encoderConf),
		zapcore.AddSync(lw),
		atomicLevel,
	)
	// log := zap.New(core).WithOptions(zap.AddCaller()).With(zap.String("field", "string")).Sugar()
	log := zap.New(core).WithOptions(zap.AddCaller()).Sugar()

	for i := 0; i < 100; i++ {
		log.Infof("info log 1111111111111")
		// log.Debugf("debug log 22222222222")
		// log.Warnf("warn log 3333333333333")
		// log.Errorf("error log 44444444444")
	}

	// ticker := time.NewTicker(100 * time.Millisecond)
	// var cnt int
	// for {
	// 	select {
	// 	case t := <-ticker.C:
	// 		cnt++
	// 		log.Info(t)
	// 		if cnt > 100 {
	// 			goto END
	// 		}
	// 	}
	// }
	// END:
	return nil
}
