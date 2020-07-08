package ecode

import (
	"errors"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	defer func() {
		errStr := recover()
		if errStr != "ecode: 1 already exist" {
			t.Logf("New duplicate ecode should cause panic")
			t.FailNow()
		}
	}()
	var _ error = New(1, "1")
	var _ error = New(2, "2")
	var _ error = New(1, "1")
}

func TestCause(t *testing.T) {
	e1 := New(4, "test4 error")
	var err error = e1
	e2 := Cause(err)
	if e2.Code() != 4 {
		t.Logf("parsed error code should be 4")
		t.FailNow()
	}
	fmt.Println(e2.Message())

	// error不存在时返回500
	var err3 = errors.New("test new")
	e3 := Cause(err3)
	if e3.Code() != -500 {
		t.Logf("invalid string parsed error code should be -500")
		t.FailNow()
	}
	if e3.Error() != "-500" {
		t.Logf("invalid string parsed error string should be `-500`")
		t.FailNow()
	}
}

func TestInt(t *testing.T) {
	e1 := Int(1)
	if e1.Code() != 1 {
		t.Logf("int parsed error code should be 1")
		t.FailNow()
	}
	if e1.Error() != "1" || e1.Message() != "1" {
		t.Logf("int parsed error string should be `1`")
		t.FailNow()
	}
}

func TestString(t *testing.T) {
	eStr := String("123")
	if eStr.Code() != 123 {
		t.Logf("string parsed error code should be 123")
		t.FailNow()
	}
	if eStr.Error() != "123" || eStr.Message() != "123" {
		t.Logf("string parsed error string should be `123`")
		t.FailNow()
	}
	eStr = String("test")
	if eStr.Code() != -500 {
		t.Logf("invalid string parsed error code should be -500")
		t.FailNow()
	}
	if eStr.Error() != "-500" {
		t.Logf("invalid string parsed error string should be `-500`")
		t.FailNow()
	}
}
