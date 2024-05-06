CREATE TABLE "user" (
                         "id" SERIAL PRIMARY KEY,
                         "email" varchar UNIQUE NOT NULL,
                         "hashed_password" varchar NOT NULL,
                         "name" varchar NOT NULL
);