ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "email_key";
ALTER TABLE IF EXISTS "users" DROP CONSTRAINT IF EXISTS "username_account_key";
