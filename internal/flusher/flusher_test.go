package flusher

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/scipie28/note-service-api/internal/app/model"
	mocksRepo "github.com/scipie28/note-service-api/internal/repo/mocks"
	"github.com/stretchr/testify/require"
)

func TestFlusher_Flush(t *testing.T) {
	var (
		mockCtrl = gomock.NewController(t)
	)

	var mockNoteRepo = mocksRepo.NewMockRepo(mockCtrl)
	noteFlusher := NewFlusher(mockNoteRepo)

	t.Run("input value equal nil", func(t *testing.T) {
		expectedRes := "invalid input values"

		_, err := noteFlusher.Flush(nil, 1)
		require.Error(t, err)
		require.Equal(t, expectedRes, err.Error())
	})

	t.Run("len input value equal zero", func(t *testing.T) {
		req := make([]model.Note, 0)
		expectedRes := "invalid input values"

		_, err := noteFlusher.Flush(req, 1)
		require.Error(t, err)
		require.Equal(t, expectedRes, err.Error())
	})

	t.Run("input one note", func(t *testing.T) {
		gomock.InOrder(mockNoteRepo.EXPECT().MultiAdd([]model.Note{
			{
				Id:          1,
				UserId:      2,
				ClassroomId: 3,
				DocumentId:  4,
			},
		},
		).Return(int64(0), nil).Times(1),
		)

		req := []model.Note{
			{
				Id:          1,
				UserId:      2,
				ClassroomId: 3,
				DocumentId:  4,
			},
		}
		var expectedRes []model.Note

		res, err := noteFlusher.Flush(req, 1)
		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("success case", func(t *testing.T) {
		gomock.InOrder(
			mockNoteRepo.EXPECT().MultiAdd([]model.Note{
				{
					Id:          1,
					UserId:      2,
					ClassroomId: 3,
					DocumentId:  4,
				},
			},
			).Return(int64(0), nil).Times(1),
			mockNoteRepo.EXPECT().MultiAdd([]model.Note{
				{
					Id:          5,
					UserId:      6,
					ClassroomId: 7,
					DocumentId:  8,
				},
			},
			).Return(int64(0), nil).Times(1),
		)

		req := []model.Note{
			{
				Id:          1,
				UserId:      2,
				ClassroomId: 3,
				DocumentId:  4,
			},
			{
				Id:          5,
				UserId:      6,
				ClassroomId: 7,
				DocumentId:  8,
			},
		}
		var expectedRes []model.Note

		res, err := noteFlusher.Flush(req, 1)
		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

}
