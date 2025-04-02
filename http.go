package http

import (
	"github.com/gin-gonic/gin"
	"untitled5/internal/store"
)

type Network struct {
	Store *store.Store
}

func (n *Network) GetMeets(ctx *gin.Context) {
	var err error
	var meets []*store.Meet
	if meets, err = n.Store.GetMeets(); err != nil {
		ctx.JSON(200, gin.H{
			"error": err,
			"message": "ошибка с бд",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"meets": meets,
	})
}

func (n *Network) GetMeetById(ctx *gin.Context) {
	var err error
	var meet *store.Meet
	var requestBody struct{
		Id int `json:"id"`
	}

	if err = ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(200, gin.H{
			"error": err,
			"message": "некорректный id в запросе",
		})
	}

	if meet, err = n.Store.GetMeetById(requestBody.Id); err != nil {
		ctx.JSON(200, gin.H{
			"error": err,
			"message": "ошибка с бд",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"meet": meet,
	})
}

//func (n *Network) CreateMeet(ctx *gin.Context) {
//	var err error
//	var meet *store.Meet
//	if err = ctx.BindJSON(&meet); err != nil {
//		ctx.JSON(200, gin.H{
//			"error" : err,
//			"message" : "данные некорректны, провертье их",
//		})
//		return
//	}
//	result := n.Store.Create(meet)
//	if result.Error != nil {
//		ctx.JSON(200, gin.H{
//			"error" : result.Error,
//			"message" : "ошибка с базой данных",
//		})
//		return
//	}
//	ctx.JSON(200, gin.H{
//		"message": "new meet created",
//	})
//	return
//}
