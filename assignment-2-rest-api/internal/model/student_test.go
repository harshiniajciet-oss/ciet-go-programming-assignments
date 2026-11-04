package model

import (
	"testing"
)

func TestStudentInput_Validate(t *testing.T) {
	tests := []struct {
		name    string
		input   StudentInput
		wantErr bool
	}{
		{
			name: "Valid input",
			input: StudentInput{
				Name:       "Ezhilarasi K",
				Department: "Computer Science and Engineering",
				Year:       2,
				Email:      "ezhilarasi.24bcs025@ciet.edu.in",
			},
			wantErr: false,
		},
		{
			name: "Empty name",
			input: StudentInput{
				Name:       "  ",
				Department: "CSE",
				Year:       2,
				Email:      "test@ciet.edu.in",
			},
			wantErr: true,
		},
		{
			name: "Empty department",
			input: StudentInput{
				Name:       "Ezhilarasi K",
				Department: " ",
				Year:       2,
				Email:      "test@ciet.edu.in",
			},
			wantErr: true,
		},
		{
			name: "Year below range (0)",
			input: StudentInput{
				Name:       "Ezhilarasi K",
				Department: "CSE",
				Year:       0,
				Email:      "test@ciet.edu.in",
			},
			wantErr: true,
		},
		{
			name: "Year above range (6)",
			input: StudentInput{
				Name:       "Ezhilarasi K",
				Department: "CSE",
				Year:       6,
				Email:      "test@ciet.edu.in",
			},
			wantErr: true,
		},
		{
			name: "Empty email",
			input: StudentInput{
				Name:       "Ezhilarasi K",
				Department: "CSE",
				Year:       2,
				Email:      "",
			},
			wantErr: true,
		},
		{
			name: "Invalid email syntax without domain dot",
			input: StudentInput{
				Name:       "Ezhilarasi K",
				Department: "CSE",
				Year:       2,
				Email:      "invalid@com",
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.input.Validate()
			if (err != nil) != tc.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
