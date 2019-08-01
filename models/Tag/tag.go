package Tag

import (
	"gin-init/models"
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	models.Model

	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface {}) (tags []Tag) {
	models.Db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetTagTotal(maps interface {}) (count int){
	models.Db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	models.Db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func AddTag(name string, state int, createdBy string) bool{
	models.Db.Create(&Tag {
		Name : name,
		State : state,
		CreatedBy : createdBy,
	})

	return true
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}