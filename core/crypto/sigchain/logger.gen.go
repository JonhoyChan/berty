// Code generated by github.com/berty/berty/core/.scripts/generate-logger.sh

package sigchain

import "go.uber.org/zap"

func logger() *zap.Logger {
	return zap.L().Named("core.crypto.sigchain")
}