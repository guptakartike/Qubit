create table users (
    id UUID primary key,
    name varchar(100) not null,
    email varchar(255) not null unique,
    status varchar(20) not null default 'active',
    email_verified_at timestamptz,
    deleted_at timestamptz,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),

    constraint user_status_check 
    check (status in ('active', 'suspended'))
);