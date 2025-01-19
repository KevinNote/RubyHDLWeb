package ruby

import (
	"net/http"
	"os"

	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/KevinZonda/RubyDHLWeb/controller/types"
	"github.com/KevinZonda/RubyDHLWeb/lib/taskdir"
	"github.com/KevinZonda/RubyDHLWeb/shared"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Controller struct{}

var _ types.IController = (*Controller)(nil)

func (c *Controller) Init(r gin.IRouter) {
	r.POST("/compile", c.Compile)
}

type CompileReq struct {
	TaskId string `json:"task_id"`
	Code   string `json:"code"`
}

type CompileRes struct {
	TaskId     string `json:"task_id"`
	CompileErr string `json:"compile_err"`
	Rbs        string `json:"rbs"`
}

func (ctr *Controller) Compile(c *gin.Context) {
	var req CompileReq
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parse request body failed"})
		return
	}

	task := taskdir.TaskInfo{
		Id: req.TaskId,
	}

	if task.Id == "" || (uuid.Validate(task.Id) != nil) {
		task, err = shared.TaskDir.NewTask()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "create task failed"})
			return
		}
	} else {
		task.Dir = shared.TaskDir.JoinTask(task.Id)
	}

	rby := task.File("main.rby")

	iox.WriteAllText(rby, req.Code)

	rcOut, err := shared.Ruby.Rc(task.Dir, rby)
	if err != nil {
		c.JSON(http.StatusBadRequest, CompileRes{
			TaskId:     task.Id,
			CompileErr: rcOut,
		})
		return
	}

	rbs, err := iox.ReadAllText(task.File("current.rbs"))
	if err != nil {
		c.JSON(http.StatusBadRequest, CompileRes{
			TaskId:     task.Id,
			CompileErr: "current.rbs not found",
		})
		return
	}

	c.JSON(http.StatusOK, CompileRes{
		TaskId: task.Id,
		Rbs:    rbs,
	})
}

type RunReq struct {
	TaskId string `json:"task_id"`
	Input  string `json:"input"`
}

type RunRes struct {
	Output string `json:"output"`
	Err    string `json:"err"`
}

func (ctr *Controller) Run(c *gin.Context) {
	var req RunReq
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parse request body failed"})
		return
	}

	task := taskdir.TaskInfo{
		Id: req.TaskId,
	}

	if task.Id == "" || (uuid.Validate(task.Id) != nil) {
		c.JSON(http.StatusBadRequest, RunRes{
			Err: "task not found",
		})
		return
	}

	task.Dir = shared.TaskDir.JoinTask(task.Id)

	if _, err := os.Stat(task.File("current.rbs")); os.IsNotExist(err) {
		c.JSON(http.StatusBadRequest, RunRes{
			Err: "compile first",
		})
		return
	}

	out, err := shared.Ruby.Re(task.Dir, task.File("current.rbs"), req.Input)
	if err != nil {
		c.JSON(http.StatusBadRequest, RunRes{
			Err: out,
		})
		return
	}

	c.JSON(http.StatusOK, RunRes{
		Output: out,
	})
}
