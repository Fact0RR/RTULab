-- +goose Up
-- +goose StatementBegin
INSERT INTO employee_photos (photob64) values(
	'sdasdasdada'
);

INSERT INTO employees (name,email,skill,login,password,verified,photo_id,correct_answers,wrong_answers,score_answer) VALUES (
    'John',
    'johndoe@mail.com',
    1,
    'loginJohn',
    crypt('johnspassword', gen_salt('bf')),
    false,
	1,
    0,
    0,
    0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM employees
WHERE email='johndoe@mail.com';
DELETE FROM employee_photos
WHERE photob64='sdasdasdada';


-- +goose StatementEnd
