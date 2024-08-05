CREATE TABLE "auth_group" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(150) NOT NULL UNIQUE
);

CREATE TABLE "auth_permission" (
    "id" SERIAL PRIMARY KEY,
    "content_type_id" INTEGER NOT NULL REFERENCES "content_type" ("id") DEFERRABLE INITIALLY DEFERRED,
    "codename" VARCHAR(100) NOT NULL,
    "name" VARCHAR(255) NOT NULL
);

CREATE TABLE "auth_group_permissions" (
    "id" SERIAL PRIMARY KEY,
    "group_id" INTEGER NOT NULL REFERENCES "auth_group" ("id") DEFERRABLE INITIALLY DEFERRED,
    "permission_id" INTEGER NOT NULL REFERENCES "auth_permission" ("id") DEFERRABLE INITIALLY DEFERRED
);

CREATE TABLE "message" (
    "id" SERIAL PRIMARY KEY,
    "body" TEXT NOT NULL,
    "updated" TIMESTAMPTZ NOT NULL,
    "created" TIMESTAMPTZ NOT NULL,
    "room_id" BIGINT NOT NULL REFERENCES "room" ("id") DEFERRABLE INITIALLY DEFERRED,
    "user_id" BIGINT NOT NULL REFERENCES "user" ("id") DEFERRABLE INITIALLY DEFERRED
);

CREATE TABLE "room" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(200) NOT NULL,
    "description" TEXT,
    "updated" TIMESTAMPTZ NOT NULL,
    "created" TIMESTAMPTZ NOT NULL,
    "host_id" BIGINT REFERENCES "user" ("id") DEFERRABLE INITIALLY DEFERRED,
    "topic_id" BIGINT REFERENCES "topic" ("id") DEFERRABLE INITIALLY DEFERRED
);

CREATE TABLE "room_participants" (
    "id" SERIAL PRIMARY KEY,
    "room_id" BIGINT NOT NULL REFERENCES "room" ("id") DEFERRABLE INITIALLY DEFERRED,
    "user_id" BIGINT NOT NULL REFERENCES "user" ("id") DEFERRABLE INITIALLY DEFERRED
);

CREATE TABLE "topic" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(200) NOT NULL
);

CREATE TABLE "user" (
    "id" SERIAL PRIMARY KEY,
    "password" VARCHAR(128) NOT NULL,
    "last_login" TIMESTAMPTZ,
    "is_superuser" BOOLEAN NOT NULL,
    "username" VARCHAR(150) NOT NULL UNIQUE,
    "first_name" VARCHAR(150) NOT NULL,
    "last_name" VARCHAR(150) NOT NULL,
    "email" VARCHAR(254) UNIQUE,
    "is_staff" BOOLEAN NOT NULL,
    "is_active" BOOLEAN NOT NULL,
    "date_joined" TIMESTAMPTZ NOT NULL,
    "bio" TEXT,
    "name" VARCHAR(200),
    "avatar" VARCHAR(100)
);

CREATE TABLE "user_groups" (
    "id" SERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL REFERENCES "user" ("id") DEFERRABLE INITIALLY DEFERRED,
    "group_id" INTEGER NOT NULL REFERENCES "auth_group" ("id") DEFERRABLE INITIALLY DEFERRED
);

CREATE TABLE "user_permissions" (
    "id" SERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL REFERENCES "user" ("id") DEFERRABLE INITIALLY DEFERRED,
    "permission_id" INTEGER NOT NULL REFERENCES "auth_permission" ("id") DEFERRABLE INITIALLY DEFERRED
);

CREATE TABLE "content_type" (
    "id" SERIAL PRIMARY KEY,
    "app_label" VARCHAR(100) NOT NULL,
    "model" VARCHAR(100) NOT NULL
);

-- Indexes for performance improvement
CREATE INDEX idx_auth_group_permissions_group_id ON auth_group_permissions (group_id);
CREATE INDEX idx_auth_group_permissions_permission_id ON auth_group_permissions (permission_id);

CREATE INDEX idx_message_room_id ON message (room_id);
CREATE INDEX idx_message_user_id ON message (user_id);

CREATE INDEX idx_room_host_id ON room (host_id);
CREATE INDEX idx_room_topic_id ON room (topic_id);

CREATE INDEX idx_room_participants_room_id ON room_participants (room_id);
CREATE INDEX idx_room_participants_user_id ON room_participants (user_id);

CREATE INDEX idx_user_groups_user_id ON user_groups (user_id);
CREATE INDEX idx_user_groups_group_id ON user_groups (group_id);

CREATE INDEX idx_user_permissions_user_id ON user_permissions (user_id);
CREATE INDEX idx_user_permissions_permission_id ON user_permissions (permission_id);

CREATE INDEX idx_auth_permission_content_type_id ON auth_permission (content_type_id);
