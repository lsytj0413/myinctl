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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='User Accounts Information' MAX_ROWS=10 MIN_ROWS=1 KEY_BLOCK_SIZE=8 /*PARTITION by KEY(`id`) PARTITIONS 10*/;

insert into user_accounts(id, name, gender, gender2, email) 
    values
    (1, "name1", "Male", "Male", "email1"),
    (2, "name2", "Male", "Male", "email2"),
    (3, "name3", "Female", "Female", "email3");