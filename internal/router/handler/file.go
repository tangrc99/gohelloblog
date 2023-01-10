package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tangrc99/gohelloblog/internal/dao"
	"github.com/tangrc99/gohelloblog/internal/model"
	"github.com/tangrc99/gohelloblog/pkg/err_code"
	"github.com/tangrc99/gohelloblog/pkg/resp"
	"path"
	"strings"
)

type File struct {
}

func (*File) Create(ctx *gin.Context) {
	res := resp.NewResponse(ctx)

	file, err := ctx.FormFile("file")
	if err != nil {
		res.SendError(err_code.InvalidParams.WithDetails(err.Error()))
		return
	}

	sp := "files/"
	ext := path.Ext(file.Filename)
	fn := strings.TrimSuffix(file.Filename, ext)
	ext = strings.TrimLeft(ext, ".")

	// 开启事务，如果创建文件写入硬盘不成功，则停止提交数据库写入
	err = dao.StartTransaction(func() error {
		fm := model.File{
			FileName: fn,
			Ext:      ext,
			Path:     sp,
		}
		e := dao.CreateFile(&fm)
		if e != nil {
			return e
		}

		e = ctx.SaveUploadedFile(file, path.Join(sp, file.Filename))

		return e
	})

	if err != nil {
		res.SendError(err_code.ServerError.WithDetails(err.Error()))
		return
	}

	res.SendJson("ok")
}
