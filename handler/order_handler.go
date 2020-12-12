package handler

import (
	"footcer-backend/helper"
	"footcer-backend/log"
	"footcer-backend/model"
	"footcer-backend/repository"
	"footcer-backend/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

type OrderHandler struct {
	OrderRepo  repository.OrderRepository
	UserRepo   repository.UserRepository
	NotifyRepo repository.NotificationRepository
}

func (o *OrderHandler) AddOrder(c echo.Context) error {
	req := model.Order{}

	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	req.OrderId = uuid.NewV1().String()
	req.UserId = claims.UserId
	req.Finish = false
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()

	teamDetails, err := o.OrderRepo.AddOrder(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	token, errToken := o.UserRepo.GetToken(c.Request().Context(), req.StadiumUserId)
	if errToken != nil {
		log.Error(errToken)
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    errToken.Error(),
			Data:       nil,
		})
	}

	var tokens []string
	tokens = append(tokens, token)
	service.PushNotification(c, model.DataNotification{
		Type: "ADD_ORDER",
		Body: model.BodyNotification{
			Title:     "Yêu cầu đặt sân",
			Content:   claims.UserName + " đã yêu cầu đặt sân của bạn vào lúc " + time.Now().String(),
			GeneralId: req.OrderId,
		},
	}, tokens,
	)
	_, err = o.NotifyRepo.AddNotification(c.Request().Context(), model.Notification{
		NotifyID:  uuid.NewV1().String(),
		Key:       "DELETE_MEMBER",
		Title:     "Mời rời đội bóng",
		Content:   claims.UserName + " đã yêu cầu đặt sân của bạn vào lúc " + time.Now().String(),
		Icon:      "",
		GeneralID: req.OrderId,
		UserId:    req.StadiumUserId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       teamDetails,
	})

}

func (o *OrderHandler) UpdateStatusOrder(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	req := model.OrderStatus{}
	if claims.Role == 0 {
		req.IsUser = true
	} else {
		req.IsUser = false
	}

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	err := o.OrderRepo.UpdateStatusOrder(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	token, errToken := o.UserRepo.GetToken(c.Request().Context(), req.UserId)
	if errToken != nil {
		log.Error(errToken)
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    errToken.Error(),
			Data:       nil,
		})
	}
	title := ""
	content := ""
	if req.Status == "ACCEPT" {
		title = "Chấp nhận đặt sân"
		content = "Sân " + req.Name + " đã chấp nhận yêu cầu đặt sân của bạn"
	}else if req.Status =="REJECT"{
		if req.IsUser {
			//noti den user
			title = "Từ chối đặt sân"
			content = "Sân " + req.Name + " đã từ chối yêu cầu đặt sân của bạn"
		}else{
			//noti den chu san
			title = "Huỷ đặt sân"
			content =   req.Name + " đã huỷ yêu cầu đặt sân của bạn"
		}

	}else if req.Status == "FINISH"{
		//noti user
		title = "Hoàn thành đặt sân"
		content =   "Hãy đánh giá sân " + req.Name
	}

	var tokens []string
	tokens = append(tokens, token)
	service.PushNotification(c, model.DataNotification{
		Type: req.Status,
		Body: model.BodyNotification{
			Title:     title,
			Content:   content,
			GeneralId: req.OrderId,
		},
	}, tokens,
	)
	_, err = o.NotifyRepo.AddNotification(c.Request().Context(), model.Notification{
		NotifyID:  uuid.NewV1().String(),
		Key:       req.Status,
		Title:     title,
		Content:   content,
		Icon:      "",
		GeneralID: req.OrderId,
		UserId:    req.UserId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}



	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       nil,
	})

}

func (o *OrderHandler) FinishOrder(c echo.Context) error {
	req := model.Order{}

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}
	req.Finish = true

	err := o.OrderRepo.FinishOrder(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       nil,
	})

}

func (o *OrderHandler) ListOrderForStadium(c echo.Context) error {
	stadiumID := c.Param("id")

	orders, err := o.OrderRepo.ListOrderForStadium(c.Request().Context(), stadiumID)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       err.Error,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       orders,
	})
}

func (o *OrderHandler) ListOrderForUser(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	orders, err := o.OrderRepo.ListOrderForUser(c.Request().Context(), claims.UserId)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       err.Error,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       orders,
	})
}

func (o *OrderHandler) OrderDetail(c echo.Context) error {
	orderId := c.Param("id")

	orders, err := o.OrderRepo.OrderDetail(c.Request().Context(), orderId)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       err.Error,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       orders,
	})
}
