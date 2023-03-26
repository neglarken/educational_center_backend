package model

import "time"

type Child struct {
	Id        int
	FirstName string
	LastName  string
	Surname   string // nullable
	UserId    int    // references to users(id)
}

type Groups struct {
	Id       int
	Title    string
	CourseId int // references to course(id)
}

type Classroom struct {
	Id       int
	Number   int
	OfficeId int // references to office(id)
}

type Course struct {
	Id          int
	Title       string
	Description string // nullable
	Price       float64
}

type Lesson struct {
	Id           int
	StartAt      time.Time
	EndAt        time.Time
	Classroom_id int // references to the classroom(id)
	CourseId     int // references to the course(id)
	TeacherId    int // references to the teacher(id)
	LessonStatus bool
}

type News struct {
	Id          int
	Title       string
	Description string
	CreatedAt   time.Time // default current_timestamp
}

type Office struct {
	Id    int
	Title string
}

type Payments struct {
	Id        int
	CreatedAt time.Time // default current_timestamp
	Count     int
	CourseId  int
	UserId    int
}

type Teacher struct {
	Id        int
	FirstName string
	LastName  string
	Surname   string // nullable
}

type ChildGroups struct {
	GroupId int // reference to the group(id)
	ChildId int // reference to the child(id)
}

type CourseTeacher struct {
	TeacherId int // reference to the teacher(id)
	CourseId  int // reference to the course(id)
}

type NewsUsers struct {
	UserId int // reference to the user(id)
	NewsId int // reference to the news(id)
}
