package route

import (
	"api/internal/interface/controller"
	"api/internal/interface/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRoutes(router *mux.Router, db *gorm.DB) {
	
	router.Use(middleware.AuthMiddleware)
	
	userController := controller.NewUserController(db)

	router.HandleFunc("/user", userController.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", userController.GetUserById).Methods("GET")
	router.HandleFunc("/user", userController.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", userController.DeleteUser).Methods("DELETE")

	announcementController := controller.NewAnnouncementController(db)

	router.HandleFunc("/announcement", announcementController.GetAnnouncements).Methods("GET")
	router.HandleFunc("/announcement/{id}", announcementController.GetAnnouncementById).Methods("GET")
	router.HandleFunc("/announcement", announcementController.CreateAnnouncement).Methods("POST")
	router.HandleFunc("/announcement/{id}", announcementController.UpdateAnnouncement).Methods("PUT")
	router.HandleFunc("/announcement/{id}", announcementController.DeleteAnnouncement).Methods("DELETE")

	attendanceController := controller.NewAttendanceController(db)

	router.HandleFunc("/attendance", attendanceController.GetAttendances).Methods("GET")
	router.HandleFunc("/attendance/{id}", attendanceController.GetAttendanceById).Methods("GET")
	router.HandleFunc("/attendance", attendanceController.CreateAttendance).Methods("POST")
	router.HandleFunc("/attendance/{id}", attendanceController.UpdateAttendance).Methods("PUT")
	router.HandleFunc("/attendance/{id}", attendanceController.DeleteAttendance).Methods("DELETE")

	classroomController := controller.NewClassroomController(db)

	router.HandleFunc("/classroom", classroomController.GetClassrooms).Methods("GET")
	router.HandleFunc("/classroom/{id}", classroomController.GetClassroomById).Methods("GET")
	router.HandleFunc("/classroom", classroomController.CreateClassroom).Methods("POST")
	router.HandleFunc("/classroom/{id}", classroomController.UpdateClassroom).Methods("PUT")
	router.HandleFunc("/classroom/{id}", classroomController.DeleteClassroom).Methods("DELETE")

	courseController := controller.NewCourseController(db)

	router.HandleFunc("/course", courseController.GetCourses).Methods("GET")
	router.HandleFunc("/course/{id}", courseController.GetCourseById).Methods("GET")
	router.HandleFunc("/course", courseController.CreateCourse).Methods("POST")
	router.HandleFunc("/course/{id}", courseController.UpdateCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", courseController.DeleteCourse).Methods("DELETE")

	enrollmentController := controller.NewEnrollmentController(db)

	router.HandleFunc("/enrollment", enrollmentController.GetEnrollments).Methods("GET")
	router.HandleFunc("/enrollment/{id}", enrollmentController.GetEnrollmentById).Methods("GET")
	router.HandleFunc("/enrollment", enrollmentController.CreateEnrollment).Methods("POST")
	router.HandleFunc("/enrollment/{id}", enrollmentController.UpdateEnrollment).Methods("PUT")
	router.HandleFunc("/enrollment/{id}", enrollmentController.DeleteEnrollment).Methods("DELETE")


	facultyController := controller.NewFacultyController(db)

	router.HandleFunc("/faculty", facultyController.GetFacultys).Methods("GET")
	router.HandleFunc("/faculty/{id}", facultyController.GetFacultyById).Methods("GET")
	router.HandleFunc("/faculty", facultyController.CreateFaculty).Methods("POST")
	router.HandleFunc("/faculty/{id}", facultyController.UpdateFaculty).Methods("PUT")
	router.HandleFunc("/faculty/{id}", facultyController.DeleteFaculty).Methods("DELETE")

	gradeController := controller.NewGradeController(db)

	router.HandleFunc("/grade", gradeController.GetGrades).Methods("GET")
	router.HandleFunc("/grade/{id}", gradeController.GetGradeById).Methods("GET")
	router.HandleFunc("/grade", gradeController.CreateGrade).Methods("POST")
	router.HandleFunc("/grade/{id}", gradeController.UpdateGrade).Methods("PUT")
	router.HandleFunc("/grade/{id}", gradeController.DeleteGrade).Methods("DELETE")

	lecturerController := controller.NewLecturerController(db)

	router.HandleFunc("/lecturer", lecturerController.GetLecturers).Methods("GET")
	router.HandleFunc("/lecturer/{id}", lecturerController.GetLecturerById).Methods("GET")
	router.HandleFunc("/lecturer", lecturerController.CreateLecturer).Methods("POST")
	router.HandleFunc("/lecturer/{id}", lecturerController.UpdateLecturer).Methods("PUT")
	router.HandleFunc("/lecturer/{id}", lecturerController.DeleteLecturer).Methods("DELETE")

	scheduleController := controller.NewScheduleController(db)

	router.HandleFunc("/schedule", scheduleController.GetSchedules).Methods("GET")
	router.HandleFunc("/schedule/{id}", scheduleController.GetScheduleById).Methods("GET")
	router.HandleFunc("/schedule", scheduleController.CreateSchedule).Methods("POST")
	router.HandleFunc("/schedule/{id}", scheduleController.UpdateSchedule).Methods("PUT")
	router.HandleFunc("/schedule/{id}", scheduleController.DeleteSchedule).Methods("DELETE")

	studentController := controller.NewStudentController(db)

	router.HandleFunc("/student", studentController.GetStudents).Methods("GET")
	router.HandleFunc("/student/{id}", studentController.GetStudentById).Methods("GET")
	router.HandleFunc("/student", studentController.CreateStudent).Methods("POST")
	router.HandleFunc("/student/{id}", studentController.UpdateStudent).Methods("PUT")
	router.HandleFunc("/student/{id}", studentController.DeleteStudent).Methods("DELETE")

	studyProgramController := controller.NewStudyProgramController(db)

	router.HandleFunc("/study-program", studyProgramController.GetStudyPrograms).Methods("GET")
	router.HandleFunc("/study-program/{id}", studyProgramController.GetStudyProgramById).Methods("GET")
	router.HandleFunc("/study-program", studyProgramController.CreateStudyProgram).Methods("POST")
	router.HandleFunc("/study-program/{id}", studyProgramController.UpdateStudyProgram).Methods("PUT")
	router.HandleFunc("/study-program/{id}", studyProgramController.DeleteStudyProgram).Methods("DELETE")

	logController := controller.NewLogController(db)

	router.HandleFunc("/log", logController.GetLogs).Methods("GET")

}

