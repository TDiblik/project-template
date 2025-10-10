-- Drop the trigger first (must be done before dropping the table)
drop trigger if exists users_updated_at_trg on users;

-- Drop functions used by indexes
drop function if exists generate_user_full_name() cascade;

-- Drop indexes
drop index if exists users_email_index;
drop index if exists users_handle_index;
drop index if exists users_user_full_name_trg;

-- Drop the table
drop table if exists users;