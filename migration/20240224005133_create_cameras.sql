-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE PROCEDURE createCamera(
    id varchar(255),
    type varchar(255),
    coordinateX real,
    coordinateY real,
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
