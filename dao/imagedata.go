package dao

import (
	"os"
	"webproject/model"
)

// GetImage 获取图片信息，返回 []model.ImageItem , error
func GetImage() ([]model.ImageItem, error) {
	// 读取文件夹
	files, err1 := os.ReadDir(model.ImagePath)
	if err1 != nil {
		return nil, err1
	}

	var imagesList []model.ImageItem
	for _, file := range files {
		if file.IsDir() { // 不读取子文件夹
			continue
		} else {
			name := model.ImagePath + file.Name()
			// 添加图片对象信息
			imagesList = append(imagesList, model.ImageItem{
				ImageName: name,
			})
		}
	}
	return imagesList, nil
}
