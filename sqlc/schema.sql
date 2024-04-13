CREATE TABLE users (
  id     BIGSERIAL PRIMARY KEY,
  name   text      NOT NULL,
  email  text      UNIQUE
);

CREATE TABLE movies (
    "id" int8 NOT NULL,
    "user_id" int8 NOT NULL,
    "name" text,
    "watchedDate" date,
    "posterUrl" text,
    "diary" text,
    "description" text,
    "locationWatched" text,
    "releaseDate" text,
    CONSTRAINT "public_Movies_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE CASCADE,
    PRIMARY KEY ("id")
);