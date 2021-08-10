CREATE TABLE accounts (
    handle VARCHAR(255) PRIMARY KEY -- handle is full address {alias}@{domain.tld}
    ,alias VARCHAR(255) UNIQUE -- id is the alias from {alias}@{domain.tld}
    ,name VARCHAR(100)
    ,avatar_url VARCHAR(255) UNIQUE
    ,private_key VARCHAR(256) UNIQUE
    ,public_key VARCHAR(256) UNIQUE
    ,address VARCHAR(160) UNIQUE
    ,email VARCHAR(255) UNIQUE
    ,mobile VARCHAR(20) UNIQUE
    ,createdAt      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
