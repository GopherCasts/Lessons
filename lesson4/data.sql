
-- ----------------------------
--  Sequence structure for books_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."books_id_seq";
CREATE SEQUENCE "public"."books_id_seq" INCREMENT 1 START 5 MAXVALUE 9223372036854775807 MINVALUE 1 CACHE 1;

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
INSERT INTO "public"."books" VALUES ('1', 'JerBear goes to the City', 'Garnee Smashington', 'A young hipster bear seeks his fortune in the wild city of Irvine.');
INSERT INTO "public"."books" VALUES ('2', 'Swarley''s Big Day', 'Barney Stinson', 'Putting his Playbook aside, one man seeks a lifetime of happiness.');
INSERT INTO "public"."books" VALUES ('3', 'All Around the Roundabound', 'Anakin Groundsitter', 'The riveting tale of a young lad taking pod-racing lessons from an instructor with a dark secret.');
INSERT INTO "public"."books" VALUES ('4', 'Mastering Crossfire: You''ll get caught up in it', 'Freddie Wong', 'It''s sometime in the future, the ultimate challenge...  Crossfire!');
COMMIT;


-- ----------------------------
--  Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."books_id_seq" RESTART 4 OWNED BY "books"."id";

-- ----------------------------
--  Primary key structure for table books
-- ----------------------------
ALTER TABLE "public"."books" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;
