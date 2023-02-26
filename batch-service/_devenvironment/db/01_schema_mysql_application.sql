-- SQL Script to create schema for dev env
-- Just a test table
CREATE TABLE if not exists test
(
    id         bigint primary key auto_increment,
    char_data  char(5),
    vchar_data varchar(5),
    sint       smallint,
    ddata      date,
    dtdata     datetime,
    tsdata     timestamp default CURRENT_TIMESTAMP not null,
    bdata      bool,
    b2data     boolean
);

create table dispatched_group
(
    id               bigint auto_increment primary key,
    uuid_v4          char(36)                              not null,
    dispatch_status  varchar(10) default 'UNKNOWN'         not null,
    records_included int                                   not null,
    created_at       datetime    default CURRENT_TIMESTAMP not null,
    updated_at       datetime    default CURRENT_TIMESTAMP not null,
    constraint dispatched_group_uuid_v4_uindex unique (uuid_v4)
);

CREATE TABLE `groups`
(
    id              BIGINT UNSIGNED auto_increment PRIMARY KEY,
    snapshot_date   DATE                                     NOT NULL,
    profession      CHAR(20)                                 not null,
    total_salary    DECIMAL(19, 2) DEFAULT 0.00              NOT NULL,
    dispatch_status TINYINT UNSIGNED                         NOT NULL,
    created_at      datetime       default CURRENT_TIMESTAMP not null,
    updated_at      datetime       default CURRENT_TIMESTAMP not null
);

-- auto-generated definition
create table persons_v2
(
    id                  bigint auto_increment primary key,
    snapshot_date       date                                     not null,
    first_name          varchar(50)                              not null,
    last_name           varchar(50)                              not null,
    email               varchar(50)                              not null,
    profession          char(15)                                 not null,
    salary              decimal(19, 2) default 0.00              not null,
    uuid_v4             char(36)                                 not null,
    created_at          datetime       default CURRENT_TIMESTAMP not null,
    updated_at          datetime       default CURRENT_TIMESTAMP not null,
    dispatched_group_id bigint                                   not null,
    groups_id           bigint unsigned                          not null,
    constraint persons_v2_uuid_v4_uindex unique (uuid_v4)
);
create index persons_v2__snapshot_date_profession_index on persons_v2 (snapshot_date, profession);
create index persons_v2__dispatched_group_id_index on persons_v2 (dispatched_group_id);
create index persons_v2__groups_id_index on persons_v2 (groups_id);

-- auto-generated definition
create table persons
(
    id         bigint auto_increment primary key,
    first_name varchar(50)                        not null,
    last_name  varchar(50)                        not null,
    profession varchar(30)                        not null,
    email      varchar(50)                        not null,
    processed  bit      default b'0'              not null,
    created_at datetime default CURRENT_TIMESTAMP not null,
    updated_at datetime default CURRENT_TIMESTAMP not null
);
create index persons__processed_index on persons (processed);
