package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/go-playground/validator.v8"
	"os"
	"os/signal"
	"time"
	. "topic.jtthink.com/src"
)

func main3(){
	c := make(chan os.Signal)
	//监听所有信号
	signal.Notify(c)
	//可以监听指定型号
	//signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)

	//阻塞直到有信号传入
	fmt.Println("启动")
	s := <-c
	fmt.Println("退出信号", s)
}

func main()  {
	count:=0
	go func() {
		for {
			fmt.Println("执行",count)
			count++
			time.Sleep(time.Second*1)
		}
	}()

	c:=make(chan os.Signal)

	go func() {
		ctx,_:=context.WithTimeout(context.Background(),time.Second*5)
		select {
			case <-ctx.Done():
				c<-os.Interrupt
		}
	}()

	signal.Notify(c)
	s:=<-c
	fmt.Println(s)

}

func main2()  {
	router:=gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("topicurl", TopicUrl)
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
