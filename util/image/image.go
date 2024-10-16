package image

import (
	"log"
	"os"
	"path/filepath"

	constant "github.com/weiwolves/aj-captcha-go/const"
	"github.com/weiwolves/aj-captcha-go/util"
)

var (
	backgroundImageArr      []string
	clickBackgroundImageArr []string
	templateImageArr        []string
	resourceAbsPath         string
	fontPath                string
)

func SetUp(resourcePath string, dirs ...string) {
	resourceAbsPath = resourcePath
	root := resourcePath

	var backgroundImageRoot, templateImageRoot, clickBackgroundImageRoot string
	//root := "/Users/skyline/go/src/aj-captcha-go"
	if len(dirs) == 4 {
		backgroundImageRoot = root + dirs[0]
		clickBackgroundImageRoot = root + dirs[1]
		templateImageRoot = root + dirs[2]
		fontPath = dirs[3]
	} else {
		backgroundImageRoot = root + constant.DefaultBackgroundImageDirectory
		templateImageRoot = root + constant.DefaultTemplateImageDirectory
		clickBackgroundImageRoot = root + constant.DefaultClickBackgroundImageDirectory
		fontPath = constant.DefaultFont
	}

	err := filepath.Walk(backgroundImageRoot, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		backgroundImageArr = append(backgroundImageArr, path)
		return nil
	})

	err = filepath.Walk(templateImageRoot, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		templateImageArr = append(templateImageArr, path)
		return nil
	})

	err = filepath.Walk(clickBackgroundImageRoot, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		clickBackgroundImageArr = append(clickBackgroundImageArr, path)
		return nil
	})

	if err != nil {
		log.Printf("初始化resource目录失败，请检查该目录是否存在 err: %v", err)
	}

}

func GetBackgroundImage() *util.ImageUtil {
	max := len(backgroundImageArr) - 1
	if max <= 0 {
		max = 1
	}
	return util.NewImageUtil(backgroundImageArr[util.RandomInt(0, max)], resourceAbsPath+fontPath)
}

func GetTemplateImage() *util.ImageUtil {
	max := len(templateImageArr) - 1
	if max <= 0 {
		max = 1
	}
	return util.NewImageUtil(templateImageArr[util.RandomInt(0, max)], resourceAbsPath+fontPath)
}

func GetClickBackgroundImage() *util.ImageUtil {
	max := len(templateImageArr) - 1
	if max <= 0 {
		max = 1
	}
	return util.NewImageUtil(clickBackgroundImageArr[util.RandomInt(0, max)], resourceAbsPath+constant.DefaultFont)
}
