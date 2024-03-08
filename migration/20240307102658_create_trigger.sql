-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION check_consensus_func() RETURNS TRIGGER AS $my_table$
   BEGIN
   
    IF isAllVoited(NEW."violation_id") THEN
        IF isAllUnanimouslyVoitedTrue(NEW."violation_id") THEN

            call changeScoreEmployees(NEW."violation_id", true);
            call deleteFromPool(NEW."violation_id");
            call setSolveInViolation(NEW."violation_id", true);

        ELSIF  not isAllUnanimouslyVoitedTrue(NEW."violation_id") THEN

            call changeScoreEmployees(NEW."violation_id", false);
            call deleteFromPool(NEW."violation_id");
            call setSolveInViolation(NEW."violation_id", false);

        ELSE

            call nextLevelSkill(NEW."violation_id");

        END IF;
    --ELSE
        --insert into test values (NEW."violation_id",0);
    END IF;
    RETURN NEW;
   END;
$my_table$ LANGUAGE plpgsql;

CREATE TRIGGER check_consensus
    AFTER UPDATE OF isViolation ON violations_employees_pool
    FOR EACH ROW
    EXECUTE PROCEDURE check_consensus_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS check_consensus_func CASCADE;
-- +goose StatementEnd