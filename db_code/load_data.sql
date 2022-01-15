-- insert me as a user
insert into public.accounts (full_name, email, opted_in) values('Dennis Kennetz', 'drkennetz@gmail.com', true);
insert into public.users (account_id, admin_user, username, grade) values(1, true, 'drkennetz', 'senior');

-- add public.categories 
insert into public.categories (category) values('math');
insert into public.categories (category) values('strings');
insert into public.categories (category) values('arrays');


-- add some sample questions
insert into public.questions (challenge_name, description, example, difficulty, complexity, completion_time) values('Sum of Squares',
    'Given a positive integer n, find the smallest number of squared integers which sum to n.',
    'SumOfSquares(13)=>2; For n=13: 3^2 + 2^2 = 9 + 4 = 13 => 2', 1, 'O(NlogN) time | O(N) space', 30);
insert into public.questions (challenge_name, description, example, difficulty, complexity, completion_time) values('Potential Palindrome',
    'Given a string, determine whether any permutation of it is a palindrome. A palindrome is any string that can be read the same both forwards and backwards, such as kayak. Case and punctuation / spaces should be ignored.',
    'PotentialPalindrome("kayak")=>true; PotentialPalindrome("Kayak")=>true; PotentialPalindrome("Yo Banana boy.")=>true', 1, 'O(N) time | O(1) space', 20);
insert into public.questions (challenge_name, description, example, difficulty, complexity, completion_time) values('Monotonic Array',
    'Given an array, determine if the array is monotonic and return a boolean. An array is monotonic if it is entirely non-decreasing or entirely non-increasing.',
    'IsMonotonic([1, 1, 1, 1])=>true, IsMonotonic([1, 2, 3, 3]=>true, IsMonotonic([3, 2, 1])=>true, IsMonotonic([1, 2, 3, 2])=>false', 2,
    'O(N) time | O(1) space', 30);

-- create some linkages between questions and categories
insert into public.question_categories (question_id, category_id) values (1, 1);
insert into public.question_categories (question_id, category_id) values (2, 2);
insert into public.question_categories (question_id, category_id) values (3, 3);

-- need multiple test cases per question id
-- add some test cases to the questions
insert into public.question_test_cases (question_id, inputs, outputs) values (1, '{"math": 13}', '{"math": 2}');
insert into public.question_test_cases (question_id, inputs, outputs) values (1, '{"math": 27}', '{"math": 3}');
insert into public.question_test_cases (question_id, inputs, outputs) values (1, '{"math": 144}', '{"math": 1}');
insert into public.question_test_cases (question_id, inputs, outputs) values (1, '{"math": 84}', '{"math": 3}');
insert into public.question_test_cases (question_id, inputs, outputs) values (1, '{"math": 85}', '{"math": 2}');
insert into public.question_test_cases (question_id, inputs, outputs) values (2, '{"strings": "kayak"}', '{"strings": true}');
insert into public.question_test_cases (question_id, inputs, outputs) values (2, '{"strings": "Kayak"}', '{"strings": true}');
insert into public.question_test_cases (question_id, inputs, outputs) values (2, '{"strings": "Yo Banana boy."}', '{"strings": true}');
insert into public.question_test_cases (question_id, inputs, outputs) values (2, '{"strings": "this is not a palindrome."}', '{"strings": false}');
insert into public.question_test_cases (question_id, inputs, outputs) values (2, '{"strings": "civic duty"}', '{"strings": false}');
insert into public.question_test_cases (question_id, inputs, outputs) values (2, '{"strings": "d civicd"}', '{"strings": true}');
insert into public.question_test_cases (question_id, inputs, outputs) values (3, '{"arrays": []}', '{"arrays": true}');
insert into public.question_test_cases (question_id, inputs, outputs) values (3, '{"arrays": [1]}', '{"arrays": true}');
insert into public.question_test_cases (question_id, inputs, outputs) values (3, '{"arrays": [1, 2, 3, 4, 5]}', '{"arrays": true}');
insert into public.question_test_cases (question_id, inputs, outputs) values (3, '{"arrays": [2, 2, 2, 2]}', '{"arrays": true}');
insert into public.question_test_cases (question_id, inputs, outputs) values (3, '{"arrays": [5, 4, 3, 2, 1, 2]}', '{"arrays": false}');
insert into public.question_test_cases (question_id, inputs, outputs) values (3, '{"arrays": [2, 1, 2, 3, 4, 5]}', '{"arrays": false}');
insert into public.question_test_cases (question_id, inputs, outputs) values (3, '{"arrays": [5, 4, 3, 2, 1]}', '{"arrays": true}');
insert into public.question_test_cases (question_id, inputs, outputs) values (3, '{"arrays": [-5, -4, -3, -2, -1]}', '{"arrays": true}');
insert into public.question_test_cases (question_id, inputs, outputs) values (3, '{"arrays": [1, 0, -1, -2, -3, -4, -5]}', '{"arrays": true}');
insert into public.question_test_cases (question_id, inputs, outputs) values (3, '{"arrays": [1, 1, 1, 2]}', '{"arrays": true}');
insert into public.question_test_cases (question_id, inputs, outputs) values (3, '{"arrays": [5, 5, 5, 2, 1]}', '{"arrays": true}');
insert into public.question_test_cases (question_id, inputs, outputs) values (3, '{"arrays": [1, 2, 5, 2, 1]}', '{"arrays": false}');
