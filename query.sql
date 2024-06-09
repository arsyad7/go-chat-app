CREATE TABLE "users" (
    user_id BIGSERIAL not NULL, 
    username varchar(255) NOT NULL, 
    password varchar(255) NOT NULL, 
    "created_at" timestamp DEFAULT now(),
    "created_by" varchar(100) default 'SYSTEM',
    "updated_at" timestamp DEFAULT now(),
    "updated_by" varchar(100) default 'SYSTEM',
    "deleted_at" timestamp,
    "deleted_by" varchar(100),
    CONSTRAINT users_pkey PRIMARY KEY (user_id)
);