package flag

import (
	"flag"

	"github.com/sirupsen/logrus"
)

func Fatal(args ...any) {
	flag.Usage()
	logrus.Fatal(args...)
}
