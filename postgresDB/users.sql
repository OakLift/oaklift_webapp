-- Define variables for database and user information
\set dbname 'oaklift_webpage'
\set dbuser 'postgres'
\set dbpassword 'oaklift@445'

-- Connect to the default database (usually postgres)
\c postgres

-- Create the database if it doesn't exist
DO $$ 
BEGIN
    IF NOT EXISTS (
        SELECT FROM pg_database 
        WHERE datname = 'oaklift_webpage'
    ) THEN
        PERFORM dblink_exec(
            'dbname=' || current_database() || ' user=' || current_user,
            'CREATE DATABASE :' || 'oaklift_webpage'
        );
    END IF;
END $$;

-- Create a new user with the specified password if it doesn't exist
DO $$ 
BEGIN
    IF NOT EXISTS (
        SELECT FROM pg_roles 
        WHERE rolname = 'postgres'
    ) THEN
        EXECUTE format('CREATE USER %I WITH ENCRYPTED PASSWORD %L', 'postgres', 'oaklift@445');
    END IF;
END $$;

CREATE EXTENSION IF NOT EXISTS dblink;

-- Grant all privileges on the database to the new user
DO $$
BEGIN
    PERFORM dblink_exec(
        'dbname=' || 'oaklift_webpage' || ' user=' || current_user,
        format('GRANT ALL PRIVILEGES ON DATABASE %I TO %I', 'oaklift_webpage', 'postgres')
    );
END $$;

-- Connect to the new database
\c :dbname

-- Create the users table if it doesn't exist
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    oauth_provider VARCHAR(50),
    oauth_id VARCHAR(255),
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    role VARCHAR(50) NOT NULL CHECK (role IN ('mentor', 'mentee')),
    bio TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Grant all privileges on the users table to the new user
GRANT ALL PRIVILEGES ON TABLE users TO :dbuser;


-- Example additional table to store detailed information for mentors and mentees
CREATE TABLE profiles (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    phone_number VARCHAR(15),
    address TEXT,
    linkedin_profile VARCHAR(100),
    experience TEXT,             -- Relevant for mentors
    interests TEXT               -- Relevant for mentees
);

-- Grant all privileges on the users table to the new user
GRANT ALL PRIVILEGES ON TABLE profiles TO :dbuser;