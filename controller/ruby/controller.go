package ruby

import (
	"net/http"
	"os"

	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/KevinZonda/RubyDHLWeb/controller/types"
	rbsP "github.com/KevinZonda/RubyDHLWeb/lib/rbs"
	"github.com/KevinZonda/RubyDHLWeb/lib/taskdir"
	"github.com/KevinZonda/RubyDHLWeb/shared"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Controller struct{}

var _ types.IController = (*Controller)(nil)

func (c *Controller) Init(r gin.IRouter) {
	r.POST("/ruby/compile", c.Compile)
	r.POST("/ruby/run", c.Run)
	r.POST("/ruby/viz", c.Viz)
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
		task = shared.TaskDir.JoinTask(task.Id)
	}

	rby := task.File("main.rby")

	iox.WriteAllText(rby, req.Code)

	rcOut, err := shared.Ruby.Rc(task.Dir, "main.rby")
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
	TaskId string `json:"task_id"`
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

	task = shared.TaskDir.JoinTask(task.Id)

	if _, err := os.Stat(task.File("current.rbs")); os.IsNotExist(err) {
		c.JSON(http.StatusBadRequest, RunRes{
			TaskId: task.Id,
			Err:    "compile first",
		})
		return
	}

	out, err := shared.Ruby.Re(task.Dir, "current.rbs", req.Input)
	if err != nil {
		c.JSON(http.StatusBadRequest, RunRes{
			TaskId: task.Id,
			Err:    out,
		})
		return
	}

	c.JSON(http.StatusOK, RunRes{
		TaskId: task.Id,
		Output: out,
	})
}

func (ctr *Controller) Viz(c *gin.Context) {
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

	task = shared.TaskDir.JoinTask(task.Id)

	rbsBs, err := iox.ReadAllText(task.File("current.rbs"))
	if err != nil {
		c.JSON(http.StatusBadRequest, RunRes{
			TaskId: task.Id,
			Err:    "current.rbs not found",
		})
		return
	}

	circuit := rbsP.ParseCircuit(string(rbsBs))
	dot := rbsP.VizToDot(circuit)

	c.JSON(http.StatusOK, RunRes{
		TaskId: task.Id,
		Output: dot,
	})

}
