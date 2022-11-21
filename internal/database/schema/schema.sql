create table if not exists screen
(
    uuid    varchar not null
        constraint screen_pk
            primary key
        constraint screen_pk
            unique,
    url     text    not null,
    element text    not null,
    status  integer not null,
    data    BLOB
);

create index if not exists screen_status_index
    on screen (status);

