CREATE DATABASE maps;

CREATE TABLE IF NOT EXISTS users
(
    id        SERIAL PRIMARY KEY,
    name      VARCHAR(20),
    surname   VARCHAR(20),
    email     VARCHAR(20),
    telephone VARCHAR(12),
    password  bytea --for sha256 binary data
);

CREATE TABLE IF NOT EXISTS polygon
(
    id          VARCHAR(36) PRIMARY KEY,
    name        VARCHAR(30),
    description VARCHAR(140),
    comment     VARCHAR(140),
    user_id     INTEGER REFERENCES users (id),
--     coordinates DOUBLE PRECISION[] --array for lat and long data instead of making second table with foreign key
    coordinates jsonb
);

CREATE TABLE IF NOT EXISTS pin
(
    id          VARCHAR(36) PRIMARY KEY,
    name        VARCHAR(30),
    description VARCHAR(140),
    comment     VARCHAR(140),
    user_id     INTEGER REFERENCES users (id),
--     coordinates DOUBLE PRECISION[] --array for lat and long data instead of making second table with foreign key
    coordinates jsonb
);


CREATE TABLE inside
(
    id      serial PRIMARY KEY,
    pinid   varchar(50) REFERENCES pin (id),
    polid   varchar(50) REFERENCES polygon (id),
    user_id INTEGER REFERENCES users (id)
);

-- TODO : write insert query to users table for real users, after project comp. this section will be written


--sql duzeltmeleri
ALTER TABLE inside
DROP CONSTRAINT IF EXISTS inside_pinid_fkey,
ADD CONSTRAINT inside_pinid_fkey
FOREIGN KEY (pinid)
REFERENCES pin (id)
ON DELETE CASCADE;

ALTER TABLE inside
DROP CONSTRAINT IF EXISTS inside_polid_fkey,
ADD CONSTRAINT inside_polid_fkey
FOREIGN KEY (polid)
REFERENCES polygon (id)
ON DELETE CASCADE;