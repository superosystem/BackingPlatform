BEGIN;

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    occupation VARCHAR(255),
    email VARCHAR(255),
    password_hash VARCHAR(255),
    avatar_file_name VARCHAR(255),
    role VARCHAR(255),
    token VARCHAR(255),
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

DROP TABLE IF EXISTS campaigns;
CREATE TABLE campaigns (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    short_description TEXT,
    description TEXT,
    perks TEXT,
    backer_count INTEGER,
    goal_amount INTEGER,
    current_amount INTEGER,
    slug VARCHAR(255),
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

DROP TABLE IF EXISTS campaign_images;
CREATE TABLE campaign_images (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    campaign_id INTEGER NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    is_primary TINYINT,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user_id INTEGER NOT NULL,
    campaign_id INTEGER NOT NULL,
    amount INTEGER,
    status VARCHAR(255),
    code VARCHAR(255),
    payment_url VARCHAR(255),
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

ALTER TABLE campaigns ADD FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE campaign_images ADD FOREIGN KEY (campaign_id) REFERENCES campaigns(id);

ALTER TABLE transactions ADD FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE transactions ADD FOREIGN KEY (campaign_id) REFERENCES campaigns(id);

COMMIT;