-- Start by dropping foreign keys to avoid dependency issues
ALTER TABLE "student_select_minors" DROP FOREIGN KEY "student_select_minors_curriculum_id_fkey";
ALTER TABLE "student_select_minors" DROP FOREIGN KEY "student_select_minors_student_code_fkey";
ALTER TABLE "student_select_minors" DROP FOREIGN KEY "student_select_minors_course_no_fkey";

ALTER TABLE "plan_courses" DROP FOREIGN KEY "plan_courses_plan_id_fkey";
ALTER TABLE "plan_courses" DROP FOREIGN KEY "plan_courses_course_no_fkey";

ALTER TABLE "plans" DROP FOREIGN KEY "plans_student_code_fkey";

ALTER TABLE "course_requisites" DROP FOREIGN KEY "course_requisites_curriculum_course_id_fkey";

ALTER TABLE "curriculum_courses" DROP FOREIGN KEY "curriculum_courses_category_id_fkey";
ALTER TABLE "curriculum_courses" DROP FOREIGN KEY "curriculum_courses_course_no_fkey";

ALTER TABLE "self_categorize_members" DROP FOREIGN KEY "self_categorize_members_self_categorize_id_fkey";

ALTER TABLE "self_categorizes" DROP FOREIGN KEY "self_categorizes_curriculum_id_fkey";

ALTER TABLE "course_categories" DROP FOREIGN KEY "course_categories_self_categorize_id_fkey";
ALTER TABLE "course_categories" DROP FOREIGN KEY "course_categories_parent_id_fkey";
ALTER TABLE "course_categories" DROP FOREIGN KEY "course_categories_cross_category_id_fkey";

ALTER TABLE "curriculums" DROP FOREIGN KEY "curriculums_fs_category_id_fkey";
ALTER TABLE "curriculums" DROP FOREIGN KEY "curriculums_ge_template_id_fkey";
ALTER TABLE "curriculums" DROP FOREIGN KEY "curriculums_ge_category_id_fkey";
ALTER TABLE "curriculums" DROP FOREIGN KEY "curriculums_fe_category_id_fkey";

ALTER TABLE "ge_categories" DROP FOREIGN KEY "ge_categories_parent_id_fkey";

ALTER TABLE "ge_templates" DROP FOREIGN KEY "ge_templates_root_id_fkey";

ALTER TABLE "Account" DROP FOREIGN KEY "Account_account_type_fkey";
ALTER TABLE "Account" DROP FOREIGN KEY "Account_organization_fkey";

-- Now drop the tables
DROP TABLE IF EXISTS "organizations";
DROP TABLE IF EXISTS "account_types";
DROP TABLE IF EXISTS "moderators";
DROP TABLE IF EXISTS "students";
DROP TABLE IF EXISTS "Account";
DROP TABLE IF EXISTS "student_select_minors";
DROP TABLE IF EXISTS "plan_courses";
DROP TABLE IF EXISTS "plans";
DROP TABLE IF EXISTS "detail_courses";
DROP TABLE IF EXISTS "course_requisites";
DROP TABLE IF EXISTS "curriculum_courses";
DROP TABLE IF EXISTS "self_categorize_members";
DROP TABLE IF EXISTS "self_categorizes";
DROP TABLE IF EXISTS "course_categories";
DROP TABLE IF EXISTS "curriculums";
DROP TABLE IF EXISTS "ge_categories";
DROP TABLE IF EXISTS "ge_templates";

-- Finally, drop the types
DROP TYPE IF EXISTS "mod_roles";
DROP TYPE IF EXISTS "plan_course_type";
DROP TYPE IF EXISTS "requisite_type";
DROP TYPE IF EXISTS "category_kind";