-- +goose Up
-- +goose StatementBegin

create or REPLACE FUNCTION getK()
RETURNS int 
   LANGUAGE plpgsql
  as
$$

BEGIN
 return (select value from constants where id = 'k');
END;
$$;
--------------------------------------------------------------------------------
CREATE or REPLACE FUNCTION createEmployeePhoto(b64 text)
   RETURNS int 
   LANGUAGE plpgsql
  as
$$
DECLARE 
 idp int;
BEGIN
	insert into employee_photos (photob64)
    values (b64) RETURNING id into idp;
 return idp;
END;
$$;
----------------------------------------------------------------------------------
create or REPLACE FUNCTION getJ()
RETURNS int 
   LANGUAGE plpgsql
  as
$$

BEGIN
 return (select value from constants where id = 'j');
END;
$$;
-----------------------------------------------------------------------------------
CREATE or REPLACE FUNCTION createEmployeePhoto(b64 text)
   RETURNS int 
   LANGUAGE plpgsql
  as
$$
DECLARE 
 idp int;
BEGIN
	insert into employee_photos (photob64)
    values (b64) RETURNING id into idp;
 return idp;
END;
$$;
-----------------------------------------------------------------------------------------

CREATE OR REPLACE PROCEDURE createEmployee(
    name varchar(255),
    email varchar(255),
    skill int,
    login varchar(255),
    password varchar(70),
    photob64 text
)

language plpgsql
as $$
DECLARE id_photo int;
begin
	
	id_photo := createEmployeePhoto(photob64);
	
	INSERT INTO employees (name,email,skill,login,password,verified,photo_id,correct_answers,wrong_answers,score_answer,score_answer_max) VALUES (
    name,
    email,
    skill,
    login,
    crypt(password, gen_salt('bf')),
    false,
	id_photo,
    0,
    0,
    0,
    0);
    
    --commit;
end;
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS createEmployee;
DROP FUNCTION IF EXISTS createEmployeePhoto;
DROP FUNCTION IF EXISTS getK;
DROP FUNCTION IF EXISTS getJ;
-- +goose StatementEnd
