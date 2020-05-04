package internal

import (
	"context"

	"github.com/anz-bank/pkg/log"
)

type TestHook struct {
	Entries []log.LogEntry
}

func (h *TestHook) OnLogged(entry *log.LogEntry) error {
	h.Entries = append(h.Entries, *entry)
	return nil
}

func NewTestContextWithLoggerHook() (context.Context, *TestHook) {
	loghook := TestHook{}
	ctx := log.WithConfigs(log.AddHooks(&loghook)).Onto(context.Background())
	return ctx, &loghook
}
