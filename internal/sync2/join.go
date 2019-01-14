// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information

package sync2

import (
	"time"

	"github.com/zeebo/errs"
)

// RunJointly runs the given jobs concurrently. As soon as one finishes,
// the timeout begins ticking. If all jobs finish before the deadline
// RunJointly returns the combined errors. If some jobs time out, a timeout
// error is returned for them.
func RunJointly(timeout time.Duration, jobs ...func() error) error {
	ch := make(chan error, len(jobs))

	for _, job := range jobs {
		go func(job func() error) {
			// run the job but turn panics into errors
			err := func() (err error) {
				defer func() {
					if rec := recover(); rec != nil {
						err = errs.New("panic: %+v", rec)
					}
				}()
				return job()
			}()
			// send the error
			ch <- err
		}(job)
	}

	errgroup := make([]error, 0, len(jobs))
	errgroup = append(errgroup, <-ch)
	timer := time.NewTimer(timeout)
	defer timer.Stop()

loop:
	for len(errgroup) < len(jobs) {
		select {
		case err := <-ch:
			errgroup = append(errgroup, err)
		case <-timer.C:
			errgroup = append(errgroup,
				errs.New("%d jobs timed out", len(jobs)-len(errgroup)))
			break loop
		}
	}

	return errs.Combine(errgroup...)
}
