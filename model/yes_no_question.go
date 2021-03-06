package model

import "errors"
import "github.com/phucledien/ask-me/model"

type YesNoQuestion struct {
	Content string
	Answers []Answer
	Users   []model.user
}

func (q *YesNoQuestion) GetQuestion() (string, error) {
	return q.Content, nil
}

func (q *YesNoQuestion) GetAnswers() ([]Answer, error) {
	return q.Answers, nil
}
func (q *YesNoQuestion) GetQuestionType() (QuestionType, error) {
	return "YES_NO", nil
}
func (q *YesNoQuestion) GetRating() (int, error) {
	return len(q.User), nil
}
func (q *YesNoQuestion) UpVote(user User) error {
	for i, u := range q.Users {
		if u.ID == user.ID {
			return errors.New("USER_EXISTED")
		}
	}

	q.Users = append(q.Users, user)
	return nil
}
func (q *YesNoQuestion) DownVote(user User) error {
	for i, u := range q.Users {
		if u.ID == user.ID {
			// Remove user q.Users
			return nil
		}
	}

	return errors.New("USER_NOTEXISTED")
}
