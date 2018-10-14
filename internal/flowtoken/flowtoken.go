package flowtoken

import "sync/atomic"

type Config struct{}

type FlowToken struct {
	cfg *Config

	successes uint32
	fails     uint32
}

func New(cfg *Config) *FlowToken {
	return &FlowToken{
		cfg: cfg,
	}
}

func (ft *FlowToken) Allow() bool {
	if ft == nil {
		return false
	}
	panic("not implemented")
}

func (ft *FlowToken) ReportSuccess() {
	if ft == nil {
		return
	}
	atomic.AddUint32(&ft.successes, 1)
}

func (ft *FlowToken) ReportFailure() {
	if ft == nil {
		return
	}
	atomic.AddUint32(&ft.fails, 1)
}
