package services

import (
	"fmt"
	"github.com/ekusuy/go_api_blog/models"
	"github.com/ekusuy/go_api_blog/repositories"
)

// PostCommentHandlerで使用することを想定したサービス
// 引数の情報をもとに新しいコメントを作り、結果を返却
func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		fmt.Println(err)
		return models.Comment{}, err
	}
	return newComment, nil
}
