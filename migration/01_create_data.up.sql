CREATE TABLE "roles"(
    "id" UUID PRIMARY KEY,
    "type" VARCHAR(20) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "users"(
    "id" UUID PRIMARY KEY,
    "role_id" UUID REFERENCES "roles"("id"),
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


INSERT INTO "roles"(id, type)
VALUES('214fd852-b158-4e9a-9004-1cdc94c72835', 'Teacher');
INSERT INTO "roles"(id, type)
VALUES('1ead7347-2c79-490d-b109-c9d75dcd0bac', 'Oquvchi');

INSERT INTO "users"(id, role_id, first_name, last_name, email, phone_number, password)
VALUES('5f6242db-5c4a-4080-99a7-646db032b6fd', '214fd852-b158-4e9a-9004-1cdc94c72835', 'Adam' , 'Johns', 'adam.johns@gmail.com', '+576432176', '123456789');
INSERT INTO "users"(id, role_id, first_name, last_name, email, phone_number, password)
VALUES('cac80e04-b3eb-4c6f-b3f7-32d5cadcaa48', '1ead7347-2c79-490d-b109-c9d75dcd0bac', 'Brat', 'Pitt', 'brat.pitt@gmail.com', '+57687432', '123457689');