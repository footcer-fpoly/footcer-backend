package message

import "errors"

var (
	UserConflict   = errors.New("Người dùng đã tồn tại")
	UserNotFound   = errors.New("Người dùng không tồn tại")
	UserIsClient   = errors.New("Bạn đã là người dùng, không thể tạo chủ sân ")
	UserNotUpdated = errors.New("Cập nhật thông tin người dùng thất bại")
	EmailExits     = errors.New("Email đã tồn tại")
	SignUpFail     = errors.New("Đăng ký thất bại")
	SomeWentWrong  = errors.New("Có lỗi xảy ra")
)
