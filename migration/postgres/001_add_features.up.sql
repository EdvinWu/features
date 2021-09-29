create table features
(
    id             uuid      not null
        constraint feature_pkey
            primary key,
    display_name   varchar,
    technical_name varchar   not null,
    expires_on     timestamp,
    description    varchar,
    inverted       bool      not null default false,
    created_at     timestamp not null,
    updated_at     timestamp,
    deleted_at     timestamp

);

create table feature_users
(
    feature_id uuid not null references features,
    user_id    uuid not null,
    constraint feature_user_id_unique unique (feature_id, user_id)
)