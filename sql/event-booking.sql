USE `test`;
DROP TABLE if EXISTS user;

CREATE TABLE `user` (
    id BIGINT AUTO_INCREMENT,
    name VARCHAR(20) NOT NULL DEFAULT "",
    tel VARCHAR(15) NOT NULL DEFAULT "",
    email VARCHAR(50) NOT NULL DEFAULT "",
    password VARCHAR(100) NOT NULL DEFAULT "",
    create_at TIMESTAMP NOT NUll DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `status` TINYINT NOT NUll DEFAULT 0,
    PRIMARY key (id),
    UNIQUE key (name),
    UNIQUE key (email),
    KEY `idx_created` (`create_at`)
) ENGINE=InnoDB Default CHARSET=utf8;

DROP TABLE IF EXISTS `event`;
CREATE TABLE `event` (
    id BIGINT AUTO_INCREMENT,
    name VARCHAR(20) NOT NULL DEFAULT "",
    price INT NOT NULL DEFAULT 0,
    dates VARCHAR(30) NOT NULL DEFAULT "",
    codes VARCHAR(200) NOT NULL DEFAULT "",    
    capacity BIGINT NOT NULL DEFAULT 1000,
    `status` TINYINT NOT NULL DEFAULT 0,
    create_by BIGINT NOT NULL DEFAULT 0,
    update_by BIGINT NOT NULL DEFAULT 0,
    create_at TIMESTAMP NOT NUll DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    KEY `idx_created` (`create_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `event_booking`;
CREATE TABLE `event_booking` (
    id BIGINT AUTO_INCREMENT,
    event_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    create_at TIMESTAMP NOT NUll DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `status` TINYINT NOT NUll DEFAULT 0,
    PRIMARY KEY (id),
    UNIQUE KEY `uniq_eid_uid` (event_id, user_id),
    KEY `idx_uid` (user_id),
    KEY `idx_created` (`create_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;