package calculator

import (
	"errors"
	"math"

	"assignment1/internal/model"
)

// PassMark defines the minimum passing mark per subject and for overall average.
const PassMark = 50.0

// CalculateTotal calculates the sum of marks obtained across all subjects.
func CalculateTotal(subjects []model.Subject) float64 {
	total := 0.0
	for _, sub := range subjects {
		total += sub.Marks
	}
	return total
}

// CalculateAverage calculates the arithmetic mean of marks.
// Returns an error if the subjects slice is empty.
func CalculateAverage(total float64, subjectCount int) (float64, error) {
	if subjectCount <= 0 {
		return 0.0, errors.New("subject count must be greater than zero")
	}
	avg := total / float64(subjectCount)
	// Round to two decimal places
	return math.Round(avg*100) / 100, nil
}

// DetermineGrade assigns a letter grade based on the documented academic criteria:
// - A: 90 - 100
// - B: 80 - 89.99...
// - C: 70 - 79.99...
// - D: 60 - 69.99...
// - E: 50 - 59.99...
// - F: below 50 (or any failed subject)
func DetermineGrade(average float64, hasFailedSubject bool) string {
	if hasFailedSubject || average < PassMark {
		return "F"
	}

	switch {
	case average >= 90.0:
		return "A"
	case average >= 80.0:
		return "B"
	case average >= 70.0:
		return "C"
	case average >= 60.0:
		return "D"
	case average >= 50.0:
		return "E"
	default:
		return "F"
	}
}

// HasFailedSubject checks if any individual subject mark is below the pass mark (50.0).
func HasFailedSubject(subjects []model.Subject) bool {
	for _, sub := range subjects {
		if sub.Marks < PassMark {
			return true
		}
	}
	return false
}

// ProcessStudent computes the Total, Average, Grade, and Pass/Fail status for a student.
func ProcessStudent(s *model.Student) error {
	if len(s.Subjects) == 0 {
		return errors.New("cannot process student with no subjects")
	}

	s.TotalMarks = CalculateTotal(s.Subjects)
	avg, err := CalculateAverage(s.TotalMarks, len(s.Subjects))
	if err != nil {
		return err
	}
	s.Average = avg

	failedSub := HasFailedSubject(s.Subjects)
	s.IsPassed = !failedSub && s.Average >= PassMark
	s.Grade = DetermineGrade(s.Average, failedSub)

	return nil
}

// CalculateClassSummary aggregates key performance metrics for a group of students.
func CalculateClassSummary(students []model.Student) model.ClassSummary {
	if len(students) == 0 {
		return model.ClassSummary{}
	}

	summary := model.ClassSummary{
		TotalStudents: len(students),
		HighestScore:  -1.0,
		LowestScore:   1000000.0,
	}

	totalAverageSum := 0.0

	for _, s := range students {
		if s.IsPassed {
			summary.PassedCount++
		} else {
			summary.FailedCount++
		}

		totalAverageSum += s.Average

		if s.Average > summary.HighestScore {
			summary.HighestScore = s.Average
			summary.HighestScorer = s.Name + " (" + s.RegNo + ")"
		}

		if s.Average < summary.LowestScore {
			summary.LowestScore = s.Average
			summary.LowestScorer = s.Name + " (" + s.RegNo + ")"
		}
	}

	summary.PassPercentage = math.Round((float64(summary.PassedCount)/float64(summary.TotalStudents)*100)*100) / 100
	summary.ClassAverage = math.Round((totalAverageSum/float64(summary.TotalStudents))*100) / 100

	return summary
}
