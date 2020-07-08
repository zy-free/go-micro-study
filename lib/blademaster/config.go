package blademaster

import "io"

type LoggerConfig struct {
	// Output is a writer where logs are written.
	// Optional. Default value is gin.DefaultWriter.
	Output io.Writer
}
