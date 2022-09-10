CREATE TABLE IF NOT EXISTS `tasks` (
    `id`            bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    `summary`       varchar(2500) NOT NULL,
    `created_by`    bigint NOT NULL,
    `performed_at`  datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_at`    datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`    datetime,
    CONSTRAINT pk_tasks_table PRIMARY KEY (id)
    )
