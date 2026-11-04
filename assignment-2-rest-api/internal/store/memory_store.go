package store

import (
	"errors"
	"fmt"
	"sort"
	"sync"

	"assignment2/internal/model"
)

var (
	ErrDuplicateID = errors.New("student with this ID already exists")
	ErrNotFound    = errors.New("student not found")
)

// Store defines the CRUD operations required for the Student resource.
type Store interface {
	Create(student model.Student) (model.Student, error)
	GetAll() []model.Student
	GetByID(id string) (model.Student, bool)
	Update(id string, input model.StudentInput) (model.Student, bool)
	Delete(id string) bool
	Reset()
}

// MemoryStore implements a thread-safe in-memory repository for Student records.
// Note: Data is kept in RAM and resets whenever the server process restarts.
type MemoryStore struct {
	mu       sync.RWMutex
	students map[string]model.Student
	nextID   int
}

// NewMemoryStore creates and returns an initialized in-memory student store.
func NewMemoryStore() *MemoryStore {
	ms := &MemoryStore{
		students: make(map[string]model.Student),
		nextID:   1,
	}
	ms.seedInitialData()
	return ms
}

// seedInitialData populates the store with initial student records for testing and demonstration.
func (ms *MemoryStore) seedInitialData() {
	initial := []model.Student{
		{
			ID:         "24BCS025",
			Name:       "Ezhilarasi K",
			Department: "Computer Science and Engineering",
			Year:       2,
			Email:      "ezhilarasi.24bcs025@ciet.edu.in",
		},
		{
			ID:         "24BCS001",
			Name:       "Aarav Sharma",
			Department: "Computer Science and Engineering",
			Year:       2,
			Email:      "aarav.sharma@ciet.edu.in",
		},
		{
			ID:         "23BIT012",
			Name:       "Kavitha M",
			Department: "Information Technology",
			Year:       3,
			Email:      "kavitha.m@ciet.edu.in",
		},
	}

	for _, s := range initial {
		ms.students[s.ID] = s
	}
}

// Create inserts a new student record into the store.
func (ms *MemoryStore) Create(s model.Student) (model.Student, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if s.ID == "" {
		s.ID = fmt.Sprintf("CIET%04d", ms.nextID)
		ms.nextID++
	}

	if _, exists := ms.students[s.ID]; exists {
		return model.Student{}, fmt.Errorf("%w: ID %s", ErrDuplicateID, s.ID)
	}

	ms.students[s.ID] = s
	return s, nil
}

// GetAll returns a slice of all students, sorted by ID for deterministic output.
func (ms *MemoryStore) GetAll() []model.Student {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	result := make([]model.Student, 0, len(ms.students))
	for _, s := range ms.students {
		result = append(result, s)
	}

	// Deterministic sorting by ID
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result
}

// GetByID retrieves a student record matching the given identifier.
func (ms *MemoryStore) GetByID(id string) (model.Student, bool) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	s, ok := ms.students[id]
	return s, ok
}

// Update modifies an existing student's mutable fields.
func (ms *MemoryStore) Update(id string, input model.StudentInput) (model.Student, bool) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	s, exists := ms.students[id]
	if !exists {
		return model.Student{}, false
	}

	s.Name = input.Name
	s.Department = input.Department
	s.Year = input.Year
	s.Email = input.Email

	ms.students[id] = s
	return s, true
}

// Delete removes a student by their identifier. Returns true if removed, false if not found.
func (ms *MemoryStore) Delete(id string) bool {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if _, exists := ms.students[id]; !exists {
		return false
	}

	delete(ms.students, id)
	return true
}

// Reset clears all records from the store (useful for clean test states).
func (ms *MemoryStore) Reset() {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.students = make(map[string]model.Student)
	ms.nextID = 1
}
