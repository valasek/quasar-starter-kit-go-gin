// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package main

import (
	"github.com/sirupsen/logrus"
	"github.com/valasek/quasar-starter-kit-go-gin/server/cmd"
	"github.com/valasek/quasar-starter-kit-go-gin/server/logger"
)

func main() {
	logger.Log = logrus.New()
	cmd.Execute()
}
