package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	. "topic.jtthink.com/src"
)


func main(){

	router := gin.Default()
	/*
	router.GET("/topic_id/:topic_id", func(c *gin.Context) {
		c.String(200,"获取topicid=%s的帖子",c.Param("topic_id"))
	})
	router.GET("v1/topics", func(c *gin.Context) {
		if c.Query("username") == ""{
			c.String(200,"获取帖子列表")
		}else{
			c.String(200,"获取topicid=%s的帖子",c.Param("topic_id"))
		}
	})
	*/

	/*
	v1 := router.Group("v1/topics")
	{
		v1.GET("", func(c *gin.Context) {
			if c.Query("username")==""{
				c.String(200,"获取帖子列表")
			}else {
				c.String(200,"获取用户名=%s的帖子列表",c.Query("username"))
			}
		})
		v1.GET("/:topic_id", func(c *gin.Context) {
			c.String(200,"获取topicid=%s的帖子",c.Param("topic_id"))
		})
	}
	*/

	if v,ok := binding.Validator.Engine().(*validator.Validate);ok{
		v.RegisterValidation("topicurl",TopicUrl)
		v.RegisterValidation("topics", TopicsValidate) //验证长度
	}

	v1:=router.Group("/v1/topics")//单条帖子
	{
		v1.GET("", GetTopicList)

		v1.GET("/:topic_id",GetTopicDetail)

		v1.Use(MustLogin())
		{
			v1.POST("",NewTopic)
			v1.DELETE("/:topic_id",DelTopic)
		}


	}

	v2:=router.Group("/v1/mtopics")//多条帖子
	{
		v2.Use(MustLogin())
		{
			v2.POST("",NewTopics)

		}
	}





	router.Run()
}
