-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE PROCEDURE createViolationPool(
    violation_id varchar(255),
    transport text,
    camera_id varchar(255),
    violation_value text,
    skill int,
    datetime timestamp,
)

language plpgsql
as $$
begin
	
	INSERT INTO violations (violation_id,transport,camera_id,violation_value,datetime,solved) VALUES (
    violation_id,
    transport,
    camera_id,
    violation_value,
    skill,
    datetime,
    false
    );

    INSERT INTO violations_employees_pool(violation_id, employee_id)
    select violations.id, employees.id
    from violations
    join employees on employees.skill = violations.skill;
    
end;
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS createViolationPool;
-- +goose StatementEnd
