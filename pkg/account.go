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
	request, err := http.NewRequest(http.MethodGet, "https://online.moysklad.ru/api/remap/1.2/entity/employee", http.NoBody)

	if err != nil {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	err = godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	request.SetBasicAuth(os.Getenv("authLogin"), os.Getenv("authPassword"))
	defer request.Body.Close()

	response, err := client.Do(request)

	if err != nil {
		logrus.Fatal(err)
		return
	}

	resBody := io.Reader(response.Body)

	// var responseObject interface{} // как в Документаций

	var responseObject struct {
		data map[string][]map[string]string
	}

	json.NewDecoder(resBody).Decode(&responseObject.data)
	c.JSON(200, responseObject.data)
}



func deleteAccount(c *gin.Context) {
	
	client := http.Client{Timeout: 5 * time.Second}
	id := c.Param("id")

	request, err := http.NewRequest(http.MethodDelete, "https://online.moysklad.ru/api/remap/1.2/entity/employee/"+ id, http.NoBody)

	if err != nil{
		c.Status(http.StatusServiceUnavailable)
		return
	}
	err = godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	request.SetBasicAuth(os.Getenv("authLogin"), os.Getenv("authPassword"))

	response, err := client.Do(request)

	if err != nil {
		logrus.Fatal(err)
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

	request, err := http.NewRequest(http.MethodPost, "https://online.moysklad.ru/api/remap/1.2/entity/employee", bytes.NewBuffer(postBody))
	if err != nil {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	err = godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	request.SetBasicAuth(os.Getenv("authLogin"), os.Getenv("authPassword"))
	defer request.Body.Close()
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		logrus.Fatal(err)
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
	postBody, err := json.Marshal(account)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	request, err := http.NewRequest(http.MethodPut, "https://online.moysklad.ru/api/remap/1.2/entity/employee/" + id, bytes.NewBuffer(postBody))

	if err != nil {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	err = godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	request.SetBasicAuth(os.Getenv("authLogin"), os.Getenv("authPassword"))
	
	defer request.Body.Close()
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		logrus.Fatal(err)
		return
	}
	
	defer response.Body.Close()

	c.JSON(200, gin.H{"message":"Successfully updated"})
}