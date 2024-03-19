-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION add_correct_answers_func() RETURNS TRIGGER AS $my_table$
   BEGIN
    IF NEW."correct_answers" >  OLD."correct_answers" THEN
    INSERT INTO analitic (employee_id,is_correct,date)  values(
        NEW."id",
        true,
        now()
    );
    END IF;
    RETURN NEW;
   END;
$my_table$ LANGUAGE plpgsql;

CREATE TRIGGER add_correct
    AFTER UPDATE OF correct_answers ON employees
    FOR EACH ROW
    EXECUTE PROCEDURE add_correct_answers_func();
----------------------------------------------------------------
CREATE OR REPLACE FUNCTION add_wrong_answers_func() RETURNS TRIGGER AS $my_table$
   BEGIN
    IF NEW."wrong_answers" > OLD."wrong_answers" THEN
    INSERT INTO analitic (employee_id,is_correct,date)  values(
        NEW."id",
        false,
        now()
    );
    END IF;
    RETURN NEW;
   END;
$my_table$ LANGUAGE plpgsql;

CREATE TRIGGER add_wrong
    AFTER UPDATE OF wrong_answers ON employees
    FOR EACH ROW
    EXECUTE PROCEDURE add_wrong_answers_func();
--------------------------------------------------------------------------------------------------
create or REPLACE FUNCTION get_employee_analitic_with_interval(emp_id int,start_date timestamp, end_date timestamp)
RETURNS TABLE(id int, employee_id int, is_correct boolean, date timestamp) AS $$
BEGIN
    RETURN QUERY select analitic.id,analitic.employee_id,analitic.is_correct, analitic.date from analitic
                where analitic.date > start_date and analitic.date < end_date and analitic.employee_id = emp_id
                order by date;
END;
$$ LANGUAGE plpgsql;
--------------------------------------------------------------------------------------------------
create or REPLACE FUNCTION check_excess_on_readiness(exc_id int)
RETURNS boolean  AS $$
BEGIN
    RETURN (SELECT count(*) from excesses where id = exc_id and isViolation = true) = 1;
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS add_correct_answers_func CASCADE;
DROP FUNCTION IF EXISTS add_wrong_answers_func CASCADE;
DROP FUNCTION IF EXISTS get_employee_analitic_with_interval;
DROP FUNCTION IF EXISTS check_excess_on_readiness;
-- +goose StatementEnd
