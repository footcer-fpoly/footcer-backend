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

type GameHandler struct {
	GameRepo   repository.GameRepository
	UserRepo   repository.UserRepository
	NotifyRepo repository.NotificationRepository
}

func (g *GameHandler) AddGame(c echo.Context) error {
	req := model.Game{}

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	req.GameId = uuid.NewV1().String()
	req.Finish = "0"
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()

	game, err := g.GameRepo.AddGame(c.Request().Context(), req)
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
		Data:       game,
	})

}

func (g *GameHandler) UpdateGame(c echo.Context) error {
	req := model.Game{}

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	_, err := g.GameRepo.UpdateGame(c.Request().Context(), req)
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

func (g *GameHandler) DeleteGame(c echo.Context) error {
	req := model.Game{}

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	err := g.GameRepo.DeleteGame(c.Request().Context(), c.Param("id"))
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

func (g *GameHandler) JoinGame(c echo.Context) error {
	req := model.GameTemp{}

	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	req.GameTempId = uuid.NewV1().String()
	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	err := g.GameRepo.JoinGame(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	token, errToken := g.UserRepo.GetToken(c.Request().Context(), req.UserNotifyId)
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
		Type: "JOIN_GAME",
		Body: model.BodyNotification{
			Title:     "Tham gia trận đấu ",
			Content:   req.NameInvite + " gửi lời thách đấu đến trận đấu của bạn",
			GeneralId: req.GameId,
		},
	}, tokens,
	)
	_, err = g.NotifyRepo.AddNotification(c.Request().Context(), model.Notification{
		NotifyID:  uuid.NewV1().String(),
		Key:       "JOIN_GAME",
		Title:     "Tham gia trận đấu ",
		Content:   req.NameInvite + " gửi lời thách đấu đến trận đấu của bạn",
		Icon:      "",
		GeneralID: req.GameId,
		UserId:    req.UserNotifyId,
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

	token, errToken = g.UserRepo.GetToken(c.Request().Context(), claims.UserId)
	if errToken != nil {
		log.Error(errToken)
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    errToken.Error(),
			Data:       nil,
		})
	}

	tokens = []string{}
	tokens = append(tokens, token)
	service.PushNotification(c, model.DataNotification{
		Type: "JOIN_GAME",
		Body: model.BodyNotification{
			Title:     "Tham gia trận đấu ",
			Content:   "Bạn vừa gửi lời mời thách đấu đến đội bóng " + req.NameHost,
			GeneralId: req.GameId,
		},
	}, tokens,
	)
	_, err = g.NotifyRepo.AddNotification(c.Request().Context(), model.Notification{
		NotifyID:  uuid.NewV1().String(),
		Key:       "JOIN_GAME",
		Title:     "Tham gia trận đấu ",
		Content:   "Bạn vừa gửi lời mời thách đấu đến đội bóng " + req.NameHost,
		Icon:      "",
		GeneralID: req.GameId,
		UserId:    claims.UserId,
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

func (g *GameHandler) AcceptJoin(c echo.Context) error {
	req := model.GameTemp{}

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	err := g.GameRepo.AcceptJoin(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	users, errToken := g.UserRepo.GetTokenForTeam(c.Request().Context(), req.TeamId) // id guest
	if errToken != nil {
		log.Error(errToken)
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    errToken.Error(),
			Data:       nil,
		})
	}
	var tokens []string
	for _, user := range users {
		tokens = append(tokens, user.TokenNotify)
		_, err = g.NotifyRepo.AddNotification(c.Request().Context(), model.Notification{
			NotifyID:  uuid.NewV1().String(),
			Key:       "ACCEPT_GAME",
			Title:     "Chấp nhận thách đấu ",
			Content:   "Đội bóng " + req.NameHost + " đã chấp nhận lời thách đấu đội bóng " + req.NameInvite + " của bạn vào ngày " + req.DateGame,
			Icon:      "",
			GeneralID: req.GameId,
			UserId:    user.UserId,
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
	}
	service.PushNotification(c, model.DataNotification{
		Type: "ACCEPT_GAME",
		Body: model.BodyNotification{
			Title:     "Chấp nhận thách đấu ",
			Content:   "Đội bóng " + req.NameHost + " đã chấp nhận lời thách đấu đội bóng " + req.NameInvite + " của bạn vào ngày " + req.DateGame,
			GeneralId: req.GameId,
		},
	}, tokens,
	)

	tokens = []string{}
	users, errToken = g.UserRepo.GetTokenForTeam(c.Request().Context(), req.UserNotifyId) // id host
	if errToken != nil {
		log.Error(errToken)
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    errToken.Error(),
			Data:       nil,
		})
	}

	for _, user := range users {
		tokens = append(tokens, user.TokenNotify)
		_, err = g.NotifyRepo.AddNotification(c.Request().Context(), model.Notification{
			NotifyID:  uuid.NewV1().String(),
			Key:       "ACCEPT_GAME",
			Title:     "Trận đấu",
			Content:   "Đội bóng " + req.NameHost + " của bạn đã xác nhận đội bóng " + req.NameInvite + " là đối thủ cho trận đấu vào ngày " + req.DateGame,
			Icon:      "",
			GeneralID: req.GameId,
			UserId:    user.UserId,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	service.PushNotification(c, model.DataNotification{
		Type: "ACCEPT_GAME",
		Body: model.BodyNotification{
			Title:     "Trận đấu",
			Content:   "Đội bóng " + req.NameHost + " của bạn đã xác nhận đội bóng " + req.NameInvite + " là đối thủ cho trận đấu vào ngày " + req.DateGame,
			GeneralId: req.GameId,
		},
	}, tokens,
	)

	// tu choi game


	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       nil,
	})

}

func (g *GameHandler) RefuseJoin(c echo.Context) error {
	req := model.GameTemp{}

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	err := g.GameRepo.RefuseJoin(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	token, errToken := g.UserRepo.GetToken(c.Request().Context(), req.UserNotifyId)
	if errToken != nil {
		log.Error(errToken)
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    errToken.Error() + " tokennn",
			Data:       nil,
		})
	}

	var tokens []string
	tokens = append(tokens, token)
	service.PushNotification(c, model.DataNotification{
		Type: "REFUSE_GAME",
		Body: model.BodyNotification{
			Title:     "Từ chối thách đấu",
			Content:   req.NameHost + " đã từ chối lời thách đấu " + req.NameInvite + " của bạn",
			GeneralId: req.GameId,
		},
	}, tokens,
	)
	_, err = g.NotifyRepo.AddNotification(c.Request().Context(), model.Notification{
		NotifyID:  uuid.NewV1().String(),
		Key:       "REFUSE_GAME",
		Title:     "Từ chối thách đấu",
		Content:   req.NameHost + " đã từ chối lời thách đấu " + req.NameInvite + " của bạn",
		Icon:      "",
		GeneralID: req.GameId,
		UserId:    req.UserNotifyId,
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

func (g *GameHandler) GetGames(c echo.Context) error {
	date := c.Param("date")

	games, err := g.GameRepo.GetGames(c.Request().Context(), date)
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
		Data:       games,
	})
}

func (g *GameHandler) GetGame(c echo.Context) error {
	gameId := c.Param("id")
	game, err := g.GameRepo.GetGame(c.Request().Context(), gameId)
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
		Data:       game,
	})
}

func (g *GameHandler) GetGameForUser(c echo.Context) error {

	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	game, err := g.GameRepo.GetGameForUser(c.Request().Context(), claims.UserId)
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
		Data:       game,
	})
}

func (g *GameHandler) UpdateScore(c echo.Context) error {
	req := model.Game{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	_, err := g.GameRepo.UpdateScore(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       nil,
	})
}
