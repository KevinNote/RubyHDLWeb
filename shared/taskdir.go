package shared

import (
	"os"
	"path/filepath"

	"github.com/KevinZonda/RubyDHLWeb/lib/taskdir"
)

var TaskDir *taskdir.TaskDir

func initTaskDir() {
	cfg := GetConfig()
	TaskDir = taskdir.NewTaskDir(cfg.TaskDir, func(info taskdir.TaskInfo) error {
		if len(Prelude) > 0 {
			return os.Symlink(PreludePath, filepath.Join(info.Dir, "prelude.rby"))
			//return iox.WriteAllBytes(filepath.Join(info.Dir, "prelude.rby"), Prelude)
		}
		return nil
	})
}
