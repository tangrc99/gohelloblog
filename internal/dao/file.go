package dao

import (
	"github.com/tangrc99/gohelloblog/global"
	"github.com/tangrc99/gohelloblog/internal/model"
	"time"
)

// CreateFile 在数据库中创建一条关于上传文件的记录
func CreateFile(file *model.File) error {
	file.CreateON = time.Now().UTC()
	return global.MySQL.Create(file).Error
}

func GetFile(file *model.File) error {
	//SQL: SELECT * FROM files WHERE id = ?
	return global.MySQL.Where("id = ?", file.ID).Take(file).Error
}

func UpdateFile(file *model.File) error {

	return global.MySQL.Model(file).Where("id = ?").Updates(file).Error
}

func DeleteFile(file *model.File) error {
	return global.MySQL.Delete(file).Error
}

// FindFileByExt 通过拓展名来查找相应的文件
func FindFileByExt(file *model.File) ([]model.File, error) {
	var objs []model.File

	//SQL: SELECT * FROM files WHERE file_name = 'fn'
	err := global.MySQL.Select("*").Where("ext = ?", file.Ext, file.FileName).Find(objs).Error

	return objs, err
}

func FindFileByFullName(file *model.File) ([]model.File, error) {

	var objs []model.File

	//SQL: SELECT * FROM files WHERE ext = 'txt' AND file_name = 'fn'
	err := global.MySQL.Select("*").Where("ext = ? AND file_name = ?", file.Ext, file.FileName).Find(objs).Error

	return objs, err
}

func FindFileMatches(file *model.File) ([]model.File, error) {
	var objs []model.File

	err := global.MySQL.Select("*").Where("filename REGEXP ?", file.FileName).Find(objs).Error

	return objs, err
}
