package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"footcer-backend/model"
	"footcer-backend/security/pro"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
)

func PushNotification(c echo.Context, data model.DataNotification, tokens []string) {
	payload := model.PayLoad{
		RegistrationIds: tokens,
		Data:            data,
	}
	payloadByte, _ := json.Marshal(payload)
	go Push(payloadByte)

}

func Push(payload []byte) {
	req, err := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "key="+pro.FIREBASE_SERVER)
	client := &http.Client{}
	res, _ := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
	}

	bytes, _ := ioutil.ReadAll(res.Body)
	resString := string(bytes)

	fmt.Println(res.StatusCode)
	fmt.Print(resString)

	defer res.Body.Close()
}
