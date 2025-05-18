package service

import (
	"context"
	"errors"
	"internal/internal/dto"
	"internal/internal/repository"
)

var (
	ErrNotFound     = errors.New("quiz not found")
	ErrNotOwner     = errors.New("you are not the creator")
	ErrDeleteFailed = errors.New("failed to delete quiz")
)

var ErrQuizNotFound = errors.New("quiz not found")
var ErrForbidden = errors.New("not your quiz")

func DeleteQuiz(ctx context.Context, quizID string) error {
	return repository.DeleteQuiz(ctx, quizID)
}

func UpdateQuiz(ctx context.Context, quizID string, input dto.QuizInput) error {
	// удалить старые вопросы и ответы
	err := repository.DeleteQuestionsByQuizID(ctx, quizID)
	if err != nil {
		return err
	}

	// обновить название
	err = repository.UpdateQuizTitle(ctx, quizID, input.Title)
	if err != nil {
		return err
	}

	// сохранить новые вопросы и ответы
	return repository.AddQuestionsWithAnswers(ctx, quizID, input.Questions)
}
