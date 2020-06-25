package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zy-free/micro-study/api/member/model"
	"github.com/zy-free/micro-study/lib/ecode"
	"github.com/zy-free/micro-study/lib/ginResult"
	"net/http"
)

func getMember(ctx *gin.Context) {
	ctx.JSON(200, "member test")
}

func addMember(ctx *gin.Context) {
	arg := &model.ArgMemberCreate{}
	if err := ctx.Bind(arg); err != nil {
		fmt.Println(err)
		ctx.JSON(ginResult.Render(http.StatusBadRequest, nil, ecode.RequestErr))
		return
	}
	id, err := svc.AddMember(arg)
	data := map[string]interface{}{
		"id": id,
	}
	ctx.JSON(ginResult.Render(http.StatusOK, data, err))

}
