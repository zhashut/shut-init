package db

import (
	"fmt"
	"go-init/global"
	"gorm.io/gorm"
)

type BaseDao struct {
}

func (d *BaseDao) GetDataByWhere(entity interface{}, where string) (interface{}, error) {
	if err := global.DB.Where(where).Find(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (d *BaseDao) GetOneByID(entity interface{}, id int) (interface{}, error) {
	if err := global.DB.Where("id = ?", id).First(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (d *BaseDao) GetOneByParam(entity interface{}, param string, value interface{}) (interface{}, error) {
	where := fmt.Sprintf("%s = ?", param)
	if err := global.DB.Where(where, value).First(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (d *BaseDao) SaveChange(entity interface{}) (interface{}, error) {
	if err := global.DB.Save(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (d *BaseDao) DeleteOneByWhere(entity interface{}, where string) error {
	if err := global.DB.Where(where).Delete(entity).Error; err != nil {
		return err
	}
	return nil
}

// Paginate 封装分页方法
func (d *BaseDao) Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 20:
			pageSize = 20
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
