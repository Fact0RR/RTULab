-- +goose Up
-- +goose StatementBegin
create or REPLACE FUNCTION getCountOf10Percent()
RETURNS int 
   LANGUAGE plpgsql
  as
$$
BEGIN
 return (select count(*) from employees)/10;
END;
$$;
-----------------------------------------------------------------------
CREATE OR REPLACE PROCEDURE endOfReportingPeriod()

language plpgsql
as $$
DECLARE id_photo int;
begin
	
    update employees emp
    set skill = case  when emp.skill< 4
        then emp.skill+1
        else emp.skill
    end
    FROM (select id from employees where (correct_answers+wrong_answers)>getJ() order by (correct_answers-wrong_answers) desc limit getCountOf10Percent()) as ql
    where ql.id = emp.id;

    update employees emp
    set skill = case  when emp.skill>2
        then emp.skill-1
        else emp.skill
    end
    FROM (select id from employees where (correct_answers+wrong_answers)>getJ() order by (correct_answers-wrong_answers) asc limit getCountOf10Percent()) as ql
    where ql.id = emp.id;

    update employees 
    set correct_answers = 0,
    wrong_answers = 0,
    score_answer = 0,
    score_answer_max = 0;

end;
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS getCountOf10Percent;
-- +goose StatementEnd
