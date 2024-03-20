-- +goose Up
-- +goose StatementBegin
create or REPLACE FUNCTION get_lern_with_interval(page int,size int,skilll int,type_in varchar(255),start_date timestamp, end_date timestamp)
RETURNS TABLE(id int,transport text,camera_id varchar(255),violation_id varchar(255), violation_value text, type varchar(255)) AS $$
BEGIN
    RETURN QUERY select e.id,e.transport,e.camera_id,e.violation_id,e.violation_value,c.type  from excesses e
                join cameras c ON c.id = e.camera_id
                where e.datetime > start_date and e.datetime < end_date and e.skill = skilll and c.type = type_in and isViolation is not null
                order by e.datetime
                limit size offset (page - 1) * size;
                
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS get_lern_with_interval;
-- +goose StatementEnd
