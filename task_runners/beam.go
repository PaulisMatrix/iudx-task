package main

import (
	"context"
	"fmt"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/xlang/kafkaio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/runners/prism"

	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/local"
)

func init() {
	// register your functions here to be run by the worker runner
	register.Function0x0(func() {})
	register.Emitter1[string]()
}

func main() {
	// beam.Init() is an initialization hook that must be called on startup.
	beam.Init()

	// Create the Pipeline object and root scope.
	p := beam.NewPipeline()
	s := p.Root()

	// Connect to kafka source
	expansionAddr := "localhost:9092"
	bootstrapServer := "bootstrap-server:9092"
	topic := "data-kaveri"

	pcol := kafkaio.Read(s, expansionAddr, bootstrapServer, []string{topic},
		kafkaio.MaxNumRecords(100), kafkaio.CommitOffsetInFinalize(true))

	fmt.Print(pcol)

	// Define the data processing pipeline here.

	// Run the pipeline on the default direct/prism runner.
	prism.Execute(context.Background(), p)
}
