package logger

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

// ChannelHook is a hook to receive logs in test scenarios
type ChannelHook struct {
	Entries chan logrus.Entry
}

// ChannelHook is a logrus.Hook
var _ logrus.Hook = &ChannelHook{}

// NewChannelLogger creates a discarding logger and installs the channel hook.
func NewChannelLogger(size int) (*logrus.Logger, *ChannelHook) {

	log := logrus.New()
	log.Out = ioutil.Discard

	hook := &ChannelHook{
		Entries: make(chan logrus.Entry, size),
	}
	log.Hooks.Add(hook)

	return log, hook
}

func (t *ChannelHook) Fire(e *logrus.Entry) error {
	t.Entries <- *e
	return nil
}

func (t *ChannelHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
