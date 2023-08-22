CREATE TABLE "users" (
    "id" UUID PRIMARY KEY,
    "login" VARCHAR NOT NULL,
    "password" VARCHAR NOT NULL,
    "name" VARCHAR NOT NULL,
    "age" INT NOT NULL
);

CREATE TABLE "phones" (
    "id" UUID PRIMARY KEY,
    "user_id" UUID NOT NULL REFERENCES "users"("id"),
    "phone" VARCHAR(12) NOT NULL,
    "descriprion" VARCHAR,
    "is_fax" BOOLEAN
);