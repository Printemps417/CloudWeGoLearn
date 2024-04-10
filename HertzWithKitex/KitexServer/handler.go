package main

import (
	"context"
	"fmt"
	demo "kitex.demo/kitex_gen/demo"
	"sync"
)

var students sync.Map

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct{}

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, student *demo.Student) (resp *demo.RegisterResp, err error) {
	// TODO: Your code here...
	fmt.Println("Register receive:", student)
	_, exists := students.Load(student.Id)
	if exists {
		fmt.Println("Student already exists\n")
		return &demo.RegisterResp{
			Success: false,
			Message: "Student already exists\n",
		}, nil
	}

	students.Store(student.Id, student)
	//fmt.Println(student)
	fmt.Println("Student append successfully\n")
	return &demo.RegisterResp{
		Success: true,
		Message: "Student added successfully\n",
	}, nil
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *demo.QueryReq) (resp *demo.Student, err error) {
	fmt.Println("Query receive:", req)
	student, exists := students.Load(req.Id)
	if !exists {
		return nil, fmt.Errorf("Student with id %d does not exist\n", req.Id)
	}
	return student.(*demo.Student), nil
}
