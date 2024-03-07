-- +goose Up
-- +goose StatementBegin
create or REPLACE FUNCTION isAllUnanimouslyVoitedTrue(v_id varchar(255))
RETURNS boolean 
   LANGUAGE plpgsql
  as
$$

BEGIN

RETURN (select isViolation from violations_employees_pool
    group by violation_id,isViolation
    having count(isViolation) = getK() and  violation_id = v_id);
END;
$$;

create or REPLACE FUNCTION isAllVoited(v_id varchar(255))
RETURNS boolean 
   LANGUAGE plpgsql
  as
$$
DECLARE
    res int;
BEGIN
res =  (select sum(c) as s from (select violation_id,isViolation , count(isViolation) as c from violations_employees_pool
    group by violation_id,isViolation
    having count(isViolation) > 0) as foo
    group by violation_id
having violation_id = v_id);

RETURN res = getK();
END;
$$;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS isAllVoited;
DROP FUNCTION IF EXISTS isAllUnanimouslyVoitedTrue;
-- +goose StatementEnd
