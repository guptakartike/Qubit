create table password_credentials (
    user_id UUID primary key,
    password_hash text not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    
    constraint password_credentials_user_fk
    foreign key(user_id)
    references users(id)
    on delete cascade
);