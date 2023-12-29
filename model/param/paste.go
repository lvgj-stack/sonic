package param

import (
	"time"
)

type Paste struct {
	Key       string    `json:"key" form:"key"`                                                // 主键:索引
	Lang      string    `json:"lang" example:"plain" form:"lang" binding:"gte=1"`              // 语言类型
	Content   string    `json:"content" example:"Hello World!" form:"content" binding:"gte=1"` // 内容，最大长度为 16777215(2^24-1) 个字符
	Password  string    `json:"password" example:"" from:"password"`                           // 密码
	ClientIP  string    `json:"client_ip" from:"client_ip"`                                    // 用户 IP
	Username  string    `json:"username" from:"username"`                                      // 用户名
	CreatedAt time.Time `json:"-"`                                                             // 存储记录的创建时间
}
