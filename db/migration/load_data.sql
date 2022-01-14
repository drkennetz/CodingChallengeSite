-- insert me as a user
insert into public.accounts (full_name, email, opted_in) values('Dennis Kennetz', 'drkennetz@gmail.com', true);
insert into public.users (account_id, admin_user, username, grade) values(1, true, 'drkennetz', 'senior');

-- add public.categories 
insert into public.categories (category) values('math');
insert into public.categories (category) values('arrays');
insert into public.categories (category) values('strings');

-- add some sample questions
insert into public.questions (challenge_name, description, example, difficulty, complexity, completion_time) values('Sum of Squares',
    'Given a positive integer n, find the smallest number of squared integers which sum to n.',
    'SumOfSquares(13)=>2; For n=13: 3^2 + 2^2 = 9 + 4 = 13 => 2', 1, 'O(NlogN) time | O(N) space', 30);
insert into public.questions (challenge_name, description, example, difficulty, complexity, completion_time) values('Potential Palindrome', 'Given a string, determine whether any permutation of it is a palindrome. A palindrome is any string that can be read the same both forwards and backwards, such as kayak',
    'PotentialPalindrome("kayak")=>true', 1, 'O(N) time | O(1) space', 20);
insert into public.questions (challenge_name, description, example, difficulty, complexity, completion_time) values('Monotonic Array',
    'Given an array, determine if the array is monotonic and return a boolean. An array is monotonic if it is entirely non-decreasing or entirely non-increasing.',
    'IsMonotonic([1, 1, 1, 1])=>true, IsMonotonic([1, 2, 3, 3]=>true, IsMonotonic([3, 2, 1])=>true, IsMonotonic([1, 2, 3, 2])=>false', 2,
    'O(N) time | O(1) space', 30);

-- create some linkages between questions and categories
insert into public.question_categories (question_id, category_id) values (1, 1);
insert into public.question_categories (question_id, category_id) values (2, 2);
insert into public.question_categories (question_id, category_id) values (3, 3);
