/*
 Navicat Premium Data Transfer

 Source Server         : Dev@Postgresql
 Source Server Type    : PostgreSQL
 Source Server Version : 150004 (150004)
 Source Host           : localhost:5432
 Source Catalog        : ccps
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 150004 (150004)
 File Encoding         : 65001

 Date: 21/02/2024 21:18:45
*/


-- ----------------------------
-- Sequence structure for api_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."api_id_seq";
CREATE SEQUENCE "public"."api_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for application_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."application_id_seq";
CREATE SEQUENCE "public"."application_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for login_log_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."login_log_id_seq";
CREATE SEQUENCE "public"."login_log_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for menu_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."menu_id_seq";
CREATE SEQUENCE "public"."menu_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for operate_log_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."operate_log_id_seq";
CREATE SEQUENCE "public"."operate_log_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for permission_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."permission_id_seq";
CREATE SEQUENCE "public"."permission_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for role_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."role_id_seq";
CREATE SEQUENCE "public"."role_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for role_permission_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."role_permission_id_seq";
CREATE SEQUENCE "public"."role_permission_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for role_user_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."role_user_id_seq";
CREATE SEQUENCE "public"."role_user_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for user_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."user_id_seq";
CREATE SEQUENCE "public"."user_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Table structure for api
-- ----------------------------
DROP TABLE IF EXISTS "public"."api";
CREATE TABLE "public"."api" (
  "id" int4 NOT NULL DEFAULT nextval('api_id_seq'::regclass),
  "title" varchar(30) COLLATE "pg_catalog"."default" NOT NULL,
  "route" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "method" varchar(10) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 'GET'::character varying,
  "ver" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT '1.0'::character varying,
  "params" json NOT NULL,
  "headers" json NOT NULL,
  "response" json NOT NULL,
  "create_time" int8 NOT NULL,
  "creator_id" int4 NOT NULL,
  "tags" varchar[] COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."api"."title" IS '名称';
COMMENT ON COLUMN "public"."api"."route" IS '路由';
COMMENT ON COLUMN "public"."api"."method" IS '方法';
COMMENT ON COLUMN "public"."api"."ver" IS '版本';
COMMENT ON COLUMN "public"."api"."params" IS '参数';
COMMENT ON COLUMN "public"."api"."headers" IS '请求头';
COMMENT ON COLUMN "public"."api"."response" IS '响应';
COMMENT ON COLUMN "public"."api"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."api"."creator_id" IS '创建人id[user_id]';
COMMENT ON TABLE "public"."api" IS 'api接口表';

-- ----------------------------
-- Records of api
-- ----------------------------

-- ----------------------------
-- Table structure for application
-- ----------------------------
DROP TABLE IF EXISTS "public"."application";
CREATE TABLE "public"."application" (
  "id" int4 NOT NULL DEFAULT nextval('application_id_seq'::regclass),
  "user_id" int4 NOT NULL,
  "name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL,
  "access_key" char(20) COLLATE "pg_catalog"."default" NOT NULL,
  "access_secret" char(64) COLLATE "pg_catalog"."default" NOT NULL,
  "remark" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "create_time" int8 NOT NULL
)
;
COMMENT ON COLUMN "public"."application"."user_id" IS '用户id';
COMMENT ON COLUMN "public"."application"."name" IS '名称';
COMMENT ON COLUMN "public"."application"."access_key" IS '访问标识';
COMMENT ON COLUMN "public"."application"."access_secret" IS '访问密钥';
COMMENT ON COLUMN "public"."application"."remark" IS '备注';
COMMENT ON COLUMN "public"."application"."create_time" IS '创建时间';
COMMENT ON TABLE "public"."application" IS '应用表';

-- ----------------------------
-- Records of application
-- ----------------------------
INSERT INTO "public"."application" VALUES (1, 1, '测试应用', 'DCwGOe1p5HzEjE4JTbJb', 'BxdgKtoLMeR4rzcDnrBJ2yGgoT8QAo0DuXY2k155u69Xz0N8UmThTuAHZrgEM7im', '', 1685720569);

-- ----------------------------
-- Table structure for login_log
-- ----------------------------
DROP TABLE IF EXISTS "public"."login_log";
CREATE TABLE "public"."login_log" (
  "id" int4 NOT NULL DEFAULT nextval('login_log_id_seq'::regclass),
  "ip" varchar(46) COLLATE "pg_catalog"."default" NOT NULL,
  "user_id" int4 NOT NULL,
  "login_time" int8 NOT NULL,
  "status" int2 NOT NULL,
  "status_msg" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "user_agent" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
COMMENT ON COLUMN "public"."login_log"."ip" IS 'ip地址';
COMMENT ON COLUMN "public"."login_log"."user_id" IS '用户id';
COMMENT ON COLUMN "public"."login_log"."login_time" IS '登陆时间';
COMMENT ON COLUMN "public"."login_log"."status" IS '登陆状态：1 - 成功，2 - 失败';
COMMENT ON COLUMN "public"."login_log"."status_msg" IS '状态对应的消息';
COMMENT ON COLUMN "public"."login_log"."user_agent" IS '用户代理';
COMMENT ON TABLE "public"."login_log" IS '登陆日志表';

-- ----------------------------
-- Records of login_log
-- ----------------------------

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS "public"."menu";
CREATE TABLE "public"."menu" (
  "id" int4 NOT NULL DEFAULT nextval('menu_id_seq'::regclass),
  "pid" int4 NOT NULL DEFAULT 0,
  "name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL,
  "path" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "permission_id" int4 NOT NULL DEFAULT 0,
  "icon" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "weight" int4 NOT NULL DEFAULT 0,
  "enabled" bool NOT NULL DEFAULT true,
  "create_time" int8 NOT NULL,
  "creator_id" int4 NOT NULL,
  "ext_lang" json NOT NULL
)
;
COMMENT ON COLUMN "public"."menu"."pid" IS '父id，0表示顶级';
COMMENT ON COLUMN "public"."menu"."name" IS '名称';
COMMENT ON COLUMN "public"."menu"."path" IS '访问路径';
COMMENT ON COLUMN "public"."menu"."permission_id" IS '访问权限id';
COMMENT ON COLUMN "public"."menu"."icon" IS '图标';
COMMENT ON COLUMN "public"."menu"."weight" IS '权重，值越大越靠前';
COMMENT ON COLUMN "public"."menu"."enabled" IS '是否启用';
COMMENT ON COLUMN "public"."menu"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."menu"."creator_id" IS '创建人id';
COMMENT ON COLUMN "public"."menu"."ext_lang" IS '扩展语言';
COMMENT ON TABLE "public"."menu" IS '菜单表';

-- ----------------------------
-- Records of menu
-- ----------------------------

-- ----------------------------
-- Table structure for operate_log
-- ----------------------------
DROP TABLE IF EXISTS "public"."operate_log";
CREATE TABLE "public"."operate_log" (
  "id" int4 NOT NULL DEFAULT nextval('operate_log_id_seq'::regclass),
  "title" varchar(30) COLLATE "pg_catalog"."default" NOT NULL,
  "type" int2 NOT NULL,
  "route" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "user_id" int4 NOT NULL,
  "operate_time" int8 NOT NULL,
  "ip" varchar(46) COLLATE "pg_catalog"."default" NOT NULL,
  "status" int2 NOT NULL,
  "status_msg" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
COMMENT ON COLUMN "public"."operate_log"."title" IS '标题';
COMMENT ON COLUMN "public"."operate_log"."type" IS '操作类型：1-新增，2-编辑，3-删除';
COMMENT ON COLUMN "public"."operate_log"."route" IS '操作路由';
COMMENT ON COLUMN "public"."operate_log"."user_id" IS '操作人id';
COMMENT ON COLUMN "public"."operate_log"."operate_time" IS '操作时间';
COMMENT ON COLUMN "public"."operate_log"."ip" IS '操作的客户端ip地址';
COMMENT ON COLUMN "public"."operate_log"."status" IS '操作状态：1-成功，2-失败';
COMMENT ON COLUMN "public"."operate_log"."status_msg" IS '操作状态对应的消息';
COMMENT ON TABLE "public"."operate_log" IS '操作日志表';

-- ----------------------------
-- Records of operate_log
-- ----------------------------

-- ----------------------------
-- Table structure for operate_log_info
-- ----------------------------
DROP TABLE IF EXISTS "public"."operate_log_info";
CREATE TABLE "public"."operate_log_info" (
  "operate_log_id" int4 NOT NULL,
  "request_headers" json NOT NULL,
  "request_params" json NOT NULL,
  "content" json NOT NULL
)
;
COMMENT ON COLUMN "public"."operate_log_info"."request_headers" IS '请求头';
COMMENT ON COLUMN "public"."operate_log_info"."request_params" IS '请求参数';
COMMENT ON COLUMN "public"."operate_log_info"."content" IS '操作内容';
COMMENT ON TABLE "public"."operate_log_info" IS '操作日志详情';

-- ----------------------------
-- Records of operate_log_info
-- ----------------------------

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS "public"."permission";
CREATE TABLE "public"."permission" (
  "id" int4 NOT NULL DEFAULT nextval('permission_id_seq'::regclass),
  "name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL,
  "key" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "route" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "remark" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "create_time" int8 NOT NULL,
  "creator_id" int4 NOT NULL,
  "pid" int4 NOT NULL DEFAULT 0,
  "realm" varchar(100) COLLATE "pg_catalog"."default" NOT NULL
)
;
COMMENT ON COLUMN "public"."permission"."name" IS '名称';
COMMENT ON COLUMN "public"."permission"."key" IS '键';
COMMENT ON COLUMN "public"."permission"."route" IS '资源路由';
COMMENT ON COLUMN "public"."permission"."remark" IS '备注';
COMMENT ON COLUMN "public"."permission"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."permission"."creator_id" IS '创建人id';
COMMENT ON COLUMN "public"."permission"."pid" IS '父id';
COMMENT ON COLUMN "public"."permission"."realm" IS '所属域';
COMMENT ON TABLE "public"."permission" IS '权限表';

-- ----------------------------
-- Records of permission
-- ----------------------------
INSERT INTO "public"."permission" VALUES (1, '权限列表', 'permission::list', '/permission', '权限列表', 1687267964, 0, 0, 'global');
INSERT INTO "public"."permission" VALUES (2, '新增权限', 'permission::add', '/permission', '权限新增', 1687267964, 0, 1, 'global');
INSERT INTO "public"."permission" VALUES (3, '编辑权限', 'permission::edit', '/permission', '权限新增', 1687267964, 0, 1, 'global');
INSERT INTO "public"."permission" VALUES (4, 'Test', 'permission::T', '/permission', '权限新增', 1687267964, 0, 2, 'global');
INSERT INTO "public"."permission" VALUES (5, 'Test1', 'permission::T1', '/permission', '权限新增', 1687267964, 0, 4, 'global');
INSERT INTO "public"."permission" VALUES (7, '新增角色', 'role::add', '/permission', '权限新增', 1687267964, 0, 6, 'global');
INSERT INTO "public"."permission" VALUES (6, '角色列表', 'role::list', '/permission', '权限新增', 1687267964, 0, 0, 'global');

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS "public"."role";
CREATE TABLE "public"."role" (
  "id" int4 NOT NULL DEFAULT nextval('role_id_seq'::regclass),
  "name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL,
  "enabled" bool NOT NULL DEFAULT true,
  "remark" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "create_time" int8 NOT NULL,
  "creator_id" int4 NOT NULL
)
;
COMMENT ON COLUMN "public"."role"."name" IS '名称';
COMMENT ON COLUMN "public"."role"."enabled" IS '是否启用';
COMMENT ON COLUMN "public"."role"."remark" IS '备注';
COMMENT ON COLUMN "public"."role"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."role"."creator_id" IS '创建人id';
COMMENT ON TABLE "public"."role" IS '角色表';

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO "public"."role" VALUES (1, 'administrator', 't', '超级管理员', 1687267964, 0);

-- ----------------------------
-- Table structure for role_permission
-- ----------------------------
DROP TABLE IF EXISTS "public"."role_permission";
CREATE TABLE "public"."role_permission" (
  "id" int4 NOT NULL DEFAULT nextval('role_permission_id_seq'::regclass),
  "role_id" int4 NOT NULL,
  "permission_id" int4 NOT NULL,
  "scope" int2 NOT NULL DEFAULT 1
)
;
COMMENT ON COLUMN "public"."role_permission"."role_id" IS '角色id';
COMMENT ON COLUMN "public"."role_permission"."permission_id" IS '权限id';
COMMENT ON COLUMN "public"."role_permission"."scope" IS '权限适用范围：1 - 全局，2 - 仅适用于登陆用户自己的数据范围内';
COMMENT ON TABLE "public"."role_permission" IS '角色权限表';

-- ----------------------------
-- Records of role_permission
-- ----------------------------
INSERT INTO "public"."role_permission" VALUES (1, 1, 1, 1);
INSERT INTO "public"."role_permission" VALUES (2, 1, 2, 1);
INSERT INTO "public"."role_permission" VALUES (3, 1, 3, 1);
INSERT INTO "public"."role_permission" VALUES (4, 1, 4, 1);
INSERT INTO "public"."role_permission" VALUES (5, 1, 5, 1);
INSERT INTO "public"."role_permission" VALUES (6, 1, 6, 1);
INSERT INTO "public"."role_permission" VALUES (7, 1, 7, 1);

-- ----------------------------
-- Table structure for role_user
-- ----------------------------
DROP TABLE IF EXISTS "public"."role_user";
CREATE TABLE "public"."role_user" (
  "id" int4 NOT NULL DEFAULT nextval('role_user_id_seq'::regclass),
  "role_id" int4 NOT NULL,
  "user_id" int4 NOT NULL
)
;
COMMENT ON COLUMN "public"."role_user"."role_id" IS '角色id';
COMMENT ON COLUMN "public"."role_user"."user_id" IS '用户id';
COMMENT ON TABLE "public"."role_user" IS '角色用户表';

-- ----------------------------
-- Records of role_user
-- ----------------------------
INSERT INTO "public"."role_user" VALUES (1, 1, 1);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "public"."user";
CREATE TABLE "public"."user" (
  "id" int4 NOT NULL DEFAULT nextval('user_id_seq'::regclass),
  "email" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "password" char(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::bpchar,
  "nickname" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "avatar" varchar(100) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "last_login_time" int8 NOT NULL DEFAULT 0,
  "last_login_ip" varchar(46) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "create_time" int8 NOT NULL,
  "enabled" bool NOT NULL DEFAULT true,
  "account" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "creator_id" int4 NOT NULL DEFAULT 0
)
;
COMMENT ON COLUMN "public"."user"."email" IS '邮箱号';
COMMENT ON COLUMN "public"."user"."password" IS '密码';
COMMENT ON COLUMN "public"."user"."nickname" IS '昵称';
COMMENT ON COLUMN "public"."user"."avatar" IS '头像';
COMMENT ON COLUMN "public"."user"."last_login_time" IS '最后登陆时间';
COMMENT ON COLUMN "public"."user"."last_login_ip" IS '最后登陆IP地址';
COMMENT ON COLUMN "public"."user"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."user"."enabled" IS '是否启用';
COMMENT ON COLUMN "public"."user"."account" IS '账号';
COMMENT ON COLUMN "public"."user"."creator_id" IS '创建人id';
COMMENT ON TABLE "public"."user" IS '用户表';

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO "public"."user" VALUES (1, '1191249254@qq.com', '$2a$04$Cxa5D/EjMz4c1ufb8XmYHus/ZnuzzUolZ1S9or5TBExY4XgxCu7C2    ', 'xianyu', '', 1687399450, '', 1687267964, 't', 'qsz', 0);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."api_id_seq"
OWNED BY "public"."api"."id";
SELECT setval('"public"."api_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."application_id_seq"
OWNED BY "public"."application"."id";
SELECT setval('"public"."application_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."login_log_id_seq"
OWNED BY "public"."login_log"."id";
SELECT setval('"public"."login_log_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."menu_id_seq"
OWNED BY "public"."menu"."id";
SELECT setval('"public"."menu_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."operate_log_id_seq"
OWNED BY "public"."operate_log"."id";
SELECT setval('"public"."operate_log_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."permission_id_seq"
OWNED BY "public"."permission"."id";
SELECT setval('"public"."permission_id_seq"', 3, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."role_id_seq"
OWNED BY "public"."role"."id";
SELECT setval('"public"."role_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."role_permission_id_seq"
OWNED BY "public"."role_permission"."id";
SELECT setval('"public"."role_permission_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."role_user_id_seq"
OWNED BY "public"."role_user"."id";
SELECT setval('"public"."role_user_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."user_id_seq"
OWNED BY "public"."user"."id";
SELECT setval('"public"."user_id_seq"', 1, true);

-- ----------------------------
-- Uniques structure for table api
-- ----------------------------
ALTER TABLE "public"."api" ADD CONSTRAINT "api_route_method_key" UNIQUE ("route", "method");

-- ----------------------------
-- Primary Key structure for table api
-- ----------------------------
ALTER TABLE "public"."api" ADD CONSTRAINT "api_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table application
-- ----------------------------
ALTER TABLE "public"."application" ADD CONSTRAINT "application_access_key_key" UNIQUE ("access_key");

-- ----------------------------
-- Primary Key structure for table application
-- ----------------------------
ALTER TABLE "public"."application" ADD CONSTRAINT "application_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table login_log
-- ----------------------------
ALTER TABLE "public"."login_log" ADD CONSTRAINT "login_log_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table menu
-- ----------------------------
ALTER TABLE "public"."menu" ADD CONSTRAINT "menu_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table operate_log
-- ----------------------------
ALTER TABLE "public"."operate_log" ADD CONSTRAINT "operate_log_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table operate_log_info
-- ----------------------------
ALTER TABLE "public"."operate_log_info" ADD CONSTRAINT "operate_log_info_pkey" PRIMARY KEY ("operate_log_id");

-- ----------------------------
-- Uniques structure for table permission
-- ----------------------------
ALTER TABLE "public"."permission" ADD CONSTRAINT "permission_key_key" UNIQUE ("key");

-- ----------------------------
-- Primary Key structure for table permission
-- ----------------------------
ALTER TABLE "public"."permission" ADD CONSTRAINT "permission_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table role
-- ----------------------------
ALTER TABLE "public"."role" ADD CONSTRAINT "role_name_key" UNIQUE ("name");

-- ----------------------------
-- Primary Key structure for table role
-- ----------------------------
ALTER TABLE "public"."role" ADD CONSTRAINT "role_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table role_permission
-- ----------------------------
ALTER TABLE "public"."role_permission" ADD CONSTRAINT "role_permission_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table role_user
-- ----------------------------
ALTER TABLE "public"."role_user" ADD CONSTRAINT "role_user_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table user
-- ----------------------------
ALTER TABLE "public"."user" ADD CONSTRAINT "user_email_key" UNIQUE ("email");
ALTER TABLE "public"."user" ADD CONSTRAINT "user_account_key" UNIQUE ("account");

-- ----------------------------
-- Primary Key structure for table user
-- ----------------------------
ALTER TABLE "public"."user" ADD CONSTRAINT "user_pkey" PRIMARY KEY ("id");
