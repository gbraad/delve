package native

import (
	"errors"
	"github.com/go-delve/delve/pkg/proc"
)

func (t *nativeThread) findHardwareBreakpoint() (*proc.Breakpoint, error) {
	// TODO(alexsaezm) Implement findHardwareBreakpoint
	panic("Unimplemented findHardwareBreakpoint method in threads_linux_ppc64le.go")
	return nil, nil
}

func (t *nativeThread) writeHardwareBreakpoint(addr uint64, wtype proc.WatchType, idx uint8) error {
	// TODO(alexsaezm) Implement writeHardwareBreakpoint
	wpstate, err := t.getWatchpoints()
	if err != nil {
		return err
	}
	if idx >= wpstate.num {
		return errors.New("hardware breakpoints exhausted")
	}
	return nil
}

func (t *nativeThread) clearHardwareBreakpoint(addr uint64, wtype proc.WatchType, idx uint8) error {
	// TODO(alexsaezm) Implement clearHardwareBreakpoint
	panic("Unimplemented clearHardwareBreakpoint method in threads_linux_ppc64le.go")
	return nil
}
