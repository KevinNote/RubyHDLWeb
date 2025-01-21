package shared

import (
	"os"
	"path/filepath"

	"github.com/KevinZonda/RubyDHLWeb/lib/taskdir"
)

var TaskDir *taskdir.TaskDir

func initTaskDir() {
	cfg := GetConfig()
	preludePath := cfg.Prelude
	var taskInit func(info taskdir.TaskInfo) error
	if len(preludePath) > 0 {
		taskInit = func(info taskdir.TaskInfo) error {
			return os.Symlink(preludePath, filepath.Join(info.Dir, "prelude.rby"))
		}
	}
	TaskDir = taskdir.NewTaskDir(cfg.TaskDir, taskInit)
}
