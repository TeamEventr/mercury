-- +goose Up
-- +goose StatemendBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE enum_account_status as ENUM ('active', 'disabled', 'banned', 'deleted');
CREATE TYPE enum_event_schedule as ENUM ('ontime', 'prepond', 'postponed', 'cancelled');
CREATE TYPE enum_event_visibility as ENUM ('draft', 'published');
CREATE TYPE enum_event_type as ENUM ('music', 'sports', 'arts', 'food', 'tech',
'business', 'others');
CREATE TYPE enum_booking_status as ENUM ('open', 'closed');
CREATE TYPE enum_hosted_status as ENUM ('under_five', 'under_twenty', 'under_fifty', 'under_hundred', 'more');
CREATE TYPE enum_gender_options as ENUM ('male', 'female', 'others');
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_account (
  id UUID NOT NULL,
  username TEXT NOT NULL,
  password_login BOOLEAN NOT NULL DEFAULT false,
  provider_id TEXT NULL,
  password TEXT NULL,
  first_name TEXT NULL,
  middle_name TEXT NULL,
  last_name TEXT NULL,
  gender enum_gender_options NULL,
  email TEXT NOT NULL,
  avatar TEXT NULL,
  city TEXT NULL,
  status enum_account_status NOT NULL DEFAULT 'active',
  loggedin_at TIMESTAMPZ NULL,
  refresh_token TEXT NULL,
  created_at TIMESTAMPZ DEFAULT NOW(),
  updated_at TIMESTAMPZ DEFAULT NOW(),
  deleted_at TIMESTAMPZ NULL,

  CONSTRAINT "user_account_pkey" PRIMARY KEY (id)
);

CREATE UNIQUE INDEX idx_user_account__username ON user_account(username);
CREATE UNIQUE INDEX idx_user_account__email ON user_account(email);
-- +goose StatementEnd

-- +goose Up
CREATE TABLE user_onboarding (
  id SERIAL NOT NULL,
  username TEXT NOT NULL,
  password TEXT NOT NULL,
  email TEXT NOT NULL,
  otp TEXT NOT NULL,
  created_at TIMESTAMPZ DEFAULT NOW(),
  expiry_at TIMESTAMPZ NOT NULL,

  CONSTRAINT "user_onboarding_pkey" PRIMARY KEY (id)
);

-- +goose Up
-- +goose StatementBegin
CREATE TABLE verification_actions (
  id SERIAL NOT NULL,
  username TEXT NOT NULL,
  purpose TEXT NOT NULL,
  otp TEXT NOT NULL,
  expiry_at TIMESTAMPZ NOT NULL,

  CONSTRAINT "verification_actions_username_fkey"
    FOREIGN KEY (username)
      REFERENCES user_account(username)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);
CREATE INDEX idx_verification__username ON verification(username);
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE host (
  id UUID NOT NULL,
  username TEXT NOT NULL,
  pan_number TEXT NULL,
  account_number TEXT NULL,
  first_name TEXT NOT NULL,
  middle_name TEXT NULL,
  last_name TEXT NOT NULL,
  phone_number TEXT NOT NULL,
  dob TEXT NOT NULL,
  company_name TEXT NOT NULL,
  company_email TEXT NOT NULL,
  backup_email TEXT NOT NULL,
  registered BOOLEAN DEFAULT false,
  registration_number TEXT NULL,
  address TEXT NULL,
  pincode INTEGER NULL,
  event_count INTEGER DEFAULT 0,
  hosted_status enum_hosted_status NOT NULL DEFAULT 'under_five',
  account_status enum_account_status NOT NULL DEFAULT 'active',
  created_at TIMESTAMPZ DEFAULT NOW(),
  updated_at TIMESTAMPZ DEFAULT NOW(),

  CONSTRAINT "host_pkey" PRIMARY KEY (id),
  CONSTRAINT "host_username_fkey"
    FOREIGN KEY (username)
      REFERENCES user_account(username)
        ON DELETE RESTRICT
        ON UPDATE CASCADE
);
CREATE UNIQUE INDEX idx_host__username ON host(username);
CREATE UNIQUE INDEX idx_host__email ON host(company_email);
CREATE INDEX idx_host__company_name ON host(company_name);
-- +goose StatementEnd

-- +goose Up
CREATE TABLE host_onboarding (
  id SERIAL NOT NULL,
  username TEXT NOT NULL,
  first_name TEXT NOT NULL,
  middle_name TEXT NULL,
  last_name TEXT NOT NULL,
  phone_number TEXT NOT NULL,
  dob TEXT NOT NULL,
  company_name TEXT NOT NULL,
  company_email TEXT NOT NULL,
  backup_email TEXT NULL,
  registered BOOLEAN NOT NULL DEFAULT false,
  registration_number TEXT NULL,
  hosted_status enum_hosted_status DEFAULT 'under_five',
  otp TEXT NOT NULL,
  created_at TIMESTAMPZ DEFAULT NOW(),
  expiry_at TIMESTAMPZ NOT NULL,

  CONSTRAINT "host_onboarding_pkey" PRIMARY KEY (id)
);

-- +goose Up
-- +goose StatementBegin
CREATE TABLE event (
  id UUID NOT NULL,
  title TEXT NOT NULL,
  type enum_event_type NOT NULL DEFAULT 'music',
  host_id UUID NOT NULL,
  description TEXT NULL,
  cover_picture_url TEXT NULL, banner_url TEXT NULL, thumbnail_url TEXT NULL,
  visibility enum_event_visibility NOT NULL DEFAULT 'draft',
  tags TEXT[],
  venue TEXT NULL,
  schedule enum_event_schedule NOT NULL DEFAULT 'ontime',
  start_time TIMESTAMPZ NULL,
  end_time TIMESTAMPZ NULL,
  age_limit INTEGER DEFAULT 18,
  created_at TIMESTAMPZ DEFAULT NOW(),
  updated_at TIMESTAMPZ DEFAULT NOW(),

  CONSTRAINT "event_pkey" PRIMARY KEY (id),
  CONSTRAINT "event_host_id_fkey"
    FOREIGN KEY (host_id)
      REFERENCES host(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE
);

CREATE INDEX idx_event_title ON event(title);
-- +goose StatementEnd

-- +goose Up
CREATE TABLE event_image (
  id SERIAL NOT NULL,
  event_id UUID NOT NULL,
  image_type TEXT NOT NULL,
  url TEXT NOT NULL,

  CONSTRAINT "event_image_pkey" PRIMARY KEY (id),
  CONSTRAINT "event_image_event_id_fkey"
    FOREIGN KEY (event_id)
      REFERENCES event(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

-- +goose Up
CREATE TABLE price_tier (
  id SERIAL NOT NULL,
  event_id UUID NOT NULL,
  name TEXT NOT NULL,
  validity_start TIMESTAMPZ NOT NULL,
  validity_end TIMESTAMPZ NOT NULL,
  price INTEGER NOT NULL,
  seat_available INTEGER NOT NULL DEFAULT 0,
  total_seat INTEGER NOT NULL DEFAULT 0,
  booking_open_time TIMESTAMPZ NULL,
  booking_close_time TIMESTAMPZ NULL,
  booking_status enum_booking_status DEFAULT 'open',

  CONSTRAINT "price_tier_pkey" PRIMARY KEY (id),
  CONSTRAINT "price_tier_event_id_fkey"
    FOREIGN KEY (event_id)
      REFERENCES event(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

-- +goose Up
CREATE TABLE ticket (
  id UUID NOT NULL,
  event_id UUID NOT NULL,
  username TEXT NOT NULL,
  tier_id INTEGER NOT NULL,
  first_name TEXT NOT NULL,
  middle_name TEXT NULL,
  last_name TEXT NOT NULL,
  phone_number TEXT NOT NULL,
  created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
  cancelled BOOLEAN DEFAULT false,
  check_in BOOLEAN NOT NULL DEFAULT false,

  CONSTRAINT "ticket_pkey" PRIMARY KEY (id),
  CONSTRAINT "ticket_username_fkey"
    FOREIGN KEY (username)
      REFERENCES user_account(username)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
  CONSTRAINT "ticket_tier_id_fkey"
    FOREIGN KEY (tier_id)
      REFERENCES price_tier(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE
);

-- +goose Up
CREATE TABLE staff (
  id TEXT NOT NULL,
  username TEXT NOT NULL,
  event_id UUID NOT NULL,
  name TEXT NOT NULL,
  created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
  shift_start TIMESTAMPZ NOT NULL,
  shift_end TIMESTAMPZ NOT NULL,
  removed BOOLEAN DEFAULT false,

  CONSTRAINT "staff_pkey" PRIMARY KEY (id),
  CONSTRAINT "staff_username_fkey"
    FOREIGN KEY (username)
      REFERENCES user_account(username)
        ON DELETE RESTRICT
        ON UPDATE CASCADE,
  CONSTRAINT "staff_event_id_fkey"
    FOREIGN KEY (event_id)
      REFERENCES event(event_id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE
);

-- +goose Up
-- +goose statementBegin
CREATE TABLE bookmark (
  id SERIAL NOT NULL,
  username TEXT NOT NULL,
  event_id UUID NOT NULL,

  CONSTRAINT "bookmark_pkey" PRIMARY KEY (id),
  CONSTRAINT "bookmark_username_fkey"
    FOREIGN KEY (username)
      REFERENCES user_account(username)
        ON DELETE RESTRICT
        ON UPDATE CASCADE,
  CONSTRAINT "bookmark_eventid_fkey"
    FOREIGN KEY (event_id)
      REFERENCES event(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE
);
CREATE UNIQUE INDEX idx_bookmark_username_event_id ON bookmark(username, event_id);
-- +goose statementEnd

-- +goose Up
-- +goose statementBegin
CREATE TABLE event_artist(
    id SERIAL NOT NULL,
    event_id UUID NOT NULL,
    username TEXT NOT NULL,
    role TEXT NOT NULL,

    CONSTRAINT "event_artist_pkey" PRIMARY KEY (id),
    CONSTRAINT "event_user_event_id_fkey"
        FOREIGN KEY (event_id)
            REFERENCES event(event_id)
                ON DELETE CASCADE
                ON UPDATE CASCADE,
    CONSTRAINT "event_user_username_fkey"
        FOREIGN KEY (username)
            REFERENCES user_account(username)
                ON DELETE CASCADE
                ON UPDATE CASCADE
);
-- +goose statementEnd
