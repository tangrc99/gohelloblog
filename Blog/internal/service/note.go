package service

type NoteCreateRequest struct {
	NoteContent string `form:"note_content" binding:"required,max=100"`
}

type NoteUpdateRequest struct {
	ID          int    `form:"id"`
	NoteContent string `form:"note_content" binding:"required,max=100"`
}

type NoteDisplay struct {
	Time    string // 格式化为的输出时间
	Content string
}

type NoteQueryRequest struct {
	MinID int `json:"min_id" binding:"min=0"`       // 记录上一次返回的最大id，用于实现数据库翻页
	Limit int `json:"limit" binding:"min=5,max=10"` // 限制返回的数量
}
