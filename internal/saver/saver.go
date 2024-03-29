package saver

//go:generate mockgen -destination=mocks/mock_saver.go -package=mocks . Saver

import (
	"errors"
	"log"

	"github.com/scipie28/note-service-api/internal/alarmer"
	"github.com/scipie28/note-service-api/internal/app/model"
	"github.com/scipie28/note-service-api/internal/flusher"
)

// Saver ...
type Saver interface {
	Save(note model.Note) error
	Init() error
	Close()
}

type saver struct {
	capacity    int64
	batchSize   int64
	flusher     flusher.Flusher
	alarmer     alarmer.Alarmer
	notes       []model.Note
	notesChan   chan model.Note
	end         chan struct{}
	lossAllData bool
	isInit      bool
}

func NewSaver(capacity, batchSize int64, flusher flusher.Flusher, alarmer alarmer.Alarmer, lossAllData bool) (Saver, error) {
	if flusher == nil {
		return nil, errors.New("invalid flusher")
	}

	if alarmer == nil {
		return nil, errors.New("invalid alarmer")
	}

	if capacity <= 0 {
		return nil, errors.New("invalid capacity")
	}

	if batchSize <= 0 {
		return nil, errors.New("invalid batch size")
	}

	return &saver{
		capacity:    capacity,
		batchSize:   batchSize,
		flusher:     flusher,
		alarmer:     alarmer,
		lossAllData: lossAllData,
		notes:       []model.Note{},
		notesChan:   make(chan model.Note),
		end:         make(chan struct{}),
		isInit:      false,
	}, nil
}

func (s *saver) Init() error {
	if s.isInit {
		return errors.New("the saver has already been initialized")
	}

	err := s.alarmer.Init()
	if err != nil {
		return err
	}

	go func() {
		defer close(s.notesChan)
		defer close(s.end)
		defer s.alarmer.Close()

		for {
			select {
			case note := <-s.notesChan:
				s.saveToBuffer(note)
			case _, ok := <-s.alarmer.Alarm():
				if ok {
					s.flushData()
				}
			case <-s.end:
				return
			}
		}
	}()

	s.isInit = true

	return nil
}

func (s *saver) Save(note model.Note) error {
	if !s.isInit {
		return errors.New("failed to initializing saver")
	}

	s.notesChan <- note

	return nil
}

func (s *saver) Close() {
	s.end <- struct{}{}
	s.flushData()
}

func (s *saver) saveToBuffer(note model.Note) {
	if int64(len(s.notes)) >= s.capacity {
		if s.lossAllData || s.capacity == 1 {
			s.notes = s.notes[:0]
		} else {
			s.notes = s.notes[1:]
		}
	}

	s.notes = append(s.notes, note)
}

func (s *saver) flushData() {

	res, err := s.flusher.Flush(s.notes, s.batchSize)
	if err != nil {
		log.Printf("failed to flushing data %s", err.Error())
	}

	s.notes = s.notes[:copy(s.notes, res)]
}
