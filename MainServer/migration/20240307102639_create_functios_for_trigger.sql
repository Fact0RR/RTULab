-- +goose Up
-- +goose StatementBegin
create or REPLACE FUNCTION isAllUnanimouslyVoitedTrue(v_id int)
RETURNS boolean 
   LANGUAGE plpgsql
  as
$$

BEGIN
RETURN (select isViolation from (select vp.excess_id,vp.isViolation from excesses_employees_pool vp
		join excesses v on v.id = vp.excess_id
		join employees em on em.id = vp.employee_id
		where em.skill = v.skill) as s
    group by excess_id,isViolation
    having count(isViolation) = getK() and  excess_id = v_id);
END;
$$;

create or REPLACE FUNCTION isAllVoited(v_id int)
RETURNS boolean 
   LANGUAGE plpgsql
  as
$$
DECLARE
    res int;
BEGIN
res =  (select sum(c) as s from (select excess_id,isViolation , count(isViolation) as c from 
		(select vp.excess_id,vp.isViolation from excesses_employees_pool vp
		join excesses v on v.id = vp.excess_id
		join employees em on em.id = vp.employee_id
		where em.skill = v.skill) as s
    group by excess_id,isViolation
    having count(isViolation) > 0) as foo
    group by excess_id
having excess_id = v_id);

RETURN res = getK();
END;
$$;
--------------------------------------------------------------------------------
                                                -------
CREATE OR REPLACE PROCEDURE changeScoreEmployees(v_id int, trueAnswer boolean)
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
FROM excesses_employees_pool vi
where vi.employee_id = emp.id and vi.excess_id = v_id;

end;
$$;
--------------------------------------------------------------------------------
CREATE OR REPLACE PROCEDURE deleteFromPool(v_id int)
LANGUAGE plpgsql
AS $$
begin
    delete from excesses_employees_pool where excess_id = v_id;
end;
$$;
--------------------------------------------------------------------------------
CREATE OR REPLACE PROCEDURE setSolveInViolation(v_id int,answer boolean)
LANGUAGE plpgsql
AS $$
begin
    update excesses
    set isViolation = answer
    where id = v_id;
end;
$$;
---------------------------------------------------------------------------------
CREATE OR REPLACE PROCEDURE nextLevelSkill(v_id int)
LANGUAGE plpgsql
AS $$
begin
    update excesses
    set skill = skill+1
    where id = v_id;

    INSERT INTO excesses_employees_pool(excess_id, employee_id)
    select excesses.id, employees.id
    from excesses
    join employees on employees.skill = excesses.skill
    where excesses.id = v_id;
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
