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
    "password" VARCHAR(100) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "check_email"(
    "request_id" UUID PRIMARY KEY,
    "email" VARCHAR(50) NOT NULL,
    "verify_code" VARCHAR(6) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "expired_at" TIMESTAMP WITH TIME ZONE
);

CREATE TABLE "courses"(
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(700) NOT NULL,
    "photo" VARCHAR(700),
    "for_who" VARCHAR(500) NOT NULL,
    "type" VARCHAR(500) NOT NULL,
    "weekly_number" NUMERIC NOT NULL,
    "duration" VARCHAR(700) NOT NULL,
    "price" NUMERIC NOT NULL,
    "beginning_date_course" VARCHAR(200) NOT NULL,
    "end_date" VARCHAR(200) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
); 

CREATE TABLE "group"(
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL,
    "course_id" UUID REFERENCES "courses"("id"),
    "status" BOOLEAN,
    "end_date" TIMESTAMP,
    "beginning_date" TIMESTAMP,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "user_of_group"(
    "id" UUID PRIMARY KEY,
    "group_id" UUID REFERENCES "group"("id"),
    "user_id" UUID REFERENCES "users"("id")
);

CREATE TABLE "course_of_users"(
    "id" UUID PRIMARY KEY,
    "user_id" UUID REFERENCES "users"("id"),
    "course_id" UUID REFERENCES "courses"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "role_of_users"(
    "id" UUID PRIMARY KEY,
    "role_id" UUID REFERENCES "roles"("id") NOT NULL,
    "user_id" UUID REFERENCES "users"("id") NOT NULL
);

CREATE TABLE "lessons"(
    "id" UUID PRIMARY KEY,
    "course_id" UUID REFERENCES "courses"("id"),
    "name" VARCHAR(40) NOT NULL,
    "video_lesson" VARCHAR(200) NOT NULL,
    "status" BOOLEAN DEFAULT false,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "lesson_of_user"(
    "id" UUID PRIMARY KEY,
    "user_id" UUID REFERENCES "users"("id"),
    "lesson_id" UUID REFERENCES "lessons"("id"),
    "status" BOOLEAN
);

CREATE TABLE "photos"(
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(255),
    "data" BYTEA
);


        SELECT 
			g.id,
			g.name,
			g.course_id,
			g.status,

			c.beginning_date_course,
			c.end_date
		FROM "group" AS g
		JOIN "courses" AS c ON g.course_id = c.id
		WHERE g.course_id = '6a109912-1e30-4120-b24c-d163c183a0c5';

SELECT 

    c.id,
    l.name,
    l.video_lesson

FROM "courses" AS c
JOIN "lessons" AS l ON c.id = l.course_id



SELECT
			
			g.id,
			g.name,
			g.course_id,
			g.status,
			g.end_date,

            COUNT(ug.user_id) as students,
            COUNT(CASE WHEN lu.status = true THEN 1 ELSE NULL END) as count_true,
            COUNT(CASE WHEN lu.status = false THEN 1 ELSE NULL END) as count_false
		FROM "group"  AS g
        LEFT JOIN "user_of_group" AS ug ON g.id = ug.group_id
        LEFT JOIN "lessons" AS l ON g.course_id = l.course_id
        LEFT JOIN "lesson_of_user" AS lu ON l.id = lu.lesson_id
        WHERE g.id = '39ddcf6c-d0e7-467c-91c7-c20e67e9b065'
        GROUP BY 
            g.id,
			g.name,
			g.course_id,
			g.status,
			g.end_date;


SELECT 
        COUNT(*) FILTER (WHERE lu.status = true) AS count_true,
        COUNT(*) FILTER (WHERE lu.status = false) AS count_false
FROM "lessons" AS l
JOIN "lesson_of_user" AS lu ON l.id = lu.lesson_id;

SELECT 

    u.first_name,
    u.last_name,
    u.email,
    u.phone_number

FROM "courses" AS c 
JOIN "group" AS g ON c.id = g.course_id
JOIN "user_of_group" AS ug ON g.id = ug.group_id
JOIN "users" AS u ON ug.user_id = u.id;


INSERT INTO "group"("id", "course_id", "name","status")
VALUES
('39ddcf6c-d0e7-467c-91c7-c20e67e9b065','c1f881fb-1a8b-4fd2-9c21-a8a1e3828c1a','Group 1', true),
('07ea2fbc-476b-41ed-91c7-90a703f75d0f','c1f881fb-1a8b-4fd2-9c21-a8a1e3828c1a','Group 2', true),
('ed599851-dede-4312-9495-9c7c0368105b','c1f881fb-1a8b-4fd2-9c21-a8a1e3828c1a','Group 3', true);



INSERT INTO "user_of_group"("id", "group_id", "user_id")
VALUES
('cb2620b7-cc8b-4a9d-a8ed-b82befe94f58','39ddcf6c-d0e7-467c-91c7-c20e67e9b065','f1ee3478-9031-427f-aafa-24d8ee5187f9'),
('81ea6c01-6b14-4188-97a5-2bd945303eb8','39ddcf6c-d0e7-467c-91c7-c20e67e9b065','97693642-b842-4efa-8fc3-cc0b59a9636e'),
('c4b57f0d-f264-4840-84fe-acc27f52b14c','39ddcf6c-d0e7-467c-91c7-c20e67e9b065','9ac8bc0d-3bb1-4c8a-8b03-3be1060188b8'),
('bf0b06e0-2570-4b2d-afca-b3b33065e8e8','39ddcf6c-d0e7-467c-91c7-c20e67e9b065','fe5a761d-4940-4864-bfda-c460b332d8e1'),
('aa279cf3-cd62-4c1e-a280-75bae96075a5','39ddcf6c-d0e7-467c-91c7-c20e67e9b065','b05ca38a-560e-4b84-b255-8f7daff55e02');


SELECT 
			c.id,
			c.name,
			c.photo,
			c.for_who,
			c.type,
			c.weekly_number,
			c.duration,
			c.price,
			c.beginning_date_course,
			c.end_date,
			c.created_at,
			c.updated_at,

            l.name,
            l.video_lesson 

		FROM "courses"  AS C
		JOIN "lessons" AS l ON c.id = l.course_id
        WHERE c.id = 'bd4ae2c8-bbfc-4c9e-8249-5e3cb939f65e';


SELECT 
			ARRAY_AGG(l.name) AS lesson_names,
			ARRAY_AGG(l.video_lesson) AS video_lessons

		FROM "courses" AS c
		JOIN "lessons" AS l ON c.id = l.course_id
		WHERE c.id = 'bd4ae2c8-bbfc-4c9e-8249-5e3cb939f65e';



INSERT INTO "lessons"(id, course_id, name, video_lesson, status)
VALUES
('39c04433-9315-4b4f-a7ca-764803bbbf93','c1f881fb-1a8b-4fd2-9c21-a8a1e3828c1a','Introduction To Golang', 'adsdadsadadsad', false),
('aacbee89-d585-4777-93c4-273574508ec8','c1f881fb-1a8b-4fd2-9c21-a8a1e3828c1a', 'Go syntax(if, for, switch)', 'adsgregyhyt', false),
('e8a0711c-780f-430d-91a6-e2b5a01945ee','c1f881fb-1a8b-4fd2-9c21-a8a1e3828c1a', 'Concurrency', 'hjgfbjhrbvre', false),
('2b59d832-fe30-46ec-8a9e-66fbf4c08060','c1f881fb-1a8b-4fd2-9c21-a8a1e3828c1a', 'Functions in Go', 'rtbrwevdfvd', false),
('813b86dc-51f6-46f3-a840-b99454529f89','c1f881fb-1a8b-4fd2-9c21-a8a1e3828c1a', 'Channels', 'onnkpjnjnk', false),
('d3f75e19-3359-4283-b078-07f8f0b51633','c1f881fb-1a8b-4fd2-9c21-a8a1e3828c1a', 'Review', 'irnvfwosn', false);


INSERT INTO "lesson_of_user"(id, user_id, lesson_id, status)
VALUES
('cedc7335-9e80-452b-81df-b13ca4762d11','b05ca38a-560e-4b84-b255-8f7daff55e02','39c04433-9315-4b4f-a7ca-764803bbbf93',false),
('6c767bc6-c66b-4315-ba48-62b63f3f3c9e','b05ca38a-560e-4b84-b255-8f7daff55e02','aacbee89-d585-4777-93c4-273574508ec8',true),
('083e7cc2-26f0-432f-87bf-8b249e39168a','b05ca38a-560e-4b84-b255-8f7daff55e02','e8a0711c-780f-430d-91a6-e2b5a01945ee',false),
('ef763866-0abf-43fa-bb16-68ee05294174','b05ca38a-560e-4b84-b255-8f7daff55e02','2b59d832-fe30-46ec-8a9e-66fbf4c08060',true),
('af18a9ba-2615-4609-875b-061928c7bf03','b05ca38a-560e-4b84-b255-8f7daff55e02','813b86dc-51f6-46f3-a840-b99454529f89',true),
('b5100dc0-6eae-4859-a84f-e08a477357a6','b05ca38a-560e-4b84-b255-8f7daff55e02','d3f75e19-3359-4283-b078-07f8f0b51633',false);

INSERT INTO "roles"(id, type)
VALUES('214fd852-b158-4e9a-9004-1cdc94c72835', 'Teacher');
INSERT INTO "roles"(id, type)
VALUES('1ead7347-2c79-490d-b109-c9d75dcd0bac', 'Oquvchi');

INSERT INTO "users"(id, role_id, first_name, last_name, email, phone_number, password)
VALUES('5f6242db-5c4a-4080-99a7-646db032b6fd', '214fd852-b158-4e9a-9004-1cdc94c72835', 'Adam' , 'Johns', 'adam.johns@gmail.com', '+576432176', '123456789');
INSERT INTO "users"(id, role_id, first_name, last_name, email, phone_number, password)
VALUES('cac80e04-b3eb-4c6f-b3f7-32d5cadcaa48', '1ead7347-2c79-490d-b109-c9d75dcd0bac', 'Brat', 'Pitt', 'brat.pitt@gmail.com', '+57687432', '123457689');


SELECT 
			u.id,
			u.role_id,
			u.first_name,
			u.last_name,
			u.email,
			u.phone_number,
			u.password,
			u.created_at, 
			u.updated_at,

			r.type
		FROM "users" AS u
		JOIN "roles" AS r ON u.role_id = r.id


SELECT 
    created_at - CURRENT_TIMESTAMP as timek
FROM "product";


ALTER TABLE "group"
ADD COLUMN "end_date" TIMESTAMP;



INSERT INTO "courses" (
    "id",
    "name",
    "photo",
    "description",
    "for_who",
    "type",
    "weekly_number",
    "duration",
    "price",
    "beginning_date_course",
    "number_of_materials",
    "end_date",
    "grade"
) VALUES
    ('e94a70e3-d8d9-4f81-94b7-28c3fc37c76e', 'Искусство кулинарии', 'kulinary_art.jpg', 'Изучение искусства кулинарии.', 'Любители кулинарии', 'Очно', 3, '6 недель', 4999, '2024-03-01', 12, '2024-04-10', 88),
    ('c7c4133e-9384-421д-9e68-7b1ef1f031b2', 'Фотография природы', 'nature_photography.jpg', 'Освоение навыков фотографии природы.', 'Фотолюбители', 'Онлайн', 4, '8 недель', 3999, '2024-02-15', 10, '2024-04-05', 85),
    ('d4e07de3-c869-4f3f-96еx-f348630db62d', 'Программирование на Java', 'java_programming.jpg', 'Обучение программированию на языке Java.', 'Студенты информатики', 'Онлайн', 5, '10 недель', 5999, '2024-03-10', 15, '2024-05-01', 90),
    ('b9f8c70c-43f5-4a17-a4b2-ba4r3e194779', 'Мода и стиль', 'fashion_style.jpg', 'Погружение в мир моды и стиля.', 'Любители моды', 'Очно', 4, '8 недель', 4499, '2024-05-01', 18, '2024-06-15', 87),
    ('b5c57bb1-75eb-4bda-9e19-61804c0ba4f2', 'Развитие личности', 'personal_development.jpg', 'Развитие навыков личной эффективности.', 'Все уровни', 'Онлайн', 6, '12 недель', 6999, '2024-03-15', 20, '2024-06-30', 92),
    ('afb1c30f-371d-4d20-8f61-c02ed138991a', 'Мастерство руководителя', 'leadership_skills.jpg', 'Обучение навыкам лидерства и управления.', 'Бизнес-профессионалы', 'Онлайн', 3, '6 недель', 2999, '2024-04-05', 10, '2024-05-25', 88),
    ('efff3273-2a98-4217-87f8-1c0bc7ef6fb0', 'Творческое письмо', 'creative_writing.jpg', 'Развитие творческих навыков в письме.', 'Любители литературы', 'Очно', 4, '8 недель', 3599, '2024-05-01', 12, '2024-06-20', 86),
    ('786d5b14-d499-4d68-89f5-5791b8d4c08d', 'Танцевальное искусство', 'dance_arts.jpg', 'Обучение различным видам танцевального искусства.', 'Любители танцев', 'Онлайн', 5, '10 недель', 5199, '2024-02-20', 15, '2024-04-15', 90),
    ('f9da2ba1-3c79-4d2f-99e2-55d6894373df', 'Разработка мобильных приложений', 'mobile_app_development.jpg', 'Практика в разработке мобильных приложений.', 'Студенты программирования', 'Онлайн', 6, '12 недель', 7999, '2024-03-01', 18, '2024-06-10', 88),
    ('c8ef41d2-47d5-4c9b-9cd6-3c1e04a99ff3', 'Искусство театра', 'theater_arts.jpg', 'Погружение в искусство театра и актерского мастерства.', 'Студенты театрального искусства', 'Очно', 4, '8 недель', 4499, '2024-04-05', 14, '2024-05-25', 85),
    ('0c2ebd22-9c4f-4e2a-a33f-6d4fe3e3c145', 'Основы бизнес-анализа', 'business_analysis.jpg', 'Введение в методы бизнес-анализа.', 'Студенты бизнес-информатики', 'Онлайн', 3, '6 недель', 2999, '2024-05-10', 10, '2024-06-15', 91),
    ('a5a3a696-4d5c-4aa1-9067-314b7a15883e', 'Управление проектами', 'project_management.jpg', 'Обучение методам управления проектами.', 'Студенты управления', 'Очно', 5, '10 недель', 5999, '2024-02-15', 18, '2024-04-25', 89),
    ('c1ec854b-d4e2-42c7-9f99-07470776c6e2', 'Графический дизайн в рекламе', 'graphic_design_advertising.jpg', 'Применение графического дизайна в рекламной сфере.', 'Студенты маркетинга', 'Онлайн', 4, '8 недель', 4499, '2024-03-10', 12, '2024-05-20', 87),
    ('4f1c2a49-78f2-48a1-b0fd-ebfa4eac5bc4', 'Основы экономики', 'economics_basics.jpg', 'Введение в основы экономики.', 'Студенты экономики', 'Онлайн', 6, '12 недель', 6999, '2024-06-01', 20, '2024-08-15', 92);






   id  ed599851-dede-4312-9495-9c7c0368105b
name       | Group 3
course_id  | bd4ae2c8-bbfc-4c9e-8249-5e3cb939f65e


id         | 07ea2fbc-476b-41ed-91c7-90a703f75d0f
name       | Group 2
course_id  | bd4ae2c8-bbfc-4c9e-8249-5e3cb939f65e


id         | 39ddcf6c-d0e7-467c-91c7-c20e67e9b065
name       | Group 1
course_id  | bd4ae2c8-bbfc-4c9e-8249-5e3cb939f65e
status     | t
