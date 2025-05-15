package service

import (
	"context"
	"errors"
	"internal/internal/repository"
)

var (
	ErrNotFound     = errors.New("quiz not found")
	ErrNotOwner     = errors.New("you are not the creator")
	ErrDeleteFailed = errors.New("failed to delete quiz")
)

func DeleteQuiz(ctx context.Context, quizID, userID string) error {
	quiz, err := repository.GetQuizByID(ctx, quizID)
	if err != nil {
		return ErrNotFound
	}

	if quiz.CreatedBy != userID {
		return ErrNotOwner
	}

	err = repository.DeleteQuiz(ctx, quizID)
	if err != nil {
		return ErrDeleteFailed
	}

	return nil
}
