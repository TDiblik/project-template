create table users (
    id uuid not null constraint "users_id_pk" primary key default uuid_generate_v4(),

    email varchar(255) unique not null,
    email_verified boolean default false,
    password_hash varchar(60),

    first_name varchar(50),
    last_name varchar(50),

    handle varchar(50) unique,
    avatar_url text,
    active boolean default true,
    
    github_id text,
    github_handle text,
    github_url text,
    
    google_id text,
    
    facebook_id text,
    facebook_url text,
    
    spotify_id text,
    spotify_url text,

    last_login_at timestamp,

    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null
);
create index if not exists users_email_index ON users(email);
create index if not exists users_handle_index ON users(handle);
create or replace trigger users_updated_at_trg 
    before update on users
    for each row execute function automatically_update_updated_at_column();