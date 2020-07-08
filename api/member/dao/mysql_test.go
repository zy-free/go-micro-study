package dao

import (
	"github.com/zy-free/micro-study/api/member/model"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func Test_InitMember(t *testing.T) {
	var (
		phone = "15791"
		name  = "1z"
	)

	convey.Convey("InitMember", t, func(ctx convey.C) {
		err := d.InitMember(&model.ArgMemberInit{
			Phone: phone,
			Name:  name,
		})
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})
}

func Test_AddMember(t *testing.T) {
	var (
		phone = "15792"
		name  = "1z"
	)

	convey.Convey("AddMember", t, func(ctx convey.C) {
		_, err := d.AddMember(&model.ArgMemberAdd{
			Phone: phone,
			Name:  name,
		})
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})
}

func Test_BatchAddMember(t *testing.T) {
	var (
		members []*model.ArgMemberAdd
	)
	members = append(members, &model.ArgMemberAdd{
		Phone: "1573",
		Name:  "zhouyu",
	})

	members = append(members, &model.ArgMemberAdd{
		Phone: "1574",
		Name:  "zhouyu",
	})

	convey.Convey("BatchAddMember", t, func(ctx convey.C) {
		_, err := d.BatchAddMember(members)
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})
}

func Test_DeleteMember(t *testing.T) {
	var (
		id = int64(1)
	)

	convey.Convey("DeleteMember", t, func(ctx convey.C) {
		err := d.DeleteMember(id)
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})
}

func Test_UpdateMember(t *testing.T) {
	var (
		id      = int64(1)
		name    = "zhouyu"
		age     = int8(28)
		address = "这就是"
	)

	convey.Convey("UpdateMember", t, func(ctx convey.C) {
		err := d.UpdateMember(model.ArgMemberUpdate{
			ID:      id,
			Name:    name,
			Age:     age,
			Address: address,
		})
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})
}

func Test_SetMember(t *testing.T) {
	var (
		phone   = "15794"
		name    = "zhouyu"
		age     = int8(28)
		address = "这就是"

		name2 = "zhouyu2"
	)

	convey.Convey("SetMember", t, func(ctx convey.C) {
		err := d.SetMember(model.ArgMemberSet{
			Phone:   phone,
			Name:    name,
			Age:     age,
			Address: address,
		})

		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})

	convey.Convey("SetMember", t, func(ctx convey.C) {
		err := d.SetMember(model.ArgMemberSet{
			Phone:   phone,
			Name:    name2,
			Age:     age,
			Address: address,
		})

		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})

}

func Test_BatchUpdateMemberSort(t *testing.T) {

	sort := model.ArgMemberSort{{1, 1}, {2, 2}, {3, 3}}

	convey.Convey("BatchUpdateMemberSort", t, func(ctx convey.C) {
		err := d.BatchUpdateMemberSort(sort)
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})
}

func Test_GetMemberByID(t *testing.T) {
	var (
		id = int64(2)
	)
	convey.Convey("GetMemberByID", t, func(ctx convey.C) {
		member, err := d.GetMemberByID(id)
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
		ctx.Convey("Then member should not be nil.", func(ctx convey.C) {
			ctx.So(member, convey.ShouldNotBeNil)
		})
	})
}

func Test_GetMemberMaxAge(t *testing.T) {
	convey.Convey("GetMemberMaxAge", t, func(ctx convey.C) {
		maxAge, err := d.GetMemberMaxAge()
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
		ctx.Convey("Then maxAge should be >=0.", func(ctx convey.C) {
			ctx.So(maxAge, convey.ShouldBeGreaterThanOrEqualTo, 0)
		})
	})
}

func Test_GetMemberSumAge(t *testing.T) {
	convey.Convey("GetMemberSumAge", t, func(ctx convey.C) {
		sumAge, err := d.GetMemberSumAge()
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
		ctx.Convey("Then maxAge should be >=0.", func(ctx convey.C) {
			ctx.So(sumAge, convey.ShouldBeGreaterThanOrEqualTo, 0)
		})
	})
}

func Test_CountMember(t *testing.T) {
	convey.Convey("CountMember", t, func(ctx convey.C) {
		count, err := d.CountMember()
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
		ctx.Convey("Then count should be >=0.", func(ctx convey.C) {
			ctx.So(count, convey.ShouldBeGreaterThanOrEqualTo, 0)
		})
	})
}

func Test_HasMemberByID(t *testing.T) {
	var (
		id = int64(100000)
	)
	convey.Convey("HasMemberByID", t, func(ctx convey.C) {
		has, err := d.HasMemberByID(id)
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
		ctx.Convey("Then has should be false.", func(ctx convey.C) {
			ctx.So(has, convey.ShouldBeFalse)
		})
	})
}

func Test_QueryMemberByName(t *testing.T) {
	var (
		name = "zhouyu"
	)
	convey.Convey("QueryMemberByName", t, func(ctx convey.C) {
		res, err := d.QueryMemberByName(name)
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
		ctx.Convey("Then res should  be =2.", func(ctx convey.C) {
			ctx.So(len(res), convey.ShouldEqual, 2)
		})
	})
}

func Test_QueryMemberByIDs(t *testing.T) {
	ids := []int64{4, 5}
	convey.Convey("QueryMemberByIDs", t, func(ctx convey.C) {
		res, err := d.QueryMemberByIDs(ids)
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
		ctx.Convey("Then len(res) should be =2.", func(ctx convey.C) {
			ctx.So(len(res), convey.ShouldEqual, 2)
		})
	})
}
