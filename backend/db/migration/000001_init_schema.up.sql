CREATE TABLE "users" (
  "id_user" serial PRIMARY KEY,
  "id_role" smallint,
  "id_kota" integer,
  "fullname" VARCHAR(50) NOT NULL,
  "username" VARCHAR(50) NOT NULL,
  "email" VARCHAR(100) unique NOT NULL,
  "no_handphone" smallint NOT NULL,
  "password" VARCHAR(250) NOT NULL,
  "tanggal_lahir" DATE NOT NULL,
  "alamat" VARCHAR(50) NOT NULL,
  "description" text,
  "created_at" timestamp not null default current_timestamp,
  "updated_at" timestamp not null default current_timestamp,
  "deleted_at" timestamp not null default current_timestamp
);
