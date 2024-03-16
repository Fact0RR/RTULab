-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION check_consensus_func() RETURNS TRIGGER AS $my_table$
   BEGIN
   
    IF isAllVoited(NEW."excess_id") THEN
        IF isAllUnanimouslyVoitedTrue(NEW."excess_id") THEN

            call changeScoreEmployees(NEW."excess_id", true);
            call deleteFromPool(NEW."excess_id");
            call setSolveInViolation(NEW."excess_id", true);

        ELSIF  not isAllUnanimouslyVoitedTrue(NEW."excess_id") THEN

            call changeScoreEmployees(NEW."excess_id", false);
            call deleteFromPool(NEW."excess_id");
            call setSolveInViolation(NEW."excess_id", false);

        ELSE

            call nextLevelSkill(NEW."excess_id");

        END IF;
    --ELSE
        --insert into test values (NEW."excess_id",0);
    END IF;
    RETURN NEW;
   END;
$my_table$ LANGUAGE plpgsql;

CREATE TRIGGER check_consensus
    AFTER UPDATE OF isViolation ON excesses_employees_pool
    FOR EACH ROW
    EXECUTE PROCEDURE check_consensus_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS check_consensus_func CASCADE;
-- +goose StatementEnd