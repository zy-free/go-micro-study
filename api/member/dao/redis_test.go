package dao

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/zy-free/micro-study/api/member/model"
	"testing"
)

func Test_ExpireMemberCache(t *testing.T) {
	var (
		id = int64(1)
	)

	convey.Convey("ExpireMemberCache", t, func(ctx convey.C) {
		_, err := d.ExpireMemberCache(id)
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})
}

func Test_AddMemberCache(t *testing.T) {
	var (
		id    = int64(1)
		phone = "157"
		name  = "157_name"
		age   = int8(28)
	)

	convey.Convey("AddMemberCache", t, func(ctx convey.C) {
		err := d.AddMemberCache(id, &model.Member{
			ID:    id,
			Phone: phone,
			Name:  name,
			Age:   age,
		})
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})
}

func Test_GetMemberCache(t *testing.T) {
	var (
		id = int64(1)
	)

	convey.Convey("GetMemberCache", t, func(ctx convey.C) {
		member, err := d.GetMemberCache(id)
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
		fmt.Println(member)
	})
}

func Test_DelMemberCache(t *testing.T) {
	var (
		id = int64(1)
	)

	convey.Convey("DelMemberCache", t, func(ctx convey.C) {
		err := d.DelMemberCache(id)
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})
}
