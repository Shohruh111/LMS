CREATE TABLE "roles"(
    "id" UUID PRIMARY KEY,
    "type" VARCHAR(20) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "users"(
    "id" UUID PRIMARY KEY,
    "role_id" UUID REFERENCES "roles" ("id"),
    "first_name" VARCHAR (50) NOT NULL,
    "last_name" VARCHAR(50) NOT NULL,
    "email" VARCHAR(100),
    "phone_number" VARCHAR(20),
    "password" VARCHAR(20) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "courses"(
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(30) NOT NULL,
    "photo" VARCHAR(20),
    "description" VARCHAR(50) NOT NULL,
    "weekly_number" NUMERIC NOT NULL,
    "duration" VARCHAR(10) NOT NULL,
    "price" VARCHAR(20) NOT NULL,
    "beginning_date_course" VARCHAR(20) NOT NULL,
    "end_date" VARCHAR(20) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
); 

CREATE TABLE "courseOfUsers"(
    "id" UUID PRIMARY KEY,
    "user_id" UUID REFERENCES "users"("id"),
    "course_id" UUID REFERENCES "courses"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "course_info"(
    "id" UUID PRIMARY KEY,
    "course_id" UUID REFERENCES "courses"("id"),
    "percent_of_done" NUMERIC NOT NULL,
    "remaining_exam" VARCHAR(20) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "course_report"(
    "id" UUID PRIMARY KEY,
    "course_id" UUID REFERENCES "courses"("id"),
    "students" NUMERIC NOT NULL,
    "type" VARCHAR(10) NOT NULL,
    "done_all" NUMERIC NOT NULL,
    "not_done" NUMERIC NOT NULL,
    "not_started" NUMERIC NOT NULL,
    "status" BOOLEAN,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);
