create table if not exists images(
    id serial primary key,
    user_id uuid references auth.users,
    status int not null default 1,
    prompt text not null,
    deleted boolean not null default 'false',
    image_location text,
    batch_id uuid not null,
    created_at timestamp not null default now(),
    deleted_at timestamp
)
