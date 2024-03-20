-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pgcrypto;

create table if not exists employee_photos(
    id serial primary key,
    photob64 text not null
);

create table IF NOT EXISTS employees (
id serial primary key,
name varchar(255) not null,
email varchar(255) not null,
skill int not null,
login varchar(255) not null UNIQUE,
password text NOT NULL,
verified boolean not null,
photo_id int not null,
correct_answers int not null,
wrong_answers int not null,
score_answer int not null,
score_answer_max int not null,
FOREIGN KEY (photo_id) REFERENCES employee_photos (id) ON DELETE CASCADE
);

create table if not exists masters(
id serial primary key,
name varchar(255) not null,
email varchar(255) not null,
login varchar(255) not null UNIQUE,
password text NOT NULL
);

create table if not exists cameras(
   id varchar(255) primary key,
   type varchar(255) not null,
   coordinateX real not null,
   coordinateY real not null,
   description text 
);

create table if not exists excesses(
id serial primary key,
transport text not null,
camera_id varchar(255) not null,
violation_id varchar(255) not null,
violation_value text not null,
skill int not null,
datetime timestamp not null,
photo text not null,
isViolation boolean,
FOREIGN KEY (camera_id) REFERENCES cameras (id) ON DELETE CASCADE
);

create table if not exists excesses_employees_pool(
    id serial primary key,
    excess_id int not null,
    employee_id int not null,
    isViolation boolean,
    FOREIGN KEY (excess_id) REFERENCES excesses (id) ON DELETE CASCADE,
    FOREIGN KEY (employee_id) REFERENCES employees (id) ON DELETE CASCADE
);

create table if not exists analitic(
    id serial primary key,
    employee_id int not null,
    is_correct boolean not null,
    date timestamp not null,
    FOREIGN KEY (employee_id) REFERENCES employees (id) ON DELETE CASCADE
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS analitic;
DROP TABLE IF EXISTS excesses_employees_pool;
DROP TABLE IF EXISTS employees;
DROP TABLE IF EXISTS masters;
DROP TABLE IF EXISTS employee_photos;
DROP TABLE IF EXISTS excesses;
DROP TABLE IF EXISTS cameras;
DROP EXTENSION IF EXISTS pgcrypto;
-- +goose StatementEnd

