package mysql

import (
	"database/sql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Student [...]
type Student struct {
	ID     uint32 `gorm:"primaryKey;column:id;type:int(11) unsigned;not null" json:"-"`
	Name   string `gorm:"column:name;type:varchar(16);not null;default:''" json:"name"`
	Age    int8   `gorm:"column:age;type:tinyint(11);not null;default:0" json:"age"`
	Grade  int8   `gorm:"column:grade;type:tinyint(4);not null;default:0" json:"grade"`
	Class  int8   `gorm:"column:class;type:tinyint(4);not null;default:0" json:"class"`
	Status int8   `gorm:"column:status;type:tinyint(4);not null;default:0" json:"status"`
}

// TableName get sql table name.获取数据库表名
func (m *Student) TableName() string {
	return "student"
}

// StudentColumns get sql column name.获取数据库列名
var StudentColumns = struct {
	ID     string
	Name   string
	Age    string
	Grade  string
	Class  string
	Status string
}{
	ID:     "id",
	Name:   "name",
	Age:    "age",
	Grade:  "grade",
	Class:  "class",
	Status: "status",
}

type StudentDao struct {
	db *gorm.DB
}

func NewStudentDao() *StudentDao {
	return &StudentDao{
		db: getDB(),
	}
}

func (dao *StudentDao) GetStudentByClass(class int) ([]Student,error) {
	var students []Student
	err := dao.db.Model(&Student{}).Where("class", class).Find(&students).Error
	if err != nil {
		if err == sql.ErrNoRows {
			return make([]Student, 0), nil
		}
		return make([]Student, 0), errors.Wrap(err, "table:student")
	}
	return students, nil
}



