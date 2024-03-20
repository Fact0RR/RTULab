-- +goose Up
-- +goose StatementBegin

INSERT INTO masters(name,email,login,password) VALUES(
    'JohnAdmin',
    'johnadmin@gmail.com',
    'loginAdmin',
    crypt('johnAdminspassword', gen_salt('bf'))

);

insert into cameras (id, type, coordinateX, coordinateY, description) values (
    '1adf43fe-1a25-4f0d-b93788f8adcf5d88a',
    'camerus1',
    123.333,
    64.3115,
    'камера первого типа, добавленнная миграцией'

);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM masters
where login = 'loginAdmin';
DELETE FROM cameras
where id = '1adf43fe-1a25-4f0d-b93788f8adcf5d88a'
-- +goose StatementEnd
