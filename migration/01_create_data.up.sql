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
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
); 

CREATE TABLE "course_of_users"(
    "id" UUID PRIMARY KEY,
    "user_id" UUID REFERENCES "users"("id"),
    "course_id" UUID REFERENCES "courses"("id")
);

CREATE TABLE "role_of_users"(
    "id" UUID PRIMARY KEY,
    "role_id" UUID NOT NULL,
    "user_id" UUID NOT NULL
);

CREATE TABLE "lessons"(
    
);

CREATE TABLE "photos"(
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(255),
    "data" BYTEA
);


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