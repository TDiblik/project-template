create table users (
    id uuid not null constraint "users_id_pk" primary key default uuid_generate_v4(),

    email varchar(255) unique not null,
    email_verified boolean default false,
    password_hash varchar(60),

    first_name varchar(50),
    last_name varchar(50),
    handle varchar(50) unique,
    full_name varchar(101) not null,
    initials varchar(2) not null,

    avatar_url text,
    active boolean default true,
    
    prefered_theme theme_possibilities,
    prefered_language translations_possibilities,
    
    github_id text,
    github_email text,
    github_handle text,
    github_url text,
    
    google_id text,
    google_email text,
    
    facebook_id text,
    facebook_email text,
    facebook_url text,
    
    spotify_id text,
    spotify_email text,
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

create or replace function generate_user_full_name()
returns trigger as $$
begin
    if new.first_name is not null and new.last_name is not null then
        new.full_name := initcap(new.first_name) || ' ' || initcap(new.last_name);
    else
        new.full_name := new.handle;
    end if;
    return new;
end;
$$ language plpgsql;
create trigger users_user_full_name_trg
    before insert or update on users
    for each row execute function generate_user_full_name();

create or replace function generate_user_initials()
returns trigger as $$
begin
    if new.first_name is not null and new.last_name is not null then
        new.initials := upper(left(new.first_name, 1) || left(new.last_name, 1));
    elsif new.handle is not null then
        new.initials := upper(left(new.handle, 2));
    else
        new.initials := null;
    end if;
    return new;
end;
$$ language plpgsql;
create trigger users_user_initials_trg
    before insert or update on users
    for each row execute function generate_user_initials();