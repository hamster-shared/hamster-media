CREATE TABLE IF NOT EXISTS t_subscribe_email (
   id              int           NOT NULL AUTO_INCREMENT PRIMARY KEY,
   email            varchar(100)  NOT NULL                           comment 'name',
   create_time     timestamp         NULL DEFAULT CURRENT_TIMESTAMP comment 'create_time'
);