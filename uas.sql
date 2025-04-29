-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               8.0.30 - MySQL Community Server - GPL
-- Server OS:                    Win64
-- HeidiSQL Version:             12.1.0.6537
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for learn_
DROP DATABASE IF EXISTS `learn_`;
CREATE DATABASE IF NOT EXISTS `learn_` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `learn_`;

-- Dumping structure for table learn_.announcements
DROP TABLE IF EXISTS `announcements`;
CREATE TABLE IF NOT EXISTS `announcements` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `created_by` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_announcements_created_by_user` (`created_by`),
  CONSTRAINT `fk_announcements_created_by_user` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table learn_.announcements: ~1 rows (approximately)
DELETE FROM `announcements`;
INSERT INTO `announcements` (`id`, `title`, `description`, `created_at`, `created_by`) VALUES
	(1, 'Jadwal UTS', 'UTS akan dilaksanakan mulai minggu depan.', '2025-04-29 14:02:27.477', 1);

-- Dumping structure for table learn_.attendances
DROP TABLE IF EXISTS `attendances`;
CREATE TABLE IF NOT EXISTS `attendances` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `student_id` bigint unsigned NOT NULL,
  `schedule_id` bigint unsigned NOT NULL,
  `absence_date` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_schedules_attendances` (`schedule_id`),
  KEY `fk_students_attendances` (`student_id`),
  CONSTRAINT `fk_schedules_attendances` FOREIGN KEY (`schedule_id`) REFERENCES `schedules` (`id`),
  CONSTRAINT `fk_students_attendances` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table learn_.attendances: ~0 rows (approximately)
DELETE FROM `attendances`;
INSERT INTO `attendances` (`id`, `student_id`, `schedule_id`, `absence_date`, `status`, `description`) VALUES
	(2, 1, 1, '2025-05-01', 'Hadir', '-'),
	(3, 1, 2, '2025-05-02', 'Tidak Hadir', 'Sakit');

-- Dumping structure for table learn_.classrooms
DROP TABLE IF EXISTS `classrooms`;
CREATE TABLE IF NOT EXISTS `classrooms` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `room_code` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `capacity` bigint NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_classrooms_room_code` (`room_code`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table learn_.classrooms: ~0 rows (approximately)
DELETE FROM `classrooms`;
INSERT INTO `classrooms` (`id`, `room_code`, `capacity`) VALUES
	(1, 'A101', 30),
	(2, 'B202', 25);

-- Dumping structure for table learn_.courses
DROP TABLE IF EXISTS `courses`;
CREATE TABLE IF NOT EXISTS `courses` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `course_code` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `course_name` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `sks` bigint NOT NULL,
  `semester` bigint NOT NULL,
  `study_program_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_courses_course_code` (`course_code`),
  KEY `fk_study_programs_courses` (`study_program_id`),
  CONSTRAINT `fk_study_programs_courses` FOREIGN KEY (`study_program_id`) REFERENCES `study_programs` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table learn_.courses: ~0 rows (approximately)
DELETE FROM `courses`;
INSERT INTO `courses` (`id`, `course_code`, `course_name`, `sks`, `semester`, `study_program_id`) VALUES
	(1, 'PWEB', 'Pemrograman Web Lanjut', 3, 4, 1),
	(2, 'DBMS', 'Sistem Basis Data', 3, 3, 1),
	(3, 'MPK', 'Matematika Program Studi', 2, 2, 2);

-- Dumping structure for table learn_.enrollments
DROP TABLE IF EXISTS `enrollments`;
CREATE TABLE IF NOT EXISTS `enrollments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `student_id` bigint unsigned NOT NULL,
  `course_id` bigint unsigned NOT NULL,
  `academic_year` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `semester` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_courses_enrollments` (`course_id`),
  KEY `fk_students_enrollments` (`student_id`),
  CONSTRAINT `fk_courses_enrollments` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`),
  CONSTRAINT `fk_students_enrollments` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table learn_.enrollments: ~0 rows (approximately)
DELETE FROM `enrollments`;
INSERT INTO `enrollments` (`id`, `student_id`, `course_id`, `academic_year`, `semester`) VALUES
	(1, 1, 1, '2024/2025', 'Ganjil'),
	(2, 1, 2, '2024/2025', 'Ganjil');

-- Dumping structure for table learn_.faculties
DROP TABLE IF EXISTS `faculties`;
CREATE TABLE IF NOT EXISTS `faculties` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `faculty_code` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `faculty_name` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_faculties_faculty_code` (`faculty_code`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table learn_.faculties: ~0 rows (approximately)
DELETE FROM `faculties`;
INSERT INTO `faculties` (`id`, `faculty_code`, `faculty_name`) VALUES
	(1, 'FTI', 'Fakultas Teknologi Informasi'),
	(2, 'FEB', 'Fakultas Ekonomi dan Bisnis');

-- Dumping structure for table learn_.grades
DROP TABLE IF EXISTS `grades`;
CREATE TABLE IF NOT EXISTS `grades` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `enrollment_id` bigint unsigned NOT NULL,
  `nilai_akhir` double NOT NULL,
  `grade_letter` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_enrollments_grades` (`enrollment_id`),
  CONSTRAINT `fk_enrollments_grades` FOREIGN KEY (`enrollment_id`) REFERENCES `enrollments` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table learn_.grades: ~0 rows (approximately)
DELETE FROM `grades`;
INSERT INTO `grades` (`id`, `enrollment_id`, `nilai_akhir`, `grade_letter`) VALUES
	(1, 1, 85.5, 'A'),
	(2, 2, 75, 'B');

-- Dumping structure for table learn_.lecturers
DROP TABLE IF EXISTS `lecturers`;
CREATE TABLE IF NOT EXISTS `lecturers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `n_id_n` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `date_of_birth` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `address` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `academic_position` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_lecturers_n_id_n` (`n_id_n`),
  UNIQUE KEY `uni_lecturers_user_id` (`user_id`),
  CONSTRAINT `fk_users_lecturer` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table learn_.lecturers: ~0 rows (approximately)
DELETE FROM `lecturers`;
INSERT INTO `lecturers` (`id`, `user_id`, `n_id_n`, `name`, `date_of_birth`, `address`, `academic_position`) VALUES
	(1, 3, 'D001', 'Dr. Andi Saputra', '1980-05-15', 'Jl. Pendidikan No. 10', 'Professor');

-- Dumping structure for table learn_.logs
DROP TABLE IF EXISTS `logs`;
CREATE TABLE IF NOT EXISTS `logs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned DEFAULT NULL,
  `action` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `entity_type` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `entity_id` bigint unsigned NOT NULL,
  `details` text COLLATE utf8mb4_unicode_ci,
  `created_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_logs_user` (`user_id`),
  CONSTRAINT `fk_logs_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table learn_.logs: ~1 rows (approximately)
DELETE FROM `logs`;
INSERT INTO `logs` (`id`, `user_id`, `action`, `entity_type`, `entity_id`, `details`, `created_at`) VALUES
	(1, NULL, 'CREATE', 'USER', 1, '{"id":1,"username":"aaeee","password":"$2a$14$v5RQjuSrxtDf5qbJoIP64eyWLbmtenQ5JN1pK0VKC3c1EN/x2tjqO","email":"aaee@gmail.com","role":"student","created_at":"2025-04-29T11:54:23.008+07:00","updated_at":"2025-04-29T11:54:23.008+07:00"}', '2025-04-29 11:54:23.033');

-- Dumping structure for table learn_.schedules
DROP TABLE IF EXISTS `schedules`;
CREATE TABLE IF NOT EXISTS `schedules` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `course_id` bigint unsigned NOT NULL,
  `lecturer_id` bigint unsigned NOT NULL,
  `day` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `start_at` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `end_at` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `classroom_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_lecturers_schedules` (`lecturer_id`),
  KEY `fk_classrooms_schedules` (`classroom_id`),
  KEY `fk_courses_schedules` (`course_id`),
  CONSTRAINT `fk_classrooms_schedules` FOREIGN KEY (`classroom_id`) REFERENCES `classrooms` (`id`),
  CONSTRAINT `fk_courses_schedules` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`),
  CONSTRAINT `fk_lecturers_schedules` FOREIGN KEY (`lecturer_id`) REFERENCES `lecturers` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table learn_.schedules: ~0 rows (approximately)
DELETE FROM `schedules`;
INSERT INTO `schedules` (`id`, `course_id`, `lecturer_id`, `day`, `start_at`, `end_at`, `classroom_id`) VALUES
	(1, 1, 1, 'Monday', '08:00', '10:00', 1),
	(2, 2, 1, 'Tuesday', '10:00', '12:00', 2);

-- Dumping structure for table learn_.students
DROP TABLE IF EXISTS `students`;
CREATE TABLE IF NOT EXISTS `students` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `nim` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `date_of_birth` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `address` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `study_program_id` bigint unsigned NOT NULL,
  `class_year` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_students_user_id` (`user_id`),
  UNIQUE KEY `uni_students_nim` (`nim`),
  KEY `fk_study_programs_students` (`study_program_id`),
  CONSTRAINT `fk_study_programs_students` FOREIGN KEY (`study_program_id`) REFERENCES `study_programs` (`id`),
  CONSTRAINT `fk_users_student` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table learn_.students: ~0 rows (approximately)
DELETE FROM `students`;
INSERT INTO `students` (`id`, `user_id`, `nim`, `name`, `date_of_birth`, `address`, `study_program_id`, `class_year`) VALUES
	(1, 2, '1234567890', 'Budi Santoso', '2000-01-01', 'Jl. Merdeka No. 1', 1, '2022');

-- Dumping structure for table learn_.study_programs
DROP TABLE IF EXISTS `study_programs`;
CREATE TABLE IF NOT EXISTS `study_programs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `study_program_code` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `study_program_name` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `faculty_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_study_programs_study_program_code` (`study_program_code`),
  KEY `fk_faculties_programs` (`faculty_id`),
  CONSTRAINT `fk_faculties_programs` FOREIGN KEY (`faculty_id`) REFERENCES `faculties` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table learn_.study_programs: ~0 rows (approximately)
DELETE FROM `study_programs`;
INSERT INTO `study_programs` (`id`, `study_program_code`, `study_program_name`, `faculty_id`) VALUES
	(1, 'TI', 'Teknik Informatika', 1),
	(2, 'SI', 'Sistem Informasi', 1),
	(3, 'EKONOMI', 'Ilmu Ekonomi', 2);

-- Dumping structure for table learn_.users
DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `role` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_username` (`username`),
  UNIQUE KEY `uni_users_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table learn_.users: ~4 rows (approximately)
DELETE FROM `users`;
INSERT INTO `users` (`id`, `username`, `password`, `email`, `role`, `created_at`, `updated_at`) VALUES
	(1, 'aaeee', '$2a$14$v5RQjuSrxtDf5qbJoIP64eyWLbmtenQ5JN1pK0VKC3c1EN/x2tjqO', 'aaee@gmail.com', 'student', '2025-04-29 11:54:23.008', '2025-04-29 11:54:23.008'),
	(2, 'admin123', '$2a$14$hashed_admin', 'admin@example.com', 'admin', '2025-04-29 14:02:27.422', '2025-04-29 14:02:27.422'),
	(3, 'mahasiswa1', '$2a$14$hashed_student', 'student@example.com', 'student', '2025-04-29 14:02:27.422', '2025-04-29 14:02:27.422'),
	(4, 'dosen1', '$2a$14$hashed_lecturer', 'lecturer@example.com', 'lecturer', '2025-04-29 14:02:27.422', '2025-04-29 14:02:27.422');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
