package saver

import (
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/scipie28/note-service-api/internal/alarmer"
	"github.com/scipie28/note-service-api/internal/app/model"
	"github.com/scipie28/note-service-api/internal/flusher"
	mocksRepo "github.com/scipie28/note-service-api/internal/repo/mocks"
	"github.com/stretchr/testify/require"
)

func TestSaver(t *testing.T) {
	var (
		mockCtrl        = gomock.NewController(t)
		mockNoteRepo    = mocksRepo.NewMockRepo(mockCtrl)
		flusher         = flusher.NewFlusher(mockNoteRepo)
		lossAllDataMode = true
		alarmerDuration = 20 * time.Millisecond

		req = []model.Note{
			{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
			{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
			{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
			{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7},
			{Id: 11, UserId: 11, ClassroomId: 123, DocumentId: 16},
			{Id: 21, UserId: 21, ClassroomId: 124, DocumentId: 17},
			{Id: 31, UserId: 31, ClassroomId: 123, DocumentId: 16},
			{Id: 41, UserId: 41, ClassroomId: 124, DocumentId: 17},
		}
	)

	t.Run("capacity equal zero", func(t *testing.T) {
		expectedError := "invalid capacity"

		alarmerTest, _ := alarmer.NewAlarmer(alarmerDuration)

		_, err := NewSaver(0, 3, flusher, alarmerTest, lossAllDataMode)
		require.NotNil(t, err)
		require.Equal(t, expectedError, err.Error())
	})

	t.Run("batch size equal zero", func(t *testing.T) {
		expectedError := "invalid batch size"

		alarmerTest, _ := alarmer.NewAlarmer(alarmerDuration)

		_, err := NewSaver(3, 0, flusher, alarmerTest, lossAllDataMode)
		require.NotNil(t, err)
		require.Equal(t, expectedError, err.Error())
	})

	t.Run("time alarmer less time write in buffer", func(t *testing.T) {
		var (
			callsNum = 8
		)

		t.Run("capacity equal one", func(t *testing.T) {
			mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(callsNum)

			wg := sync.WaitGroup{}
			wg.Add(callsNum)

			alarmerTest, _ := alarmer.NewAlarmer(5 * time.Millisecond)

			saverTest, err := NewSaver(1, 3, flusher, alarmerTest, lossAllDataMode)
			require.NoError(t, err)

			err = saverTest.Init()
			require.NoError(t, err)
			defer saverTest.Close()

			for _, val := range req {
				errSave := saverTest.Save(val)
				require.NoError(t, errSave)

				time.Sleep(20 * time.Millisecond)

			}

			require.Nil(t, err)
		})

		t.Run("capacity less slice notes len", func(t *testing.T) {
			mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(callsNum)

			alarmerTest, _ := alarmer.NewAlarmer(5 * time.Millisecond)

			saverTest, err := NewSaver(2, 3, flusher, alarmerTest, lossAllDataMode)
			require.NoError(t, err)

			err = saverTest.Init()
			require.NoError(t, err)
			defer saverTest.Close()

			for _, val := range req {
				errSave := saverTest.Save(val)
				require.NoError(t, errSave)

				time.Sleep(20 * time.Millisecond)
			}

			require.Nil(t, err)
		})

		t.Run("capacity more slice notes len", func(t *testing.T) {
			mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(callsNum)

			alarmerTest, _ := alarmer.NewAlarmer(5 * time.Millisecond)

			saverTest, err := NewSaver(9, 3, flusher, alarmerTest, lossAllDataMode)
			require.NoError(t, err)

			err = saverTest.Init()
			require.NoError(t, err)
			defer saverTest.Close()

			for _, val := range req {
				errSave := saverTest.Save(val)
				require.NoError(t, errSave)

				time.Sleep(20 * time.Millisecond)
			}

			require.Nil(t, err)
		})
	})

	//t.Run("time write in buffer less time alarmer", func(t *testing.T) {
	//	t.Run("capacity equal one", func(t *testing.T) {
	//		mockNoteRepo.EXPECT().MultiAddNotes(gomock.All()).Return(int64(0), nil).Times(2)
	//
	//		alarmerTest, _ := alarmer.NewAlarmer(alarmerDuration)
	//
	//		saverTest, err := NewSaver(1, 3, flusher, alarmerTest, lossAllDataMode)
	//		require.NoError(t, err)
	//
	//		err = saverTest.Init()
	//		require.NoError(t, err)
	//		defer saverTest.Close()
	//
	//		for _, val := range req {
	//			errSave := saverTest.Save(val)
	//			require.NoError(t, errSave)
	//
	//			time.Sleep(5 * time.Millisecond)
	//		}
	//
	//		require.Nil(t, err)
	//	})
	//
	//	t.Run("capacity less slice notes len", func(t *testing.T) {
	//		t.Run("batch size equal capacity", func(t *testing.T) {
	//			mockNoteRepo.EXPECT().MultiAddNotes(gomock.All()).Return(int64(0), nil).Times(2)
	//
	//			alarmerTest, _ := alarmer.NewAlarmer(alarmerDuration)
	//
	//			saverTest, err := NewSaver(6, 6, flusher, alarmerTest, lossAllDataMode)
	//			require.NoError(t, err)
	//
	//			err = saverTest.Init()
	//			require.NoError(t, err)
	//			defer saverTest.Close()
	//
	//			for _, val := range req {
	//				errSave := saverTest.Save(val)
	//				require.NoError(t, errSave)
	//
	//				time.Sleep(5 * time.Millisecond)
	//			}
	//
	//			require.Nil(t, err)
	//		})
	//
	//		t.Run("batch size less capacity", func(t *testing.T) {
	//			mockNoteRepo.EXPECT().MultiAddNotes(gomock.All()).Return(int64(0), nil).Times(5)
	//
	//			alarmerTest, _ := alarmer.NewAlarmer(alarmerDuration)
	//
	//			saverTest, err := NewSaver(6, 2, flusher, alarmerTest, lossAllDataMode)
	//			require.NoError(t, err)
	//
	//			err = saverTest.Init()
	//			require.NoError(t, err)
	//			defer saverTest.Close()
	//
	//			for _, val := range req {
	//				errSave := saverTest.Save(val)
	//				require.NoError(t, errSave)
	//
	//				time.Sleep(5 * time.Millisecond)
	//			}
	//
	//			require.Nil(t, err)
	//		})
	//
	//		t.Run("batch size more capacity", func(t *testing.T) {
	//			mockNoteRepo.EXPECT().MultiAddNotes(gomock.All()).Return(int64(0), nil).Times(2)
	//
	//			alarmerTest, _ := alarmer.NewAlarmer(alarmerDuration)
	//
	//			saverTest, err := NewSaver(2, 6, flusher, alarmerTest, lossAllDataMode)
	//			require.NoError(t, err)
	//
	//			err = saverTest.Init()
	//			require.NoError(t, err)
	//			defer saverTest.Close()
	//
	//			for _, val := range req {
	//				errSave := saverTest.Save(val)
	//				require.NoError(t, errSave)
	//
	//				time.Sleep(5 * time.Millisecond)
	//			}
	//
	//			require.Nil(t, err)
	//		})
	//	})
	//})

	//t.Run("capacity more slice notes len", func(t *testing.T) {
	//	t.Run("batch size equal capacity", func(t *testing.T) {
	//		mockNoteRepo.EXPECT().MultiAddNotes(gomock.All()).Return(int64(0), nil).Times(1)
	//
	//		alarmerTest, _ := alarmer.NewAlarmer(40 * time.Millisecond)
	//
	//		saverTest, err := NewSaver(9, 9, flusher, alarmerTest, lossAllDataMode)
	//		require.NoError(t, err)
	//
	//		err = saverTest.Init()
	//		require.NoError(t, err)
	//		defer saverTest.Close()
	//
	//		for _, val := range req {
	//			errSave := saverTest.Save(val)
	//			require.NoError(t, errSave)
	//
	//			time.Sleep(5 * time.Millisecond)
	//		}
	//
	//		require.Nil(t, err)
	//	})
	//
	//	t.Run("batch size less capacity", func(t *testing.T) {
	//		mockNoteRepo.EXPECT().MultiAddNotes(gomock.All()).Return(int64(0), nil).Times(2)
	//
	//		alarmerTest, _ := alarmer.NewAlarmer(40 * time.Millisecond)
	//
	//		saverTest, err := NewSaver(9, 4, flusher, alarmerTest, lossAllDataMode)
	//		require.NoError(t, err)
	//
	//		err = saverTest.Init()
	//		require.NoError(t, err)
	//		defer saverTest.Close()
	//
	//		for _, val := range req {
	//			errSave := saverTest.Save(val)
	//			require.NoError(t, errSave)
	//
	//			time.Sleep(5 * time.Millisecond)
	//		}
	//
	//		require.Nil(t, err)
	//	})
	//
	//	t.Run("batch size more capacity", func(t *testing.T) {
	//		mockNoteRepo.EXPECT().MultiAddNotes(gomock.All()).Return(int64(0), nil).Times(1)
	//
	//		alarmerTest, _ := alarmer.NewAlarmer(40 * time.Millisecond)
	//
	//		saverTest, err := NewSaver(9, 14, flusher, alarmerTest, lossAllDataMode)
	//		require.NoError(t, err)
	//
	//		err = saverTest.Init()
	//		require.NoError(t, err)
	//		defer saverTest.Close()
	//
	//		for _, val := range req {
	//			errSave := saverTest.Save(val)
	//			require.NoError(t, errSave)
	//
	//			time.Sleep(5 * time.Millisecond)
	//		}
	//
	//		require.Nil(t, err)
	//	})
	//})

	//t.Run("time alarmer equal time write in buffer", func(t *testing.T) {
	//	t.Run("capacity equal one", func(t *testing.T) {
	//		mockNoteRepo.EXPECT().MultiAddNotes(gomock.All()).Return(int64(0), nil).Times(8)
	//
	//		alarmerTest, _ := alarmer.NewAlarmer(10 * time.Millisecond)
	//
	//		saverTest, err := NewSaver(1, 3, flusher, alarmerTest, lossAllDataMode)
	//		require.NoError(t, err)
	//
	//		err = saverTest.Init()
	//		require.NoError(t, err)
	//		defer saverTest.Close()
	//
	//		for _, val := range req {
	//			errSave := saverTest.Save(val)
	//			require.NoError(t, errSave)
	//
	//			time.Sleep(10 * time.Millisecond)
	//		}
	//
	//		require.Nil(t, err)
	//	})
	//
	//	t.Run("capacity more slice notes len", func(t *testing.T) {
	//		mockNoteRepo.EXPECT().MultiAddNotes(gomock.All()).Return(int64(0), nil).Times(8)
	//
	//		alarmerTest, _ := alarmer.NewAlarmer(10 * time.Millisecond)
	//
	//		saverTest, err := NewSaver(9, 6, flusher, alarmerTest, lossAllDataMode)
	//		require.NoError(t, err)
	//
	//		err = saverTest.Init()
	//		require.NoError(t, err)
	//		defer saverTest.Close()
	//
	//		for _, val := range req {
	//			errSave := saverTest.Save(val)
	//			require.NoError(t, errSave)
	//
	//			time.Sleep(10 * time.Millisecond)
	//		}
	//
	//		require.Nil(t, err)
	//	})
	//
	//	t.Run("capacity less slice notes len", func(t *testing.T) {
	//		mockNoteRepo.EXPECT().MultiAddNotes(gomock.All()).Return(int64(0), nil).Times(8)
	//
	//		alarmerTest, _ := alarmer.NewAlarmer(10 * time.Millisecond)
	//
	//		saverTest, err := NewSaver(3, 2, flusher, alarmerTest, lossAllDataMode)
	//		require.NoError(t, err)
	//
	//		err = saverTest.Init()
	//		require.NoError(t, err)
	//		defer saverTest.Close()
	//
	//		for _, val := range req {
	//			errSave := saverTest.Save(val)
	//			require.NoError(t, errSave)
	//
	//			time.Sleep(10 * time.Millisecond)
	//		}
	//
	//		require.Nil(t, err)
	//	})
	//})
}
