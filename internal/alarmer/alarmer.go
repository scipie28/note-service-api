package alarmer

//go:generate mockgen -destination=mocks/mock_alarmer.go -package=mocks . Alarmer

import (
	"errors"
	"time"
)

// Alarmer ...
type Alarmer interface {
	Alarm() <-chan struct{}
	Init() error
	Close()
}

type alarmer struct {
	duration time.Duration
	alarm    chan struct{}
	end      chan struct{}
	isInit   bool
}

func NewAlarmer(duration time.Duration) (Alarmer, error) {
	if duration <= 0 {
		return nil, errors.New("invalid duration")
	}

	return &alarmer{
		duration: duration,
		alarm:    make(chan struct{}),
		end:      make(chan struct{}),
		isInit:   true,
	}, nil
}

func (a *alarmer) Init() error {
	if !a.isInit {
		return errors.New("the alarm has already been initialized")
	}

	go func() {
		ticker := time.NewTicker(a.duration)
		defer ticker.Stop()
		defer close(a.end)
		defer close(a.alarm)

		for {
			select {
			case <-ticker.C:
				select {
				case a.alarm <- struct{}{}:
				default:
				}

			case <-a.end:
				return
			}
		}
	}()

	a.isInit = false

	return nil
}

func (a *alarmer) Alarm() <-chan struct{} {
	return a.alarm
}

func (a *alarmer) Close() {
	a.end <- struct{}{}
}
