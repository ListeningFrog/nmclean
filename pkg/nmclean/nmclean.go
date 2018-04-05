package nmclean

import (
  logging "github.com/op/go-logging"
)

type Nmcleaner struct {
	dir   string
	log   * logging.Logger
}
