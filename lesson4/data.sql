-- ----------------------------
--  Sequence structure for books_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."books_id_seq";
CREATE SEQUENCE "public"."books_id_seq" INCREMENT 1 START 6 MAXVALUE 9223372036854775807 MINVALUE 1 CACHE 1;

-- ----------------------------
--  Table structure for books
-- ----------------------------
DROP TABLE IF EXISTS "public"."books";
CREATE TABLE "public"."books" (
	"id" int4 NOT NULL DEFAULT nextval('books_id_seq'::regclass),
	"title" varchar(255) NOT NULL COLLATE "default"
)
WITH (OIDS=FALSE);

-- ----------------------------
--  Records of books
-- ----------------------------
BEGIN;
INSERT INTO "public"."books" VALUES ('2', 'JerBear goes to the City');
INSERT INTO "public"."books" VALUES ('3', 'Life of a Gangsta''');
INSERT INTO "public"."books" VALUES ('1', 'Swarley''s Big Day');
INSERT INTO "public"."books" VALUES ('4', 'All Around the Roundabound');
INSERT INTO "public"."books" VALUES ('5', 'Tommy Takes Seattle');
INSERT INTO "public"."books" VALUES ('6', 'Mastering Crossfire: You''ll get caught up in it');
COMMIT;


-- ----------------------------
--  Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."books_id_seq" RESTART 7 OWNED BY "books"."id";
-- ----------------------------
--  Primary key structure for table books
-- ----------------------------
ALTER TABLE "public"."books" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

