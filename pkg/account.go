package pkg

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)


func getAccount(c *gin.Context) {
	client := http.Client{Timeout: 5 * time.Second}
	request := BasicAuth(http.MethodGet, "", nil)
	response, err := client.Do(request)
	if err != nil {
		logrus.Errorf("Error during request", err)
		return
	}

	resBody := io.Reader(response.Body)
	
	// var responseObject interface{} // Выдача данных как в документаций мойСлад
	var responseObject map[string][]map[string]string // только данные аккаунтов
	error := json.NewDecoder(resBody).Decode(&responseObject)
	if error != nil {
		logrus.Errorf("Cannot unmarshal object",error)
	}
	c.JSON(200, responseObject)
}


func deleteAccount(c *gin.Context) {
	client := http.Client{Timeout: 5 * time.Second}
	id := c.Param("id")

	req := BasicAuth(http.MethodDelete, id, nil)

	response, err := client.Do(req)
	if err != nil {
		logrus.Errorf("Error during request", err)
		return
	}
	defer response.Body.Close()
	c.JSON(200, gin.H{"message":"Successfully deleted"})
}


func createAccount(c *gin.Context) {
	client := http.Client{Timeout: 5 * time.Second}

	var account interface{}
	c.Bind(&account)

	postBody, err := json.Marshal(account)
	if err != nil {
		logrus.Errorf("Error: ", err)
		return
	}

	request := BasicAuth(http.MethodPost, "", postBody)
	request.Body.Close()

	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		logrus.Errorf("Error: ", err)
		return
	}
	defer response.Body.Close()
	
	c.JSON(200, gin.H{"message":"Successfully created"})
}

func updateAccount(c *gin.Context) {
	client := http.Client{Timeout: 5 * time.Second}
	id := c.Param("id")

	var account interface{}
	c.Bind(&account)

	updateBody, err := json.Marshal(account)
	if err != nil {
		logrus.Errorf("Error: ", err)
		return
	}

	request := BasicAuth(http.MethodPut, id, updateBody)
	request.Body.Close()
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		logrus.Errorf("Error: ", err)
		return
	}
	
	response.Body.Close()
	c.JSON(200, gin.H{"message":"Successfully updated"})
}


func BasicAuth(method, id string, body []uint8) *http.Request {
	bodyPost := bytes.NewBuffer(body)
	request, err := http.NewRequest(method, "https://online.moysklad.ru/api/remap/1.2/entity/employee/" + id, bodyPost)
	if err != nil {
		logrus.Errorf("Error: ", err)
	}
	err = godotenv.Load()
	if err != nil {
		logrus.Errorf("Error loading .env file ", err)
	}
	request.SetBasicAuth(os.Getenv("authLogin"), os.Getenv("authPassword"))
	defer request.Body.Close()
	return request
}


