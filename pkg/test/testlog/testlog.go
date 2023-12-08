package testlog

import (
	"log/slog"

	"github.com/onsi/ginkgo/v2"
	"github.com/open-panoptes/opni/pkg/logger"
	"github.com/open-panoptes/opni/pkg/test/testruntime"
	"github.com/samber/lo"
)

var Log = lo.TernaryF(testruntime.IsTesting, func() *slog.Logger {
	return logger.New(logger.WithLogLevel(logger.DefaultLogLevel.Level()), logger.WithWriter(ginkgo.GinkgoWriter)).WithGroup("test")
}, func() *slog.Logger {
	return logger.New(logger.WithLogLevel(logger.DefaultLogLevel.Level())).WithGroup("test")
})
