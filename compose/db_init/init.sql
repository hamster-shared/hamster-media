CREATE TABLE IF NOT EXISTS t_middleware_ip_record (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    ip varchar(64) NOT NULL UNIQUE,
    email varchar(64) DEFAULT NULL,
    social_platform varchar(64) DEFAULT NULL,
    social_account varchar(64) DEFAULT NULL
    );
