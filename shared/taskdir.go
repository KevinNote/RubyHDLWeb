package shared

import (
	"path/filepath"

	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/KevinZonda/RubyDHLWeb/lib/taskdir"
)

var TaskDir *taskdir.TaskDir

func initTaskDir() {
	cfg := GetConfig()
	TaskDir = taskdir.NewTaskDir(cfg.TaskDir, func(info taskdir.TaskInfo) error {
		if len(Prelude) > 0 {
			return iox.WriteAllBytes(filepath.Join(info.Dir, "prelude.rby"), Prelude)
		}
		return nil
	})
}
