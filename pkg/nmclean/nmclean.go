package nmclean

import (
  logging "github.com/op/go-logging"
)

// Stats for a nmclean.
type Stats struct {
	FilesTotal   int64
	FilesRemoved int64
	SizeRemoved  int64
}

type Nmcleaner struct {
	dir   string
	log   * logging.Logger
}

// Option function.
type Option func(*Nmcleaner)

// WithDir option.
func WithDir(s string) Option {
	return func(v *Nmcleaner) {
		v.dir = s
	}
}

// New with the given options.
func New(log * logging.Logger, options ...Option) *Nmcleaner {
	v := &Nmcleaner{
		dir:   ".",
		log:   log,
	}

	for _, o := range options {
		o(v)
	}

	return v
}


// Prune performs the pruning.
func (p *Nmcleaner) Nmclean() (*Stats, error) {
	var stats Stats;

	p.log.Info("Hello from nmcleaner!")

	return &stats, nil
}
