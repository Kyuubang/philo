package api

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

type CoursesResponse struct {
	Courses []Courses `json:"courses"`
}

type Courses struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type LabsResponse struct {
	Course string `json:"course"`
	Labs   []Labs `json:"labs"`
}

type Labs struct {
	ID  int    `json:"id"`
	Lab string `json:"name"`
}

type CreateLabsResponse struct {
	ID       int    `json:"id"`
	Message  string `json:"message"`
	Name     string `json:"lab"`
	CourseID int    `json:"course_id"`
}

type CreateCourseResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	Name    string `json:"course"`
}

type CreateClassResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	Name    string `json:"class"`
}

type GetScoreResponse struct {
	LabName string `json:"lab_name"`
	Score   int    `json:"score"`
}

type ScoreData struct {
	Username string `json:"username"`
	Lab      string `json:"lab"`
	Score    int    `json:"score"`
}

type GeneralResponse struct {
	Message string `json:"message"`
}
