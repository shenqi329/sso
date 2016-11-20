use sys;

grant select on db_sso.* to user_select@`%` identified by 'user_select';

grant select,update on db_sso.* to user_update@`%` identified by 'user_update';

grant select,update,delete,insert on db_sso.* to user_connect@`%` identified by 'user_connect';

grant all privileges on db_sso.* to dba@`%` identified by 'dba';


CREATE DATABASE IF NOT EXISTS db_sso DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
use db_sso;


SHOW VARIABLES LIKE 'event_scheduler';

SET GLOBAL event_scheduler = ON;
SET GLOBAL event_scheduler = ON;


CREATE DATABASE IF NOT EXISTS db_easynote DEFAULT CHARSET utf8 COLLATE utf8_general_ci;

grant select on db_easynote.* to easynote_select@`%` identified by 'easynote_select';

grant select,update on db_easynote.* to easynote_update@`%` identified by 'easynote_update';

grant select,update,delete,insert on db_easynote.* to easynote_connect@`%` identified by 'easynote_connect';

grant all privileges on db_easynote.* to easynote_dba@`%` identified by 'easynote_dba';