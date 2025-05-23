package main

import (
	"github.com/riza/linx/internal/banner"
	"github.com/riza/linx/internal/options"
	"github.com/riza/linx/internal/scanner"
	"github.com/riza/linx/pkg/logger"
)

const Version = "v0.2.0"

func main() {
	banner.Show(Version)

	opts, err := options.Get().Parse()
	if err != nil {
		logger.Get().Fatal(err)
	}

	scanner := scanner.NewScanner(opts)
	err = scanner.Run()
	if err != nil {
		logger.Get().Error(err)
	}
}
