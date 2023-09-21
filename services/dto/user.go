package dto

type Hello struct {
	Name string `form:"name" comment:"名称" binding:"required"`
}
