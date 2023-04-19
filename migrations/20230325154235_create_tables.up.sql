CREATE TABLE child
(
  id      SERIAL PRIMARY KEY NOT NULL,
  first_name    TEXT   NOT NULL,
  last_name TEXT   NOT NULL,
  surname TEXT,
  user_id INT NOT NULL
);

CREATE TABLE child_groups
(
  group_id INT NOT NULL,
  child_id INT NOT NULL
);

CREATE TABLE classroom
(
  id        SERIAL PRIMARY KEY NOT NULL,
  number    INT NOT NULL,
  office_id INT NOT NULL
);

CREATE TABLE course
(
  id          SERIAL PRIMARY KEY NOT NULL,
  title       TEXT   NOT NULL,
  description TEXT,
  price       float  NOT NULL
);

CREATE TABLE course_teacher
(
  teacher_id INT NOT NULL,
  course_id  INT NOT NULL
);

CREATE TABLE groups
(
  id        SERIAL PRIMARY KEY NOT NULL,
  title     TEXT   NOT NULL,
  course_id INT NOT NULL
);

CREATE TABLE lesson
(
  id           SERIAL PRIMARY KEY NOT NULL,
  start_at     TIMESTAMP   NOT NULL,
  end_at       TIMESTAMP   NOT NULL,
  classroom_id INT NOT NULL,
  course_id    INT NOT NULL,
  teacher_id   INT NOT NULL,
  lesson_status BOOLEAN NOT NULL
);

CREATE TABLE news
(
  id           SERIAL PRIMARY KEY NOT NULL,
  title        TEXT   NOT NULL,
  description  TEXT   NOT NULL,
  created_at TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE news_users
(
  user_id INT NOT NULL,
  news_id INT NOT NULL
);

CREATE TABLE office
(
  id    SERIAL PRIMARY KEY NOT NULL,
  title TEXT   NOT NULL
);

CREATE TABLE payments
(
  id        SERIAL PRIMARY KEY NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  count     INT    NOT NULL,
  course_id INT    NOT NULL,
  user_id   INT    NOT NULL
);

CREATE TABLE teacher
(
  id         SERIAL PRIMARY KEY NOT NULL,
  first_name TEXT   NOT NULL,
  last_name  TEXT   NOT NULL,
  surname    TEXT
);

CREATE TABLE users
(
  id           SERIAL PRIMARY KEY NOT NULL,
  login        TEXT   NOT NULL UNIQUE,
  password     TEXT   NOT NULL,
  first_name    TEXT   NOT NULL,
  last_name     TEXT   NOT NULL,
  surname      TEXT,
  phone_number TEXT
);

ALTER TABLE child
  ADD CONSTRAINT FK_users_TO_child
    FOREIGN KEY (user_id)
    REFERENCES users (id);

ALTER TABLE classroom
  ADD CONSTRAINT FK_office_TO_classroom
    FOREIGN KEY (office_id)
    REFERENCES office (id);

ALTER TABLE lesson
  ADD CONSTRAINT FK_classroom_TO_lesson
    FOREIGN KEY (classroom_id)
    REFERENCES classroom (id);

ALTER TABLE lesson
  ADD CONSTRAINT FK_course_TO_lesson
    FOREIGN KEY (course_id)
    REFERENCES course (id);

ALTER TABLE lesson
  ADD CONSTRAINT FK_teacher_TO_lesson
    FOREIGN KEY (teacher_id)
    REFERENCES teacher (id);

ALTER TABLE course_teacher
  ADD CONSTRAINT FK_teacher_TO_course_teacher
    FOREIGN KEY (teacher_id)
    REFERENCES teacher (id);

ALTER TABLE course_teacher
  ADD CONSTRAINT FK_course_TO_course_teacher
    FOREIGN KEY (course_id)
    REFERENCES course (id);

ALTER TABLE groups
  ADD CONSTRAINT FK_course_TO_groups
    FOREIGN KEY (course_id)
    REFERENCES course (id);

ALTER TABLE child_groups
  ADD CONSTRAINT FK_groups_TO_child_groups
    FOREIGN KEY (group_id)
    REFERENCES groups (id);

ALTER TABLE child_groups
  ADD CONSTRAINT FK_child_TO_child_groups
    FOREIGN KEY (child_id)
    REFERENCES child (id);

ALTER TABLE payments
  ADD CONSTRAINT FK_course_TO_payments
    FOREIGN KEY (course_id)
    REFERENCES course (id);

ALTER TABLE payments
  ADD CONSTRAINT FK_users_TO_payments
    FOREIGN KEY (user_id)
    REFERENCES users (id);

ALTER TABLE news_users
  ADD CONSTRAINT FK_users_TO_news_users
    FOREIGN KEY (user_id)
    REFERENCES users (id);

ALTER TABLE news_users
  ADD CONSTRAINT FK_news_TO_news_users
    FOREIGN KEY (news_id)
    REFERENCES news (id);

ALTER TABLE news_users
  ADD PRIMARY KEY (user_id, news_id);