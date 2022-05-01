package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	globFile := GetFileNameFormer()

	f := excelize.NewFile()
	index := f.NewSheet("文件(名)清单")
	for id, files := range globFile {
		err := f.SetCellValue("文件(名)清单", "A"+strconv.Itoa(id+1), files)
		if err != nil {
			fmt.Println(err)
		}
		f.SetActiveSheet(index)
		err = f.SaveAs("Book1.xlsx")
		if err != nil {
			fmt.Println(err)
		}
	}

}

func CreateFiles(num int) {
	for i := 1; i <= 6; i++ {
		_, err := os.Create(fmt.Sprintf("./T8b-1N-%d_裁剪模型_裁剪图片.dwg", i))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func GetFileNameFormer() []string {
	glob, err := filepath.Glob("./*.dwg")
	if err != nil {
		fmt.Println(err)
	}
	fileName := []string{}
	for _, files := range glob {
		fileNameWithSuffix := path.Base(files)
		fileType := path.Ext(fileNameWithSuffix)
		fileNameOnly := strings.TrimSuffix(fileNameWithSuffix, fileType)
		fileName = append(fileName, fileNameOnly)
	}
	return fileName
}
