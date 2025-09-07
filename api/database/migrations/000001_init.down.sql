-- Drop functions
drop function if exists automatically_update_updated_at_column() cascade;
drop function if exists prohibit_update() cascade;
drop function if exists prohibit_insert() cascade;

-- Optionally drop extension
drop extension if exists "uuid-ossp";