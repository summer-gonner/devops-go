package controllers

import (
	"devops-go/basicdata/services"
	"devops-go/basicdata/services/impl"
)

func (router RouterGroup) TestController() {
	testServiceImpl := impl.TestServiceImpl{}
	var testService services.TestService = testServiceImpl

	authRouters := router.Group("/test")
	{

		authRouters.POST("/trace", testService.Trace)

	}
}
