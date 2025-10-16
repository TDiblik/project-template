-- Add plugins
create extension if not exists "uuid-ossp";

-- create domains
create domain theme_possibilities as text
    check (value in ('light', 'dark'));
create domain translations_possibilities as text
    check (value in ('cs', 'en'));

-- create general functions
create or replace function automatically_update_updated_at_column()
returns trigger as $$
begin
    NEW.updated_at = CURRENT_TIMESTAMP;
    return NEW;
end;
$$ language plpgsql;

create or replace function prohibit_update()
returns trigger as $$
begin
    raise exception 'Updates are not allowed on this table';
end;
$$ language plpgsql;

create or replace function prohibit_insert()
returns trigger as $$
begin
    raise exception 'Inserts are not allowed on this table';
end;
$$ language plpgsql;