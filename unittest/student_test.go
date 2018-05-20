package main

import "testing"

func TestStudent_Save(t *testing.T) {
	type fields struct {
		Name string
		Age  int
		Male string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"load1",fields{"stu1",12,"male"},false},}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Student{
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
				Male: tt.fields.Male,
			}
			if err := s.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Student.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStudent_Load(t *testing.T) {
	type fields struct {
		Name string
		Age  int
		Male string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"load1",fields{"stu1",12,"male"},false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Student{
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
				Male: tt.fields.Male,
			}
			if err := s.Load(); (err != nil) != tt.wantErr {
				t.Errorf("Student.Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
