package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gin-gonic/gin"
	"github.com/satyanurhutama/elastic-search-crud-golang/constant"
	"github.com/satyanurhutama/elastic-search-crud-golang/dto"
	"github.com/satyanurhutama/elastic-search-crud-golang/library"
	"github.com/satyanurhutama/elastic-search-crud-golang/utils"
)

func CreateStudent(c *gin.Context) {
	result := &dto.GeneralResponse{}

	var student dto.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		result.Error = true
		result.Message = err.Error()
		result.Code = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, result)
		return
	}

	body, err := json.Marshal(student)
	if err != nil {
		result.Error = true
		result.Message = "Error encoding student data"
		result.Code = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	req := esapi.IndexRequest{
		Index:      constant.ESIndexStudent,
		DocumentID: student.ID,
		Body:       bytes.NewReader(body),
		Refresh:    "true",
	}

	client := library.GetESClient()

	res, err := req.Do(c, client)
	if err != nil || res.IsError() {
		result.Error = true
		result.Message = "Error creating student"
		result.Code = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	defer res.Body.Close()

	result.Error = false
	result.Code = http.StatusOK
	result.Message = "Student created successfully"
	result.Data = student

	c.JSON(http.StatusCreated, result)
}

func GetStudent(c *gin.Context) {
	result := &dto.GeneralResponse{}

	id := c.Param("id")
	req := esapi.GetRequest{
		Index:      constant.ESIndexStudent,
		DocumentID: id,
	}

	client := library.GetESClient()

	res, err := req.Do(c, client)
	if err != nil || res.IsError() {
		result.Error = true
		if res != nil && res.StatusCode == http.StatusNotFound {
			result.Message = "Student not found"
			result.Code = http.StatusNotFound
			c.JSON(http.StatusNotFound, result)
			return
		}

		result.Message = "Error getting student"
		result.Code = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	defer res.Body.Close()

	student, err := utils.ParseESResponse[dto.Student](res.Body)
	if err != nil {
		result.Error = true
		result.Message = "Error parsing response"
		result.Code = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	result.Error = false
	result.Code = http.StatusOK
	result.Message = http.StatusText(http.StatusOK)
	result.Data = student

	c.JSON(http.StatusOK, result)
}

func UpdateStudent(c *gin.Context) {
	result := &dto.GeneralResponse{}

	var student dto.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		result.Error = true
		result.Message = err.Error()
		result.Code = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, result)
		return
	}

	body, err := json.Marshal(map[string]interface{}{
		"doc": student, // Include the student object as the "doc" field
	})
	if err != nil {
		result.Error = true
		result.Message = "Error encoding student data"
		result.Code = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	req := esapi.UpdateRequest{
		Index:      constant.ESIndexStudent,
		DocumentID: student.ID,
		Body:       bytes.NewReader(body),
		Refresh:    "true",
	}

	client := library.GetESClient()

	res, err := req.Do(c, client)
	if err != nil || res.IsError() {
		result.Error = true
		if res != nil && res.StatusCode == http.StatusNotFound {
			result.Message = "Student not found"
			result.Code = http.StatusNotFound
			c.JSON(http.StatusNotFound, result)
			return
		}

		result.Message = "Error updating student"
		result.Code = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	defer res.Body.Close()

	result.Error = false
	result.Code = http.StatusOK
	result.Message = "Student updated successfully"
	result.Data = student

	c.JSON(http.StatusOK, result)
}

func DeleteStudent(c *gin.Context) {
	result := &dto.GeneralResponse{}

	id := c.Param("id")
	req := esapi.DeleteRequest{
		Index:      constant.ESIndexStudent,
		DocumentID: id,
		Refresh:    "true",
	}

	client := library.GetESClient()

	res, err := req.Do(c, client)
	if err != nil || res.IsError() {
		result.Error = true
		if res != nil && res.StatusCode == http.StatusNotFound {
			result.Message = "Student not found"
			result.Code = http.StatusNotFound
			c.JSON(http.StatusNotFound, result)
			return
		}

		result.Message = "Error deleting student"
		result.Code = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	defer res.Body.Close()

	result.Error = false
	result.Code = http.StatusOK
	result.Message = "Student deleted successfully"

	c.JSON(http.StatusOK, result)
}
