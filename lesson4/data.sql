-- ----------------------------
--  Sequence structure for books_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."books_id_seq";
CREATE SEQUENCE "public"."books_id_seq" INCREMENT 1 START 6 MAXVALUE 9223372036854775807 MINVALUE 1 CACHE 1;

-- ----------------------------
--  Sequence structure for reviews_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."reviews_id_seq";
CREATE SEQUENCE "public"."reviews_id_seq" INCREMENT 1 START 1 MAXVALUE 9223372036854775807 MINVALUE 1 CACHE 1;

-- ----------------------------
--  Table structure for books
-- ----------------------------
DROP TABLE IF EXISTS "public"."books";
CREATE TABLE "public"."books" (
	"id" int4 NOT NULL DEFAULT nextval('books_id_seq'::regclass),
	"title" varchar(255) NOT NULL COLLATE "default",
	"author" varchar(40) NOT NULL COLLATE "default",
	"description" text COLLATE "default"
)
WITH (OIDS=FALSE);

-- ----------------------------
--  Records of books
-- ----------------------------
BEGIN;
INSERT INTO "public"."books" VALUES ('1', 'Swarley''s Big Day', 'Barney Stinson', null);
INSERT INTO "public"."books" VALUES ('2', 'JerBear goes to the City', 'Garnee Smashington', null);
INSERT INTO "public"."books" VALUES ('3', 'Life of a Certified G''', 'Jeremy Saenz', null);
INSERT INTO "public"."books" VALUES ('4', 'All Around the Roundabound', 'Anakin Groundsitter', null);
INSERT INTO "public"."books" VALUES ('5', 'Tommy Takes Seattle', 'Nate Beck', null);
INSERT INTO "public"."books" VALUES ('6', 'Mastering Crossfire: You''ll get caught up in it', 'Freddie Wong', null);
COMMIT;

-- ----------------------------
--  Table structure for reviews
-- ----------------------------
DROP TABLE IF EXISTS "public"."reviews";
CREATE TABLE "public"."reviews" (
	"id" int4 NOT NULL DEFAULT nextval('reviews_id_seq'::regclass),
	"content" text COLLATE "default",
	"rating" int4 NOT NULL DEFAULT 0
)
WITH (OIDS=FALSE);

-- ----------------------------
--  Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."books_id_seq" RESTART 7 OWNED BY "books"."id";
ALTER SEQUENCE "public"."reviews_id_seq" RESTART 2 OWNED BY "reviews"."id";
-- ----------------------------
--  Primary key structure for table books
-- ----------------------------
ALTER TABLE "public"."books" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Primary key structure for table reviews
-- ----------------------------
ALTER TABLE "public"."reviews" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;