//go:build cmount && ((linux && cgo) || (darwin && cgo) || (freebsd && cgo) || (openbsd && cgo) || windows) && (!race || !windows)
// +build cmount
// +build linux,cgo darwin,cgo freebsd,cgo openbsd,cgo windows
// +build !race !windows

// FIXME this doesn't work with the race detector under Windows either
// hanging or producing lots of differences.

package cmount

import (
	"runtime"
	"testing"

	"github.com/rclone/rclone/fstest/testy"
	"github.com/rclone/rclone/vfs/vfstest"
)

func TestMount(t *testing.T) {
	// Disable tests under macOS and the CI since they are locking up
	if runtime.GOOS == "darwin" {
		testy.SkipUnreliable(t)
	}
	vfstest.RunTests(t, false, mount)
}
