package test

import (
	i "github.com/weiwolves/aj-captcha-go/util/image"
	"testing"
)

func TestImage_GetBackgroundImage(t *testing.T) {
	background := i.GetBackgroundImage()
	template := i.GetBackgroundImage()

	if background == nil {
		t.Fatal("背景图片获取失败")
	}
	if template == nil {
		t.Fatal("模板图片获取失败")
	}
}
