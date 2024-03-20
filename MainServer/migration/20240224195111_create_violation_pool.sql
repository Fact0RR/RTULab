-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE PROCEDURE createExcessesPool(
    transport text,
    camera_id varchar(255),
    violation_id varchar(255),
    violation_value text,
    skill int,
    datetime timestamp,
    photo text
)

language plpgsql
as $$
DECLARE
   excess_id int;
begin
	
	INSERT INTO excesses (transport,camera_id,violation_id,violation_value,skill,datetime,photo) VALUES (
    transport,
    camera_id,
    violation_id,
    violation_value,
    skill,
    datetime,
    photo
    ) RETURNING id into excess_id;

    INSERT INTO excesses_employees_pool(excess_id, employee_id)
    select excesses.id, employees.id
    from excesses
    join employees on employees.skill = excesses.skill
    where excesses.id = excess_id;
    
end;
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS createExcessesPool;
-- +goose StatementEnd
