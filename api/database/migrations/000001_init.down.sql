-- Drop functions
drop function if exists automatically_update_updated_at_column() cascade;
drop function if exists prohibit_update() cascade;
drop function if exists prohibit_insert() cascade;

-- Drop domains
drop domain if exists theme_possibilities cascade;
drop domain if exists translations_possibilities cascade;

-- Optionally drop extension
drop extension if exists "uuid-ossp";