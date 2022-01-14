-- insert me as a user
insert into public.accounts (full_name, email, opted_in) values('Dennis Kennetz', 'drkennetz@gmail.com', true);
insert into public.users (account_id, username, grade) values(1, 'drkennetz', 'senior');

-- add public.categories 
insert into public.categories (category) values('arrays');
insert into public.categories (category) values('strings');
insert into public.categories (category) values('hashmaps');

-- we definitely want question_test_cases to be added as a table as well
-- id, question_id (one to many, FK), input, output