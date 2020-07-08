package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zy-free/micro-study/api/member/model"
	bm "github.com/zy-free/micro-study/lib/blademaster"
	"github.com/zy-free/micro-study/lib/ecode"
	"net/http"
)

func getMember(ctx *gin.Context) {
	arg := &model.ArgMemberGet{}
	if err := ctx.Bind(arg); err != nil {
		ctx.JSON(bm.Render(http.StatusBadRequest, nil, ecode.RequestErr))
		return
	}
	member, err := svc.GetMember(arg.ID)
	ctx.JSON(bm.Render(http.StatusOK, member, err))
}

func addMember(ctx *gin.Context) {
	arg := &model.ArgMemberAdd{}
	if err := ctx.Bind(arg); err != nil {
		ctx.JSON(bm.Render(http.StatusBadRequest, nil, ecode.RequestErr))
		return
	}
	id, err := svc.AddMember(arg)
	data := map[string]interface{}{
		"id": id,
	}
	ctx.JSON(bm.Render(http.StatusOK, data, err))
}

func deleteMember(ctx *gin.Context) {
	arg := &model.ArgMemberDelete{}
	if err := ctx.Bind(arg); err != nil {
		ctx.JSON(bm.Render(http.StatusBadRequest, nil, ecode.RequestErr))
		return
	}
	err := svc.DeleteMember(arg.ID)
	ctx.JSON(bm.Render(http.StatusOK, nil, err))
}

func setMember(ctx *gin.Context) {
	arg := &model.ArgMemberSet{}
	if err := ctx.Bind(arg); err != nil {
		ctx.JSON(bm.Render(http.StatusBadRequest, nil, ecode.RequestErr))
		return
	}
	err := svc.SetMember(arg)
	ctx.JSON(bm.Render(http.StatusOK, nil, err))
}

func updateMember(ctx *gin.Context) {
	arg := &model.ArgMemberUpdate{}
	if err := ctx.Bind(arg); err != nil {
		fmt.Println(err)
		ctx.JSON(bm.Render(http.StatusBadRequest, nil, ecode.RequestErr))
		return
	}
	err := svc.UpdateMember(arg)
	ctx.JSON(bm.Render(http.StatusOK, nil, err))
}
