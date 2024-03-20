-- +goose Up
-- +goose StatementBegin
call createEmployee('Denis','denis@mail.ru',2,'denislogin','pasrwes_my','da123');
call createEmployee('Pash','ash@mail.ru',2,'passh','pass_m6y','asdowe');
call createEmployee('Roma','oma@mail.ru',2,'romalog','pass_my','io43we');
call createEmployee('Irina','rinas@mail.ru',1,'iri','pa543ss_my','io87qpowe');
call createEmployee('Rufus','rufus@mail.ru',1,'rrr','pas3s_my','iojoij23we');
call createEmployee('Gordon','freeman@mail.ru',1,'gabengdehl3','pass_m534y','i86eqpowe');
call createEmployee('Fry','yyyy@mail.ru',1,'fry','pas3412s_my','i9owe');
call createEmployee('Vova','vovov@mail.ru',3,'vova','pas1s_my','iojosa53qpowe');
call createEmployee('Lexa','axel@mail.ru',3,'lexa','pa3ss_my','iojh5powe');
call createEmployee('Lex','leeex@mail.ru',3,'lexlogin','passjhg_my','iojoij53qpowe');
call createEmployee('Egor','egoeeor@mail.ru',3,'egrlogin','parewss_my','i234jeqpowe');
call createEmployee('Ryden','mgr@mail.ru',3,'mgrlogin','pasew_my','iojoije5233we');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM employees
WHERE name = 'Denis' or name ='Pash' or name ='Roma' or name = 'Irina' or name = 'Rufus' or name = 'Gordon'
 or name = 'Fry' or name = 'Vova' or name = 'Lexa' or name ='Lex' or name ='Egor' or name = 'Ryden';
-- +goose StatementEnd