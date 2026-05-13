-- USERS
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    sync_cursor BIGINT NOT NULL DEFAULT 0
);

-- USER KEYS
CREATE TABLE IF NOT EXISTS user_keys (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    key_type TEXT NOT NULL,
    salt BYTEA NOT NULL,
    wrapped_master_key BYTEA NOT NULL,
    updated_at TIMESTAMP NOT NULL,

    CONSTRAINT fk_user_keys_user 
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,

    CONSTRAINT unique_user_key_type
    UNIQUE (user_id, key_type)
);

-- PROJECTS
CREATE TABLE IF NOT EXISTS projects (
    id UUID PRIMARY KEY,
    user_id INT NOT NULL,
    encrypted_content BYTEA NOT NULL,

    version INT NOT NULL DEFAULT 1,
    sync_sequence BIGINT NOT NULL DEFAULT 0,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,

    CONSTRAINT fk_projects_user
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_projects_user_id ON projects(user_id);
CREATE INDEX idx_projects_sync ON projects(user_id, sync_sequence);

-- NOTES
CREATE TABLE IF NOT EXISTS notes (
    id UUID PRIMARY KEY,
    user_id INT NOT NULL,
    project_id UUID NULL,

    encrypted_title BYTEA NOT NULL,
    encrypted_content BYTEA NOT NULL,

    version INT NOT NULL DEFAULT 1,
    sync_sequence BIGINT NOT NULL DEFAULT 0,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,

    CONSTRAINT fk_notes_user
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,

    CONSTRAINT fk_notes_project
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE SET NULL
);

CREATE INDEX idx_notes_user_id ON notes(user_id);
CREATE INDEX idx_notes_sync ON notes(user_id, sync_sequence);

-- REFRESH TOKENS
CREATE TABLE IF NOT EXISTS refresh_tokens (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,

    token_hash TEXT NOT NULL,

    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    revoked_at TIMESTAMP NULL,

    CONSTRAINT fk_refresh_tokens_user
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_refresh_tokens_user_id ON refresh_tokens(user_id);

-- EMAIL VERIFICATIONS
CREATE TABLE IF NOT EXISTS email_verifications (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    code_hash TEXT NOT NULL,

    expires_at TIMESTAMP NOT NULL,
    used_at TIMESTAMP NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_email_verifications_email ON email_verifications(email);
