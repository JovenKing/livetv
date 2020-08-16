package service

import (
	"github.com/jinzhu/gorm"
	"github.com/zjyl1994/livetv/global"
	"github.com/zjyl1994/livetv/model"
)

func GetConfig(key string) (string, error) {
	if confValue, ok := global.ConfigCache.Load(key); ok {
		return confValue.(string), nil
	} else {
		var value string
		err := global.DB.Where("name = ?", key).First(&value).Error
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return "", global.ErrConfigNotFound
			} else {
				return "", err
			}
		} else {
			global.ConfigCache.Store(key, value)
			return value, nil
		}
	}
}

func SetConfig(key, value string) error {
	data := model.Config{Name: key, Data: value}
	err := global.DB.Save(&data).Error
	if err == nil {
		global.ConfigCache.Store(key, value)
	}
	return err
}

func LoadConfigCache() error {
	var datas []model.Config
	err := global.DB.Find(&datas).Error
	if err != nil {
		return err
	} else {
		for _, v := range datas {
			global.ConfigCache.Store(v.Name, v.Data)
		}
		return nil
	}
}
