-- +goose Up
-- +goose StatementBegin
INSERT INTO employee_photos (photob64) values(
	'sdasdasdada'
);

INSERT INTO employees (name,email,skill,login,password,verified,photo_id,correct_answers,wrong_answers,score_answer,score_answer_max) VALUES (
    'John',
    'johndoe@mail.com',
    1,
    'loginJohn',
    crypt('johnspassword', gen_salt('bf')),
    false,
	1,
    0,
    0,
    0,
    0
);


insert into cameras (id, type, coordinateX, coordinateY, description) values (
    'id-1232132-1231231',
    'camerus1',
    123.333,
    64.3115,
    'камера добавленная автоматически в целях тестирования'

);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM employees
WHERE email='johndoe@mail.com';
DELETE FROM employee_photos
WHERE photob64='sdasdasdada';
DELETE FROM cameras
WHERE id='id-1232132-1231231';


-- +goose StatementEnd
