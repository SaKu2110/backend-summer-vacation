package controller

import(
	"log"

	// import gin library
	"github.com/gin-gonic/gin"

	// import gorm library
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

	// import sample API packages
	"github.com/miraikeitai2020/backend-summer-vacation/pkg/server/model"
	"github.com/miraikeitai2020/backend-summer-vacation/pkg/stamp"
	"github.com/miraikeitai2020/backend-summer-vacation/pkg/zeller"
	"github.com/miraikeitai2020/backend-summer-vacation/pkg/crypto"
)

var(
	user 		model.User
	calcArgs	model.ZellerElements
	// sign data format
	sign		model.DBUser
	dbUser		model.DBUser
)

type Controller struct {
	DB	*gorm.DB
}

func (ctrl *Controller)HelloWorld(context *gin.Context) {
	context.JSON(200, gin.H{"message": "hello world"})
}

func (ctrl *Controller)SayHello(context *gin.Context) {
	err := context.BindJSON(&user)
	if err != nil {
		log.Println("[ERROR] Faild Bind JSON")
		context.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	context.JSON(200, gin.H{"message": "hello " + user.Name})
}

// 課題1
// 説明：
// 現在の日付と時間を返す.
// JSONの生成は gin.H を用いても良い
// 
// リクエスト => なし
// レスポンス => 
// {
//   "timestamp": string,
//   "detail": {
//     "date": string, //例： 2020-09-02
//     "time": string  //例: 00:00:00
//   }
// }
func (ctrl *Controller)Task1(context *gin.Context) {
	context.JSON(200, stamp.Now())
}

// 課題2
// 説明：
// ツェラーの公式でリクエストで投げた日付の曜日を返す
// JSONの生成は encoding/json を使用すること
// 
// リクエスト => 
// {
//   "year": Int,
//   "month": Int,
//   "day": Int
// }
// レスポンス => 
// {
//   "week": string //例： Monday
// }
func (ctrl *Controller)Task2(context *gin.Context) {
	err := context.BindJSON(&calcArgs)
	if err != nil {
		log.Println("[ERROR] Faild Bind JSON")
		context.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	context.JSON(200, gin.H{
		"week": zeller.Zeller(calcArgs.Year, calcArgs.Month, calcArgs.Day),
	})
}

// 課題3
// 説明：
// ユーザーIDとパスワードをデータベースに登録して, 発行したトークンを返す
// パスワードはハッシュ化したものをデータベースに登録する
// JSONの生成は encoding/json を使用すること
// 
// リクエスト => 
// {
//   "id": string,
//   "password": string
// }
// レスポンス => 
// {
//   "token": string
// }
func (ctrl *Controller)SignUp(context *gin.Context) {
	err := context.BindJSON(&sign)
	if err != nil {
		log.Println("[ERROR] Faild Bind JSON.")
		context.JSON(500, gin.H{
			"status": 500,
			"message": "Faild bind request JSON.",
		})
		return
	}
	ctrl.DB.Raw(model.QUERY_FORMAT_GET_USER, sign.ID).Scan(&dbUser)
	if dbUser.ID != "" {
		context.JSON(412, gin.H{
			"status": 412,
			"message": "Already user exist.",
		})
		return
	}
	ctrl.DB.Exec(model.QUERY_FORMAT_SET_USER, sign.ID, crypto.CreateHashWithPass(sign.Password))
	if token, err := crypto.CreateToken(sign); err == nil {
		context.JSON(201, gin.H{"token": token})
		return
	}
	log.Println("[ERROR] Faild Create Token.")
	context.JSON(412, gin.H{
		"status": 412,
		"message": "Failed create token.",
	})
}

// 課題4
// 説明：
// ユーザーIDとパスワードをデータベースに登録されたものかを照合する
// 照合が終わったら結果を返す
// JSONの生成は encoding/json を使用すること
// 
// リクエスト => 
// {
//   "id": string,
//   "password": string
// }
// レスポンス => 
// {
//   "certification": boolean
// }
func (ctrl *Controller)SignIn(context *gin.Context) {
	err := context.BindJSON(&sign)
	if err != nil {
		log.Println("[ERROR] Faild Bind JSON.")
		context.JSON(500, gin.H{
			"status": 500,
			"message": "Faild bind request JSON.",
		})
		return
	}
	ctrl.DB.Raw("SELECT id, password FROM users WHERE id = ?", sign.ID).Scan(&dbUser)
	if dbUser.ID == "" {
		context.JSON(412, gin.H{
			"status": 412,
			"message": "Already user exist.",
		})
		return
	}
	if dbUser.Password != crypto.CreateHashWithPass(sign.Password) {
		context.JSON(412, gin.H{
			"status": 412,
			"message": "Miss match password.",
		})
		return
	}
	if token, err := crypto.CreateToken(sign); err == nil {
		context.JSON(201, gin.H{"token": token})
		return
	}
	log.Println("[ERROR] Faild Create Token.")
	context.JSON(412, gin.H{
		"status": 412,
		"message": "Failed create token",
	})
}