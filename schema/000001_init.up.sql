CREATE TABLE users(
                      id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                      name varchar(255) NOT NULL,
                      password_hash varchar(255) NOT NULL,
                      createdAt timestamp DEFAULT CURRENT_TIMESTAMP,
                      updatedAt timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE posts(
                      id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                      userId uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                      content varchar(255) NOT NULL,
                      createdAt timestamp DEFAULT CURRENT_TIMESTAMP,
                      updatedAt timestamp DEFAULT CURRENT_TIMESTAMP
);

