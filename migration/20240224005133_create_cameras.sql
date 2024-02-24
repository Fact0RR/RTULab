-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE PROCEDURE createCamera(
    id varchar(255) primary key,
    type varchar(255) not null,
    coordinateX real not null,
    coordinateY real not null,
    description text
)

language plpgsql
as $$
begin
	
	INSERT INTO cameras (id,type,coordinateX,coordinateY,description) VALUES (
    id,
    type,
    coordinateX,
    coordinateY,
    description
    );
    
end;
$$;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS createCamera;
-- +goose StatementEnd
