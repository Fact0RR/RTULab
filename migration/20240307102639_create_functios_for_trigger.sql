-- +goose Up
-- +goose StatementBegin
create or REPLACE FUNCTION isAllUnanimouslyVoitedTrue(v_id varchar(255))
RETURNS boolean 
   LANGUAGE plpgsql
  as
$$

BEGIN
RETURN (select isViolation from (select vp.violation_id,vp.isViolation from violations_employees_pool vp
		join violations v on v.id = vp.violation_id
		join employees em on em.id = vp.employee_id
		where em.skill = v.skill) as s
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
res =  (select sum(c) as s from (select violation_id,isViolation , count(isViolation) as c from 
		(select vp.violation_id,vp.isViolation from violations_employees_pool vp
		join violations v on v.id = vp.violation_id
		join employees em on em.id = vp.employee_id
		where em.skill = v.skill) as s
    group by violation_id,isViolation
    having count(isViolation) > 0) as foo
    group by violation_id
having violation_id = v_id);

RETURN res = getK();
END;
$$;
--------------------------------------------------------------------------------
CREATE OR REPLACE PROCEDURE changeScoreEmployees(v_id varchar(255), trueAnswer boolean)
LANGUAGE plpgsql
AS $$
begin

UPDATE employees emp
SET correct_answers = case when trueAnswer = vi.isViolation 
    THEN emp.correct_answers+1
    else emp.correct_answers
end,
wrong_answers = case when not trueAnswer = vi.isViolation 
    THEN emp.wrong_answers +1
    else emp.wrong_answers 
end,
score_answer = case 
    when trueAnswer = vi.isViolation
    THEN emp.score_answer+1
    when trueAnswer = not vi.isViolation
    THEN 0
    else emp.score_answer
   
end,
score_answer_max = case when emp.score_answer>emp.score_answer_max
    THEN emp.score_answer
    ELSE emp.score_answer_max
end
FROM violations_employees_pool vi
where vi.employee_id = emp.id;

end;
$$;
--------------------------------------------------------------------------------
CREATE OR REPLACE PROCEDURE deleteFromPool(v_id varchar(255))
LANGUAGE plpgsql
AS $$
begin
    delete from violations_employees_pool where violation_id = v_id;
end;
$$;
--------------------------------------------------------------------------------
CREATE OR REPLACE PROCEDURE setSolveInViolation(v_id varchar(255),answer boolean)
LANGUAGE plpgsql
AS $$
begin
    update violations
    set isViolation = answer
    where id = v_id;
end;
$$;
---------------------------------------------------------------------------------
CREATE OR REPLACE PROCEDURE nextLevelSkill(v_id varchar(255))
LANGUAGE plpgsql
AS $$
begin
    update violations
    set skill = skill+1
    where id = v_id;

    INSERT INTO violations_employees_pool(violation_id, employee_id)
    select violations.id, employees.id
    from violations
    join employees on employees.skill = violations.skill
    where violations.id = v_id;
end;
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS isAllVoited;
DROP FUNCTION IF EXISTS isAllUnanimouslyVoitedTrue;
DROP PROCEDURE IF EXISTS changeScoreEmployees;
DROP PROCEDURE IF EXISTS deleteFromPool;
DROP PROCEDURE IF EXISTS setSolveInViolation;
DROP PROCEDURE IF EXISTS nextLevelSkill;
-- +goose StatementEnd
