package taskdir

import (
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type TaskInfo struct {
	Id  string
	Dir string
}

func (t TaskInfo) File(name string) string {
	return filepath.Join(t.Dir, name)
}

type TaskDir struct {
	dir     string
	InitDir func(info TaskInfo) error
}

func NewTaskDir(dir string, initDir func(info TaskInfo) error) *TaskDir {
	return &TaskDir{
		dir:     dir,
		InitDir: initDir,
	}
}

func (t *TaskDir) NewTask() (info TaskInfo, err error) {
	info.Id = uuid.New().String()
	info.Dir = filepath.Join(t.dir, info.Id)
	err = os.MkdirAll(info.Dir, 0755)
	if err != nil {
		return
	}
	if t.InitDir != nil {
		err = t.InitDir(info)
	}
	return
}

func (t *TaskDir) RemoveTask(taskId string) error {
	return os.RemoveAll(filepath.Join(t.dir, taskId))
}

func (t *TaskDir) JoinTask(taskId string) string {
	return filepath.Join(t.dir, taskId)
}
