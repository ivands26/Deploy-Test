package data

import (
	"github.com/AltaProject/AltaSocialMedia/domain"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Comment   string `json:"comment" form:"comment"`
	UserID    int
	ContentID int
}

func (comm *Comment) ToModel() domain.Comment {
	return domain.Comment{
		ID:      int(comm.ID),
		Comment: comm.Comment,
	}
}
func ParseToArr(arr []Comment) []domain.Comment {
	var res []domain.Comment

	for _, val := range arr {
		res = append(res, val.ToModel())
	}
	return res
}

func FromModel(data domain.Comment) Comment {
	var res Comment
	res.Comment = data.Comment
	res.ID = uint(data.ID)
	res.ContentID = data.ContentID
	res.UserID = data.UserID
	return res
}
