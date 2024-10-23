package database

import (
	"PeerEdu-BackEnd/util/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func Init() {
	if db, err := gorm.Open(mysql.Open(config.Config.Dsn), &gorm.Config{}); err != nil {
		panic(err)
	} else {
		DB = db
	}

	if err := DB.AutoMigrate(
		&Question{},
		&Test{},
		&Class{},
		&Attend{},
		&CourseMember{},
		&Course{},
		&Question{},
		&Comment{},
		&Like{},
		&Poster{},
		&User{}); err != nil {
		panic(err)
	}
}

// gorm.Model 的定义
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 联合索引Model (CompoundIndex Model)
type CModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 枚举用户类型(需要与数据库中定义一致)
const (
	Role_Student = iota
	Role_Teacher
	Role_Admin
)

// 用户
type User struct {
	Model
	UserName string `gorm:"not null;unique" json:"username"` //用户名
	Password string `gorm:"not null" json:"password"`        //密码
	Role     uint   `gorm:"default:0" json:"role"`           //用户类型
	Avatar   string `json:"avatar"`                          //头像URL
}

// 帖子
type Poster struct {
	Model
	Title      string `gorm:"not null" json:"title"`
	Text       string `gorm:"not null" json:"text"`
	Images     string `json:"images"`
	Uid        string `gorm:"not null;index" json:"uid"`
	LikeNum    uint   `gorm:"default:0" json:"like_num"`    //点赞数
	CommentNum uint   `gorm:"default:0" json:"comment_num"` //评论数
}

// 点赞
type Like struct {
	CModel
	Obj string `gorm:"primaryKey" json:"obj"` //点赞对象
	Uid uint   `gorm:"primaryKey" json:"uid"` //用户id
}

// 评论
type Comment struct {
	Model
	Obj        string `gorm:"not null;index" json:"obj"`    //评论对象
	Uid        uint   `gorm:"not null;index" json:"uid"`    //用户id
	Text       string `gorm:"not null" json:"text"`         //评论内容
	LikeNum    uint   `gorm:"default:0" json:"like_num"`    //点赞数
	CommentNum uint   `gorm:"default:0" json:"comment_num"` //评论数
	ReplyObj   string `json:"reply_obj"`                    //回复对象
	ReplyUid   uint   `gorm:"default:0" json:"reply_uid"`   //回复用户id
}

// 题库
type Question struct {
	Model
	Text          string `gorm:"not null" json:"text"`     //题目内容
	Images        string `gorm:"not null" json:"images"`   //图片mids，以" "拼接，不超过10个
	OptionA       string `gorm:"not null" json:"option_a"` // 四个选项的文字描述
	OptionB       string `gorm:"not null" json:"option_b"`
	OptionC       string `gorm:"not null" json:"option_c"`
	OptionD       string `gorm:"not null" json:"option_d"`
	CorrectAnwser string `gorm:"not null" json:"correct_anwser"` // 正确答案，多选题由空格进行分隔
}

// 课程
type Course struct {
	Model
	Name       string `gorm:"not null" json:"name"`              //课程名
	Text       string `gorm:"not null" json:"text"`              //课程描述
	CreatorUid uint   `gorm:"not null;index" json:"creator_uid"` //创建课程老师id
	Tids       string `gorm:"not null;default:''" json:"tids"`   //其他加入该课程老师id，使用空格分隔，不超过10人

	AttendNum uint `gorm:"default:0" json:"attend_num"` //入课学生人数
}

// 入课成员
type CourseMember struct {
	CModel
	Cid uint `gorm:"primaryKey" json:"obj"` //课程id
	Uid uint `gorm:"primaryKey" json:"uid"` //用户id
}

// 考勤到场
type Attend struct {
	CModel
	Cid uint `gorm:"primaryKey" json:"obj"` //课程id
	Uid uint `gorm:"primaryKey" json:"uid"` //用户id
}

// 课堂
type Class struct {
	Model
	Cid        uint    `gorm:"not null;index" json:"cid"`    //课程id
	AttendRate float64 `gorm:"default:0" json:"attend_rate"` //到场率
}

// 测试
type Test struct {
	Model
	Cid        uint       `gorm:"not null;index" json:"cid"`                 //课程id
	AttendRate float64    `gorm:"default:0" json:"attend_rate"`              //整体正确率
	Questions  []Question `gorm:"many2many:test_questions" json:"questions"` //所有题目
}

// 作答情况
type Answer struct {
	Model
	Uid uint   `gorm:"not null;index" json:"uid"` //用户id
	Tid uint   `gorm:"not null;index" json:"tid"` //测试id
	Qid uint   `gorm:"not null;index" json:"qid"` //题目id
	Ans string `gorm:"not null" json:"ans"`       //正确答案，多选题由空格进行分隔
}
