package taskdir

import (
	"log"
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

func (t *TaskDir) idToDir(id string) string {
	return filepath.Join(t.dir, id)
}

func (t *TaskDir) initDir(info TaskInfo) error {
	if t.InitDir != nil {
		return t.InitDir(info)
	}
	return nil
}

func (t *TaskDir) NewTask() (info TaskInfo, err error) {
	info.Id = uuid.New().String()
	info.Dir = t.idToDir(info.Id)
	err = os.MkdirAll(info.Dir, 0755)
	if err != nil {
		return
	}
	err = t.initDir(info)
	return
}

func (t *TaskDir) RemoveTask(taskId string) error {
	return os.RemoveAll(t.idToDir(taskId))
}

func (t *TaskDir) JoinTask(taskId string) string {
	path := t.idToDir(taskId)
	_, err := os.Stat(path)
	log.Println(err)
	if err == nil {
		return path
	}
	if !os.IsNotExist(err) {
		// TODO: FATAL
		return path
	}

	if info, err := t.NewTask(); err == nil {
		return info.Dir
	}
	return path
}
