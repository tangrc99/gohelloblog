package service

type ArtReq struct {
	Title    string `form:"title" binding:"required,max=100"`
	Author   string `form:"author"`                                  // 文章作者
	FileName string `form:"filename"`                                // 文章对应文件名
	Content  []byte `form:"content" binding:"required,max=16000000"` // 文章内容
}

type ListReq struct {
	Page   int    `form:"page"`
	LastId string `form:"last_id"`
}
