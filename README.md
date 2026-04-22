
```
import (
    github.com/dzamyatin/process-manager/process
    "github.com/pkg/errors"
	"go.uber.org/zap"
)

signalListener := process.NewSignalListener(logger)

shutdowner := process.NewShutdownerRegistry()
shutdowner..Add("db", func() error {
    return db.Close()
})

manager := process.NewProcessShutdownerManager(
    logger,
    shutdowner,
)

manager.Run(
    ctx,
    process.NewProcessIniter(
        "tracer",
        func(ctx context.Context) (process.ProcessStarter, error) {
            if err := r.trace.Run(); err != nil {
                return nil, errors.Wrap(err, "tracer error")
            }

            return r.trace, nil
        },
    ),
    process.NewProcess("grpc server", grpcServer),
    process.NewProcess("http server", httpServer),
    process.NewProcess("signal listener", signalListener),
)

```