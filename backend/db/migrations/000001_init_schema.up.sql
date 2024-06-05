CREATE TABLE "user" (
                         "id" SERIAL PRIMARY KEY,
                         "email" varchar UNIQUE NOT NULL,
                         "hashed_password" varchar NOT NULL,
                         "name" varchar NOT NULL,
                         "is_admin" boolean
);

CREATE TABLE "tutorial" (
                        "id" SERIAL PRIMARY KEY,
                        "user_id" SERIAL UNIQUE NOT NULL,
                        "material" text NOT NULL,
                        "title" varchar NOT NULL
);