-- USERS
CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     email TEXT NOT NULL UNIQUE,
                                     password_hash TEXT NOT NULL,
                                     salt TEXT NOT NULL,
                                     wmk_pin TEXT NOT NULL,
                                     wmk_recovery TEXT NOT NULL,
                                     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                     sync_cursor INT NOT NULL DEFAULT 0
);

-- ENCRYPTED DATA
CREATE TABLE IF NOT EXISTS encrypted_data (
                                              id SERIAL PRIMARY KEY,
                                              cipher_text TEXT NOT NULL,
                                              nonce TEXT NOT NULL,
                                              auth_tag TEXT NOT NULL,
                                              created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                              updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- PROJECTS
CREATE TABLE IF NOT EXISTS projects (
                                        id SERIAL PRIMARY KEY,
                                        user_id INT NOT NULL,
                                        encrypted_data_id INT NOT NULL,
                                        sync_sequence INT NOT NULL DEFAULT 0,
                                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                        deleted_at TIMESTAMP NULL,

                                        CONSTRAINT fk_projects_user
                                        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,

    CONSTRAINT fk_projects_encrypted_data
    FOREIGN KEY (encrypted_data_id) REFERENCES encrypted_data(id) ON DELETE CASCADE
    );

-- NOTES
CREATE TABLE IF NOT EXISTS notes (
                                     id SERIAL PRIMARY KEY,
                                     user_id INT NOT NULL,
                                     project_id INT NULL,
                                     encrypted_data_id INT NOT NULL,
                                     version INT NOT NULL DEFAULT 1,
                                     sync_sequence INT NOT NULL DEFAULT 0,
                                     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                     deleted_at TIMESTAMP NULL,

                                     CONSTRAINT fk_notes_user
                                     FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,

    CONSTRAINT fk_notes_project
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE SET NULL,

    CONSTRAINT fk_notes_encrypted_data
    FOREIGN KEY (encrypted_data_id) REFERENCES encrypted_data(id) ON DELETE CASCADE
    );

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