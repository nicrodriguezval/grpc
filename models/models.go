package models

type Student struct {
	Id   string
	Name string
	Age  int32
}

type Test struct {
	Id   string
	Name string
}

type Question struct {
	Id       string
	Question string
	Answer   string
	TestId   string
}

type Enrollment struct {
	Id        string
	StudentId string
	TestId    string
}
