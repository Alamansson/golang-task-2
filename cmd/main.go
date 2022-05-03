package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	// golangtask "github.com/alamansson/golang-task-2"
	// golangtask2 "github.com/alamansson/golang-task-2"
	"github.com/gin-gonic/gin"
)





func main() {
	router := gin.Default()
	router.GET("/accounts", func(c *gin.Context) {
		client := http.Client{Timeout: 5 * time.Second}
		req, err := http.NewRequest(http.MethodGet, "https://online.moysklad.ru/api/remap/1.2/entity/employee", http.NoBody)
		if err != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		req.SetBasicAuth("admin@alamanovdev", "02eaaf82505f")
		defer req.Body.Close()
		res, err := client.Do(req)

		resBody := io.Reader(res.Body)

		if err != nil {
			log.Fatal(err)
		}

		var v struct {
			data map[string][]map[string]string
		}

		json.NewDecoder(resBody).Decode(&v.data)
		c.JSON(200, v.data)
	})

	router.POST("/account", func(c *gin.Context) {
		var emp map[string]interface{}
		// c.BindJSON(&emp)
		c.Bind(&emp)
		fmt.Printf("Bind: %T, %v   ",emp, emp)

		// emp["attributes"] = golangtask.Meta

		fmt.Printf("Bind: %T, %v   ",emp, emp)


		// x := c.Request
		// fmt.Printf("\n\n\n c.Request.Body: %T, %v   ",x, x)

		// c.IndentedJSON(http.StatusCreated, emp)
		// fmt.Printf("\n\n\nIndentedJSON: %T, %v   ",emp, emp)


		postBody, _ := json.Marshal(emp)
		fmt.Printf("\n\n\nMarshal:  %T, %v   ",postBody, postBody)

		// var jsonData = []byte(postBody)
		client := http.Client{Timeout: 5 * time.Second}

		object := bytes.NewBuffer(postBody)


				

		req, err := http.NewRequest(http.MethodPost, "https://online.moysklad.ru/api/remap/1.2/entity/employee/", object)
		if err != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		req.SetBasicAuth("admin@alamanovdev", "02eaaf82505f")
		defer req.Body.Close()


		if err != nil {
			log.Fatalln(err)
		}
		resp, err := client.Do(req)
		defer resp.Body.Close()


		c.JSON(200, emp)
		c.JSON(200, gin.H{"message":"Successfully created"})

	})

	router.DELETE("/delete/:id", func(c *gin.Context) {
		
		client := http.Client{Timeout: 5 * time.Second}
		id := c.Param("id")

		req, err := http.NewRequest(http.MethodDelete, "https://online.moysklad.ru/api/remap/1.2/entity/employee/"+ id, http.NoBody)
		if err != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		req.SetBasicAuth("admin@alamanovdev", "02eaaf82505f")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		c.JSON(200, gin.H{"message":"Successfully deleted"})
	})


	router.PUT("/update/:id", func(c *gin.Context) {
		// client := http.Client{Timeout: 5 * time.Second}
		// id := c.Param("id")
		
		// req, err := http.NewRequest(http.MethodPut, "https://online.moysklad.ru/api/remap/1.2/entity/employee/"+id, http.NoBody)
		// if err != nil {
		// 	c.Status(http.StatusServiceUnavailable)
		// 	return
		// }

		// req.SetBasicAuth("admin@alamanovdev", "02eaaf82505f")
		
		// resp, err := client.Do(req)

		// if err != nil {
		// 	log.Fatalln(err)
		// }

	})
	router.Run(":8000")
}
