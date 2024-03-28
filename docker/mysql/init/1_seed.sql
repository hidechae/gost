CREATE TABLE if not exists `test`
(
    `id`              bigint unsigned not null auto_increment,
    `tiny_int`        tinyint,
    `small_int`       smallint,
    `varchar_string`  varchar(255),
    `text_string`     text,
    `longtext_string` longtext,
    `created_at`      datetime        not null default current_timestamp,
    primary key (`id`)
) engine = InnoDB
  default charset = utf8mb4;