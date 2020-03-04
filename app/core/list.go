package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)


func List(c *gin.Context) {
	var itemList string
	filepath.Walk(os.Getenv("WORK_DIR"), func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		itemList += fmt.Sprintf(`<li><a href="/item/%s">%s</a></li>`,info.Name(),info.Name())
		return nil
	})
	htmlContext := strings.Replace(LIST_VIEW,"{{context}}",itemList,-1)
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200,htmlContext)
}

func Item(c *gin.Context) {
	name := c.Param("name")
	fullPath := fmt.Sprintf("%s/%s",os.Getenv("WORK_DIR"),name)
	f,err := os.Open(fullPath)
	if err != nil {
		c.String(200,err.Error())
		return
	}
	defer f.Close()
	data,err := ioutil.ReadAll(f)
	if err != nil {
		c.String(200,err.Error())
		return
	}
	htmlContext := strings.Replace(ITEM_VIEW,"{{item}}",name,-1)
	htmlContext = strings.Replace(htmlContext,"{{context}}",string(data),-1)
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200,htmlContext)
}

func Edit(c *gin.Context) {
	name := c.Param("name")
	context := c.PostForm("context")
	fullPath := fmt.Sprintf("%s/%s",os.Getenv("WORK_DIR"),name)
	err := ioutil.WriteFile(fullPath,[]byte(context),644)
	if err != nil {
		c.String(200,err.Error())
		return
	}
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200,"<h1>保存成功</h1>")
}

func Add(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200,ADD_VIEW)
}

func ItemAdd(c *gin.Context) {
	name := c.PostForm("name")
	context := c.PostForm("context")
	fullPath := fmt.Sprintf("%s/%s",os.Getenv("WORK_DIR"),name)
	err := ioutil.WriteFile(fullPath,[]byte(context),644)
	if err != nil {
		c.String(200,err.Error())
		return
	}
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200,"<h1>保存成功</h1>")
}


// 基础认证拦截器
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if len(auth) < 1 {
			c.String(401,"Unauthozied")
			return
		}
		authList := strings.Split(auth, " ")
		log.Println(auth)
		if authList[0] != "Basic" {
			c.Abort()
			c.String(401,"Unauthozied")
			return
		}

		// 处理请求
		c.Next()
	}
}