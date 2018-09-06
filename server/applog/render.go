package applog

import (
	"github.com/tsaikd/KDGoLib/webutil/render"
)

// ConsoleRender implement render.Error interface
type ConsoleRender struct {
	err error
}

func (t *ConsoleRender) Error(err error) {
	if err == nil {
		return
	}

	t.err = err
	Trace(err)
}

// GetError return last error
func (t ConsoleRender) GetError() error {
	return t.err
}

var _ render.Error = &ConsoleRender{}
