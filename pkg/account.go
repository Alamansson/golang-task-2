package pkg

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)


func getAccount(c *gin.Context) {
	client := http.Client{Timeout: 5 * time.Second}
	request, err := BasicAuth(http.MethodGet, "", nil)
	if err != nil {
		logrus.Errorf("Error during request", err)
		return 
	}
	response, err := client.Do(request)
	if err != nil {
		logrus.Errorf("Error during request", err)
		return 
	}

	resBody := io.Reader(response.Body)
	
	var responseObject interface{} // Выдача данных как в документаций мойСлад
	// var responseObject map[string][]map[string]string // только данные аккаунтов
	err = json.NewDecoder(resBody).Decode(&responseObject)
	if err != nil {
		logrus.Errorf("Cannot unmarshal object: ", err)
		return 
	}
	c.JSON(200, responseObject)
	return 
}


func deleteAccount(c *gin.Context) {
	client := http.Client{Timeout: 5 * time.Second}
	id := c.Param("id")

	request, err := BasicAuth(http.MethodDelete, id, nil)
	if err != nil {
		logrus.Errorf("Error during request: ", err)
		return 
	}
	response, err := client.Do(request)
	if err != nil {
		logrus.Errorf("Error during request: ", err)
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
		logrus.Errorf("Error return: ", err)
		return 
	}

	request, err := BasicAuth(http.MethodPost, "", postBody)
	if err != nil {
		logrus.Errorf("Error during request", err)
		return 
	}

	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		logrus.Errorf("Error during request", err)
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
		logrus.Errorf("Error return: ", err)
		return 
	}

	request, err := BasicAuth(http.MethodPut, id, updateBody)
	if err != nil {
		logrus.Errorf("Error during request", err)
		return 
	}

	request.Body.Close()
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		logrus.Errorf("Error return: ", err)
		return 
	}
	
	response.Body.Close()
	c.JSON(200, gin.H{"message":"Successfully updated"})
}


func BasicAuth(method, id string, body []uint8) (*http.Request, error) {
	bodyPost := bytes.NewBuffer(body)
	request, err := http.NewRequest(method, "https://online.moysklad.ru/api/remap/1.2/entity/employee/" + id, bodyPost)
	if err != nil {
		logrus.Errorf("Error return: ", err)
		return nil, http.ErrBodyNotAllowed
	}
	err = godotenv.Load()
	if err != nil {
		logrus.Errorf("Error loading .env file ", err)
		return nil, http.ErrNotSupported
	}
	request.SetBasicAuth(os.Getenv("authLogin"), os.Getenv("authPassword"))
	defer request.Body.Close()
	return request, nil
}


