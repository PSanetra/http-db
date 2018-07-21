package main

import (
	"os"
	"github.com/sirupsen/logrus"
)

func main()  {

	cmd := newCmd()

	cmd.CancelOnSignals()

	err := cmd.Execute()

	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
