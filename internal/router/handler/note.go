package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/tangrc99/gohelloblog/internal/dao"
	"github.com/tangrc99/gohelloblog/internal/model"
	"github.com/tangrc99/gohelloblog/internal/service"
	"github.com/tangrc99/gohelloblog/pkg/app"
	"github.com/tangrc99/gohelloblog/pkg/err_code"
	"github.com/tangrc99/gohelloblog/pkg/resp"
)

// Note
// @tag Note
type Note struct {
}

// Create handles user's note create request.
// @Summary 创建一条留言
// @tags Note
// @Accept application/x-www-form-urlencoded,application/json
// @Produce json
// @Param note_content body string true "留言内容"
// @Success 200
// @Failure 400 {object} err_code.Error
// @Failure 401 {object} err_code.Error "需要登录后才能使用 api 操作"
// @Failure 500 {object} err_code.Error
// @Router /api/notes [post]
func (r *Note) Create(ctx *gin.Context) {

	res := resp.NewResponse(ctx)
	req := service.NoteCreateRequest{}

	valid, errs := app.BindAndValid(ctx, &req)
	if !valid {
		res.SendError(err_code.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	note := model.Note{
		Content: req.NoteContent,
	}

	err := dao.CreateNote(&note)
	if err != nil {
		res.SendError(err_code.ServerError.WithDetails(errs.Errors()...))
	}

	res.SendJson(gin.H{
		"id": note.ID,
	})
}

// Update
// @Summary 更新一条留言
// @tags Note
// @Accept application/x-www-form-urlencoded,application/json
// @Produce json
// @Param id path string true "留言id"
// @Param note_content body string true "留言内容"
// @Success 200
// @Failure 400 {object} err_code.Error
// @Failure 401 {object} err_code.Error "需要登录后才能使用 api 操作"
// @Failure 500 {object} err_code.Error
// @Router /api/notes/{id} [put]
func (r *Note) Update(ctx *gin.Context) {
	res := resp.NewResponse(ctx)
	req := service.NoteCreateRequest{}

	valid, errs := app.BindAndValid(ctx, &req)
	if !valid {
		res.SendError(err_code.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	note := model.Note{
		ID:      cast.ToInt(ctx.Param("id")),
		Content: req.NoteContent,
	}

	err := dao.UpdateNote(&note)
	if err != nil {
		return
	}

	res.SendJson(gin.H{
		"id": note.ID,
	})
}

// Get
// @Summary 更新一条留言
// @tags Note
// @Produce json
// @Param id path string true "留言id"
// @Success 200 {object} model.Note
// @Failure 400 {object} err_code.Error
// @Failure 401 {object} err_code.Error "需要登录后才能使用 api 操作"
// @Failure 500 {object} err_code.Error
// @Router /api/notes/{id} [get]
func (r *Note) Get(ctx *gin.Context) {
	res := resp.NewResponse(ctx)

	note := model.Note{
		ID: cast.ToInt(ctx.Param("id")),
	}
	err := dao.GetNote(&note)
	if err != nil {
		res.SendError(err_code.ServerError.WithDetails(err.Error()))
		return
	}
	res.SendJson(note)
}

// List
// @Summary 获得留言列表
// @tags Note
// @Produce json
// @Param queries query service.NoteQueryRequest false "控制展示的范围"
// @Success 200 {object} []service.NoteDisplay
// @Failure 400 {object} err_code.Error
// @Failure 401 {object} err_code.Error "需要登录后才能使用 api 操作"
// @Failure 500 {object} err_code.Error
// @Router /api/notes/ [get]
func (r *Note) List(ctx *gin.Context) {
	res := resp.NewResponse(ctx)

	req := service.NoteQueryRequest{
		MinID: cast.ToInt(ctx.Query("min_id")),
		Limit: cast.ToInt(ctx.Query("limit")),
	}

	if req.Limit == 0 {
		req.Limit = 5
	}

	notes, err := dao.ListNote(req.MinID, req.Limit)

	if err != nil {
		res.SendError(err_code.ServerError.WithDetails(err.Error()))
		return
	}

	var dList []service.NoteDisplay
	for _, note := range notes {

		// 这里需要将时间进行格式化处理
		dList = append(dList, service.NoteDisplay{
			Time:    note.CreatedOn.Format("2006-01-02 15:04:05"),
			Content: note.Content,
		})
	}
	ctx.HTML(200, "note_list.tmpl", gin.H{
		"HEADLINE": "Notes",
		"List":     dList,
	})
}

// Delete
// @Summary 删除一条留言
// @tags Note
// @Produce json
// @Success 200
// @Failure 400 {object} err_code.Error "url结尾必须为int类型"
// @Failure 401 {object} err_code.Error "需要登录后才能使用 api 操作"
// @Failure 500 {object} err_code.Error "可能为数据库查询出现问题"
// @Router /api/notes/{id} [delete]
func (r *Note) Delete(ctx *gin.Context) {
	res := resp.NewResponse(ctx)
	note := model.Note{
		ID: cast.ToInt(ctx.Param("id")),
	}

	if note.ID == 0 {
		res.SendError(err_code.InvalidParams.WithDetails("url结尾必须为int类型"))
	}

	err := dao.DeleteNote(&note)
	if err != nil {
		res.SendError(err_code.ServerError.WithDetails(err.Error()))
		return
	}

	res.SendJson(note.ID)
}
