CREATE TABLE "public"."user" (
  "code" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "username" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "password" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "sex" varchar(8) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "birthday" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "phone" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "email" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "province" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "city" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "county" varchar(90) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "address" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "reference" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "regtime" int8 NOT NULL DEFAULT 0,
  "remark" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "is_active" bool NOT NULL DEFAULT false,
  "is_superuser" bool NOT NULL DEFAULT false,
  "flag" bool NOT NULL DEFAULT false,
  "state" bool NOT NULL DEFAULT false,
  CONSTRAINT "user_pkey" PRIMARY KEY ("code")
)
;

ALTER TABLE "public"."user" 
  OWNER TO "postgres";

COMMENT ON COLUMN "public"."user"."code" IS '用户编号';

COMMENT ON COLUMN "public"."user"."username" IS '用户名';

COMMENT ON COLUMN "public"."user"."password" IS '用户密码';

COMMENT ON COLUMN "public"."user"."name" IS '真实姓名';

COMMENT ON COLUMN "public"."user"."sex" IS '性别';

COMMENT ON COLUMN "public"."user"."birthday" IS '生日';

COMMENT ON COLUMN "public"."user"."phone" IS '电话';

COMMENT ON COLUMN "public"."user"."email" IS '邮箱';

COMMENT ON COLUMN "public"."user"."province" IS '省份';

COMMENT ON COLUMN "public"."user"."city" IS '城市';

COMMENT ON COLUMN "public"."user"."county" IS '县区';

COMMENT ON COLUMN "public"."user"."address" IS '详细地址';

COMMENT ON COLUMN "public"."user"."reference" IS '参照着';

COMMENT ON COLUMN "public"."user"."regtime" IS '注册时间';

COMMENT ON COLUMN "public"."user"."remark" IS '备注';

COMMENT ON COLUMN "public"."user"."is_active" IS '是否激活';

COMMENT ON COLUMN "public"."user"."is_superuser" IS '是否超级管理员';

COMMENT ON COLUMN "public"."user"."flag" IS '标记';

COMMENT ON COLUMN "public"."user"."state" IS '状态';