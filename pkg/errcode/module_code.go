package errcode

var (
	ErrorGetTagListFail       = NewError(20010001, "获取标签列表失败")
	ErrorCreateTagFail        = NewError(20010002, "创建标签失败")
	ErrorUpdateTagFail        = NewError(20010003, "更新标签失败")
	ErrorDeleteTagFail        = NewError(20010004, "删除标签失败")
	ErrorCountTagFail         = NewError(20010005, "统计标签失败")
	UnauthorizedAuthNotExist  = NewError(20010006, "权限验证失败")
	UnauthorizedTokenGenerate = NewError(20010007, "生成token失败")
	UnauthorizedTokenTimeout  = NewError(20010008, "token已过期")
	UnauthorizedTokenError    = NewError(20010009, "token错误")
)
