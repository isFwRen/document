###  ginchat项目搭建

#### 环境搭建：

##### 1：引入GORM

1. ​	打开go依赖官网：[Go Packages - Go Packages](https://pkg.go.dev/)

2. **Install：在终端输入命令安装**

   ```
   go get -u gorm.io/gorm
   go get -u gorm.io/driver/mysql
   ```

3. 创建一个test目录，在目录下创建一个testGorm_test.go文件，内容如下：

   ```
   package test
   
   import (
   	"fmt"
   	"ginchat/models"
   	"testing"
   	"time"
   
   	"gorm.io/driver/mysql"
   	"gorm.io/gorm"
   )
   
   func TestUserBasic(t *testing.T) {
   	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=true&loc=Local"), &gorm.Config{})
   	if err != nil {
   		panic("failed to connect database")
   	}
   
   	// 迁移 schema,数据库没有这个表会自动创建表
   	db.AutoMigrate(&models.UserBasic{})
   
   	// Create
   	user := &models.UserBasic{}
   	user.Name = "小米"
   	user.LoginTime = time.Date(1990, 6, 15, 0, 0, 0, 0, time.UTC)
   	user.HeartbeatTime = time.Date(1990, 6, 15, 0, 0, 0, 0, time.UTC)
   	user.LoginOutTime = time.Date(1990, 6, 15, 0, 0, 0, 0, time.UTC)
   	db.Create(user)
   
   	// Read
   	fmt.Println(db.First(user, 1)) // 根据整型主键查找
   	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
   
   	// Update - 将 product 的 price 更新为 200
   	db.Model(user).Update("Password", "1234")
   	// Update - 更新多个字段
   	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
   	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
   
   	// Delete - 删除 product
   	//db.Delete(&product, 1)
   }
   
   ```

   

##### 2：引入gin框架

1. **Install：在终端输入以下命令**

   ```
   go get -u github.com/gin-gonic/gin
   ```

2. 在main包下，将gin引入到代码中：

   ```go
   终端输入：import "github.com/gin-gonic/gin"
   
   //项目目录下main.go
   package mai
   import (
   	"ginchat/router"
   )
   func main() {
   	r := router.Router()
   	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
   }
   
   //router包下app.go
   package router
   
   import (
   	"ginchat/service"
   
   	"github.com/gin-gonic/gin"
   )
   func Router() *gin.Engine {
   	r := gin.Default()
   	r.GET("/index", service.GetIndex)
   	return r
   }
   
   //service包下index.go
   package service
   
   import (
   	"github.com/gin-gonic/gin"
   )
   func GetIndex(c *gin.Context) {
   	c.JSON(200, gin.H{
   		"message": "welcome",
   	})
   }
   
   ```

   

#### 项目解决的bug

1. **出现的错误**：go在插入mysql数据时，出现`Error 1292 (22007): Incorrect datetime value: '0000-00-00' for column 'login_time' at row 1`
   **出现错误的原因：**这个错误时有Mysql 的严格模式(SQL Models)引起的，它要求日期时间段具有有效的日期时间值，而不允许`'0000-00-00'`这样的无效日期时间值
   **要解决这个错误，有以下方法：**

   - **使用有效的日期时间值：**最佳的解决方法是确保插入的日期时间值是有效的，不要插入`'0000-00-00'`，而是使用合法的日期时间值

   - **更改SQL模式：**如果不想更改数据，可以考虑更改MySQL的SQL模式，已允许无效的日期时间值。这是不推荐的做法，因为它可能导致数据的一致性问题。要更改SQL模式，可以执行一下SQL命令：

     ```
     SET sql_mode = 'NO_ZERO_DATE'
     ```

     或者在mysql的配置文件中永久更改SQL模式

   - **使用NULL值代替无效：**如果不能提供有效的日期时间值，可以考虑将相应的字段设置为'NULL',已表述缺少值

     ```
     INSET INTO your_table (login_time) VALUES (NULL);
     ```

     

​		