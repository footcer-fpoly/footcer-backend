package message

import "errors"

var (
	//user
	UserConflict   = errors.New("Người dùng đã tồn tại")
	UserNotFound   = errors.New("Người dùng không tồn tại")
	UserIsAdmin    = errors.New("Bạn đã chủ sân, không thể tạo người dùng ")
	UserNotUpdated = errors.New("Cập nhật thông tin người dùng thất bại")
	EmailExits     = errors.New("Email đã tồn tại")
	SignUpFail     = errors.New("Đăng ký thất bại")

	//stadium
	StadiumNotFound   = errors.New("Sân bóng không tồn tại")
	StadiumNotUpdated = errors.New("Cập nhật thông tin sân thất bại")

	//team
	TeamMemberExits  = errors.New("Thành viên đã tồn tại")
	TeamMemberDelete = errors.New("Xoá thành viên thành công")
	TeamDelete       = errors.New("Xoá nhóm thành công")
	AdminIsTeam      = errors.New("Bạn là trưởng nhóm, không thể rời nhóm")

	//
	Permission    = errors.New("Lỗi quyền truy cập")
	SomeWentWrong = errors.New("Có lỗi xảy ra")
)
