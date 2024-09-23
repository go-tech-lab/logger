package test

import (
	"github.com/go-tech-lab/framework/logger/src/boot"
	_ "github.com/stretchr/testify/require"
	"testing"
)

func TestLogger_CreateFromFile(t *testing.T) {
	logger := boot.CreateLoggerFromFile("/Users/heguang/Documents/shopee/code/src/logger/config/logger.ini")
	logger.Info("dsfdsf")
}
