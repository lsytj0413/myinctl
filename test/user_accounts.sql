CREATE TABLE `user_accounts` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(32) NOT NULL DEFAULT '' COMMENT 'User Name',
    `gender` ENUM('UNKNOWN', 'Male', 'Female') DEFAULT 'UNKNOWN' COMMENT 'User Gender',
    `gender2` ENUM('Male', 'Female') DEFAULT 'Male',
    `email` VARCHAR(20) NOT NULL DEFAULT '' COMMENT 'User Email',
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`) COMMENT 'Primary Key id',
    UNIQUE INDEX idx_user_email_name(`email`, `name`) COMMENT 'UNIQUE INDEX email & name'
) ENGINE=InnoDB ROW_FORMAT=Compact DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='User Accounts Information' MAX_ROWS=10 MIN_ROWS=1 /* KEY_BLOCK_SIZE=16 */ /*PARTITION by KEY(`id`) PARTITIONS 10*/;
-- If we use the ROW_FORMAT option, the KEY_BLOCK_SIZE must not be used.
-- And current we only support COMPACT row format.

insert into user_accounts(id, name, gender, gender2, email) 
    values
    (1, "name1", "Male", "Male", "email1"),
    (2, "name2", "Male", "Male", "email2"),
    (3, "name3", "Female", "Female", "email3");

delimiter $$$
create procedure batch_insert_accounts()
begin
declare i bigint default 0;
declare user_name VARCHAR(32);
declare user_email VARCHAR(20);
set i = 10;
start transaction;
while i < 80000 do
    set user_name=CONCAT('name_', i);
    set user_email=CONCAT('email_', i);
    insert into user_accounts(id, name, gender, gender2, email)
        values
        (i, user_name, "Male", "Male", user_email);
    set i = i+1;
end while;
commit;
end$$$
delimiter ;

call batch_insert_accounts();
drop procedure batch_insert_accounts;