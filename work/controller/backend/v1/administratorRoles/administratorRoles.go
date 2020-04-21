package administratorRoles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-edu/work/common"
	"go-edu/work/services"
	"net/http"
	"strconv"
)


func Index(c *gin.Context) {
	pageTmp := c.DefaultQuery("page", "1")
	pageSizeTmp := c.DefaultQuery("pageSize", "20")
	page,_ := strconv.Atoi(pageTmp)
	pageSize,_ := strconv.Atoi(pageSizeTmp)
	obj := &services.IndexRolesService{
		Page: page,
		PageSize: pageSize,
	}
	result := obj.Index()
	c.JSON(http.StatusOK, result)
}

// 添加角色
func Create(c *gin.Context) {
	service := &services.CreateRolesService{}
	if err := c.ShouldBind(service); err != nil {
		fmt.Printf("params err:%v\n", err)
		c.JSON(http.StatusOK, common.ValidateResponse(err))
		return
	}
	res := service.Create()
	c.JSON(http.StatusOK, res)
}



func UpdateStatus(c *gin.Context) {
	status := &services.StatusRolesService{}
	if err := c.ShouldBindJSON(status); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, common.ValidateResponse(err))
		return
	}
	res := status.UpdateStatus()
	c.JSON(http.StatusOK, res)
}

func Edit(c *gin.Context)  {
	_id := c.Param("id")
	id,_ := strconv.Atoi(_id)
	service := services.EditRolesService{Id: id}
	resp := service.Edit()
	 c.JSON(http.StatusOK, resp)
}


func Update(c *gin.Context) {
	service := &services.UpdateRolesService{}
	if err := c.ShouldBindJSON(service);err !=nil {
		c.JSON(http.StatusOK, common.ValidateResponse(err))
	}
	resp := service.Update()
	c.JSON(http.StatusOK, resp)
}
