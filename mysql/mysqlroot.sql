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

show EVENTS;

delimiter $$  
drop event if exists e_wom_stat;  
create event e_wom_stat  
on schedule   
EVERY 1 minute  
 STARTS '2013-01-01 03:00:00'  
ON COMPLETION  PRESERVE ENABLE  
do  
begin  
  delete from t_wom_random_num where time<(CURRENT_TIMESTAMP()+INTERVAL -25 MINUTE);  
end $$  
delimiter ;  