CREATE TABLE IF NOT EXISTS t_contact_us (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    contact_type varchar(100) NOT NULL comment '联系类型',
    contact_name varchar(100) NOT NULL comment '联系方名字',
    contact_email_address varchar(100) NOT NULL comment '联系邮件地址',
    contact_platform varchar(100) NULL comment '联系平台',
    contact_information varchar(100) NULL comment '联系方式',
    topic text NULL comment '话题',
    email_seed_flag int NOT NULL comment '是否发送邮件',
    create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间'
    );