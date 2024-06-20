package controllers

import (
	"devops-go/services"
	"devops-go/services/impl"
)

func (router RouterGroup) TestController() {
	testServiceImpl := impl.TestServiceImpl{}
	var testService services.TestService = testServiceImpl

	authRouters := router.Group("/test")
	{

		authRouters.POST("/trace", testService.Trace)

	}
}
