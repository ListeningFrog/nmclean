package nmclean

import (
	"os"
	logging "github.com/op/go-logging"
	"runtime"
	"sync"
	// "sync/atomic"
	"path/filepath"
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
	ch    chan func()
	wg    sync.WaitGroup
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
		ch:    make(chan func()),
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
	p.log.Debugf("- num CPUs: %d", runtime.NumCPU())

	p.startN(1) // runtime.NumCPU())
	defer p.stop()

	err := filepath.Walk(p.dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		p.log.Debugf("> %s", path)
		stats.FilesTotal++
		return nil
	})

	return &stats, err
}


// startN starts n loops.
func (p *Nmcleaner) startN(n int) {
	for i := 0; i < n; i++ {
		p.wg.Add(1)
		go p.start()
	}
}

// start loop.
func (p *Nmcleaner) start() {
	defer p.wg.Done()
	for fn := range p.ch {
		fn()
	}
}

// stop loop.
func (p *Nmcleaner) stop() {
	close(p.ch)
	p.wg.Wait()
}
