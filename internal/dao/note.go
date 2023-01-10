package dao

import (
	"github.com/tangrc99/gohelloblog/global"
	"github.com/tangrc99/gohelloblog/internal/model"
	"time"
)

func CreateNote(obj *model.Note) error {
	obj.CreatedOn = time.Now().UTC()
	return global.MySQL.Create(obj).Error
}

func UpdateNote(obj *model.Note) error {

	// 将需要修改的内容放在一个表中，然后传给 gorm
	values := map[string]interface{}{}

	if obj.Content != "" {
		values["content"] = obj.Content
	}

	return global.MySQL.Model(obj).Where("id = ?", obj.ID).Updates(values).Error
}

func GetNote(obj *model.Note) error {

	// SQL: SELECT * FROM notes WHERE id = obj.ID
	return global.MySQL.Where("id = ?", obj.ID).Take(obj).Error
}

func DeleteNote(obj *model.Note) error {
	return global.MySQL.Delete(obj).Error
}

func ListNote(minId, limit int) ([]model.Note, error) {

	var obj []model.Note

	// SQL: SELECT * FROM(SELECT * FROM notes WHERE id > minId LIMIT limit) as res ORDER BY id DESC;
	err := global.MySQL.Model(model.Note{}).Where("id > ?", minId).Limit(limit).Find(&obj).Order("id desc").Error

	return obj, err
}
