package model

// Subject represents a single subject and its associated score.
type Subject struct {
	Name  string  `json:"name"`
	Marks float64 `json:"marks"`
}

// Student holds personal information, subject marks, and calculated academic results.
type Student struct {
	Name       string    `json:"name"`
	RegNo      string    `json:"reg_no"`
	Subjects   []Subject `json:"subjects"`
	TotalMarks float64   `json:"total_marks"`
	Average    float64   `json:"average"`
	Grade      string    `json:"grade"`
	IsPassed   bool      `json:"is_passed"`
}

// ClassSummary aggregates performance statistics across all evaluated students.
type ClassSummary struct {
	TotalStudents  int     `json:"total_students"`
	PassedCount    int     `json:"passed_count"`
	FailedCount    int     `json:"failed_count"`
	PassPercentage float64 `json:"pass_percentage"`
	ClassAverage   float64 `json:"class_average"`
	HighestScorer  string  `json:"highest_scorer"`
	HighestScore   float64 `json:"highest_score"`
	LowestScorer   string  `json:"lowest_scorer"`
	LowestScore    float64 `json:"lowest_score"`
}
