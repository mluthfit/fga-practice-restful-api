package controllers

import (
	"fga-practice-rest-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetPerson(c *gin.Context) {
	var person structs.Person
	var result gin.H

	var id = c.Param("id")
	var err = idb.DB.Where("id = ?", id).First(&person).Error

	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetPersons(c *gin.Context) {
	var persons []structs.Person
	var result gin.H

	idb.DB.Find(&persons)

	if len(persons) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": persons,
			"count":  len(persons),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreatePerson(c *gin.Context) {
	var person structs.Person
	var result gin.H

	var firstName = c.PostForm("first_name")
	var lastName = c.PostForm("last_name")

	person.FirstName = firstName
	person.LastName = lastName
	idb.DB.Create(&person)

	result = gin.H{
		"result": person,
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdatePerson(c *gin.Context) {
	var id = c.Query("id")
	var firstName = c.PostForm("first_name")
	var lastName = c.PostForm("last_name")

	var person, newPerson structs.Person
	var result gin.H

	var err = idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	newPerson.FirstName = firstName
	newPerson.LastName = lastName

	err = idb.DB.Model(&person).Updates(newPerson).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "successfully updated data",
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeletePerson(c *gin.Context) {
	var person structs.Person
	var result gin.H

	var id = c.Param("id")
	var err = idb.DB.First(&person, id).Error

	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	err = idb.DB.Delete(&person).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "data deleted succesfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
