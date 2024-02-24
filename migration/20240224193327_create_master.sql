-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE PROCEDURE createMaster(
    name varchar(255),
    email varchar(255),
    login varchar(255),
    password varchar(70)
)

language plpgsql
as $$
begin
	
	INSERT INTO masters (name,email,login,password) VALUES (
    name,
    email,
    login,
    password
    );
    
end;
$$;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS createMaster;
-- +goose StatementEnd
