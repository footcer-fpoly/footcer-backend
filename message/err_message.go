package message

import "errors"

var (
	UserConflict   = errors.New("Người dùng đã tồn tại")
	UserNotFound   = errors.New("Người dùng không tồn tại")
	UserIsAdmin   = errors.New("Bạn đã chủ sân, không thể tạo người dùng ")
	UserNotUpdated = errors.New("Cập nhật thông tin người dùng thất bại")
	EmailExits     = errors.New("Email đã tồn tại")
	SignUpFail     = errors.New("Đăng ký thất bại")
	SomeWentWrong  = errors.New("Có lỗi xảy ra")

	//stadium
	StadiumNotFound   = errors.New("Sân bóng không tồn tại")


)
