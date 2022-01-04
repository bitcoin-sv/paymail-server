CREATE TABLE accounts (
    paymail VARCHAR(255) PRIMARY KEY -- handle is full paymail {localPart@domain.tld}
    ,localPart VARCHAR(255) UNIQUE -- id is the local part from {localPart}@domain.tld
    ,name VARCHAR(100) -- name is the full legal name of the entity - person or department or organization
    ,avatar_url VARCHAR(255) UNIQUE --- avatar_url is the URL of an image file for identifying the entity
    ,private_key VARCHAR(256) UNIQUE
    ,public_key VARCHAR(256) UNIQUE
    ,address VARCHAR(160) UNIQUE
    ,email VARCHAR(255) UNIQUE
    ,mobile VARCHAR(20) UNIQUE
    ,createdAt      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
