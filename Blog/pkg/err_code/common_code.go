package err_code

var (
	Success                   = registryError(0, "成功")
	ServerError               = registryError(10000000, "服务内部错误")
	InvalidParams             = registryError(10000001, "入参错误")
	NotFound                  = registryError(10000002, "找不到")
	UnauthorizedAuthNotExist  = registryError(10000003, "鉴权失败，找不到对应的user和password")
	UnauthorizedTokenError    = registryError(10000004, "鉴权失败，JWT错误")
	UnauthorizedTokenTimeout  = registryError(10000005, "鉴权失败，JWT超时")
	UnauthorizedTokenGenerate = registryError(10000006, "鉴权失败，JWT生成失败")
	TooManyRequests           = registryError(10000007, "请求过多")
	MethodNotAllowed          = registryError(10000008, "请求方法错误")
	PermissionDenied          = registryError(20000000, "用户权限不足")
)
