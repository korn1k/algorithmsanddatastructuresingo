package practice

import (
	"fmt"
	"math/rand"
	"time"
)

type Storage struct {
	data []int
	indexes map[int]int
}

func (s *Storage) Insert(val int) error {
	if _, exists := s.indexes[val]; exists {
		return fmt.Errorf("item already exists")
	}

	s.data = append(s.data, val)
	s.indexes[val] = len(s.data)-1

	return nil
}

func (s *Storage) Delete(val int) error {
	index, exists := s.indexes[val]
	if !exists {
		return fmt.Errorf("unable to delete this item")
	}

	s.data[index] = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	delete(s.indexes, val)

	return nil
}

func (s *Storage) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(s.data))
	return s.data[randomIndex]
}