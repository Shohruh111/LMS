CREATE TABLE "roles"(
    "id" UUID PRIMARY KEY,
    "type" VARCHAR(20) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "users"(
    "id" UUID PRIMARY KEY,
    "role_id" UUID,
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
    "user_id" UUID ,
    "course_id" UUID,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "course_info"(
    "id" UUID PRIMARY KEY,
    "course_id" UUID ,
    "percent_of_done" NUMERIC NOT NULL,
    "remaining_exam" VARCHAR(20) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "course_report"(
    "id" UUID PRIMARY KEY,
    "course_id" UUID,
    "students" NUMERIC NOT NULL,
    "type" VARCHAR(10) NOT NULL,
    "done_all" NUMERIC NOT NULL,
    "not_done" NUMERIC NOT NULL,
    "not_started" NUMERIC NOT NULL,
    "status" BOOLEAN,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);


ALTER TABLE "users" ADD CONSTRAINT "ur_users_roles" FOREIGN KEY ("role_id") REFERENCES "roles"("id");

ALTER TABLE "courseOfUsers" ADD CONSTRAINT "cu_course_users" FOREIGN KEY ("user_id") REFERENCES "users"("id");

ALTER TABLE "courseOfUsers" ADD CONSTRAINT "cu_users_course" FOREIGN KEY ("course_id") REFERENCES "courses"("id");

ALTER TABLE "course_info" ADD CONSTRAINT "cic_info_course" FOREIGN KEY ("course_id") REFERENCES "courses"("id");

ALTER TABLE "course_report" ADD CONSTRAINT "crc_report_course" FOREIGN KEY ("course_id") REFERENCES "courses"("id");

