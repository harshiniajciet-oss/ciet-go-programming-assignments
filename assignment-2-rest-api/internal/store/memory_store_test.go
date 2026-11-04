package store

import (
	"sync"
	"testing"

	"assignment2/internal/model"
)

func TestMemoryStore_CRUD(t *testing.T) {
	s := NewMemoryStore()
	s.Reset()

	// 1. Test Create
	newStudent := model.Student{
		ID:         "24BCS025",
		Name:       "Ezhilarasi K",
		Department: "CSE",
		Year:       2,
		Email:      "ezhilarasi@ciet.edu.in",
	}

	created, err := s.Create(newStudent)
	if err != nil {
		t.Fatalf("Failed to create student: %v", err)
	}
	if created.ID != "24BCS025" {
		t.Errorf("got ID %s, want 24BCS025", created.ID)
	}

	// 2. Test Duplicate ID Error
	_, err = s.Create(newStudent)
	if err == nil {
		t.Fatalf("Expected duplicate error, got nil")
	}

	// 3. Test Auto ID generation
	autoStudent := model.Student{
		Name:       "Auto Student",
		Department: "ECE",
		Year:       1,
		Email:      "auto@ciet.edu.in",
	}
	autoCreated, err := s.Create(autoStudent)
	if err != nil {
		t.Fatalf("Failed auto create: %v", err)
	}
	if autoCreated.ID == "" {
		t.Errorf("Expected auto-generated ID, got empty")
	}

	// 4. Test GetByID
	retrieved, found := s.GetByID("24BCS025")
	if !found {
		t.Fatalf("Expected to find student 24BCS025")
	}
	if retrieved.Name != "Ezhilarasi K" {
		t.Errorf("got Name %s, want Ezhilarasi K", retrieved.Name)
	}

	_, notFound := s.GetByID("NONEXISTENT")
	if notFound {
		t.Errorf("Expected notFound to be false for NONEXISTENT")
	}

	// 5. Test GetAll
	all := s.GetAll()
	if len(all) != 2 {
		t.Errorf("got %d students, want 2", len(all))
	}

	// 6. Test Update
	updateInput := model.StudentInput{
		Name:       "Ezhilarasi K Updated",
		Department: "CSE (AI & ML)",
		Year:       3,
		Email:      "ezhilarasi.updated@ciet.edu.in",
	}
	updated, ok := s.Update("24BCS025", updateInput)
	if !ok {
		t.Fatalf("Expected update to succeed")
	}
	if updated.Name != updateInput.Name || updated.Year != 3 {
		t.Errorf("Update fields did not match: %+v", updated)
	}

	_, updateNotFound := s.Update("NONEXISTENT", updateInput)
	if updateNotFound {
		t.Errorf("Expected update on NONEXISTENT to fail")
	}

	// 7. Test Delete
	deleted := s.Delete("24BCS025")
	if !deleted {
		t.Fatalf("Expected delete to succeed")
	}

	_, stillFound := s.GetByID("24BCS025")
	if stillFound {
		t.Errorf("Student still found after delete")
	}

	deleteAgain := s.Delete("24BCS025")
	if deleteAgain {
		t.Errorf("Expected second delete to return false")
	}
}

func TestMemoryStore_ConcurrentAccess(t *testing.T) {
	s := NewMemoryStore()
	s.Reset()

	var wg sync.WaitGroup
	numRoutines := 50

	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			stud := model.Student{
				Name:       "Concurrent Student",
				Department: "CSE",
				Year:       1,
				Email:      "concurrent@ciet.edu.in",
			}
			created, err := s.Create(stud)
			if err == nil {
				_, _ = s.GetByID(created.ID)
				_ = s.GetAll()
			}
		}(i)
	}

	wg.Wait()

	if len(s.GetAll()) != numRoutines {
		t.Errorf("Expected %d students after concurrent writes, got %d", numRoutines, len(s.GetAll()))
	}
}
