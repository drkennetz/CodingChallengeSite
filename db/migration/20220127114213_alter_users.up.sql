ALTER TABLE "accounts" ADD CONSTRAINT "email_key" UNIQUE ("email");
ALTER TABLE "users" ADD CONSTRAINT "username_account_key" UNIQUE ("username", "account_id");
