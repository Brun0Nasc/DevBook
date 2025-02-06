USE devbook;

CREATE USER 'golang'@'localhost' IDENTIFIED BY 'dev123456';

CREATE TABLE users(
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    nickname VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    pass VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
) ENGINE=INNODB;

CREATE TABLE followers(
    user_id INT NOT NULL,
    follower_id INT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY(follower_id) REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY(user_id, follower_id)
) ENGINE=INNODB;

CREATE TABLE posts(
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    content VARCHAR(300) NOT NULL,
    author_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
    likes INT DEFAULT 0,
    FOREIGN KEY(author_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=INNODB;

INSERT INTO users (username, nickname, email, pass) VALUES ('user1', 'nick1', 'user1@mail.com', '$2a$10$NS.dlHs4SfUQoPg9cHfX9uayR8NcgI.eLDsckvTF67OvdP6bGPEAq'),
                                                           ('user2', 'nick2', 'user2@mail.com', '$2a$10$NS.dlHs4SfUQoPg9cHfX9uayR8NcgI.eLDsckvTF67OvdP6bGPEAq'),
                                                           ('user3', 'nick3', 'user3@mail.com', '$2a$10$NS.dlHs4SfUQoPg9cHfX9uayR8NcgI.eLDsckvTF67OvdP6bGPEAq'),
                                                           ('user4', 'nick4', 'user4@mail.com', '$2a$10$NS.dlHs4SfUQoPg9cHfX9uayR8NcgI.eLDsckvTF67OvdP6bGPEAq');

INSERT INTO followers (user_id, follower_id) VALUES (1, 2),
                                                    (1, 3),
                                                    (2, 1),
                                                    (3, 1),
                                                    (4, 1);

GRANT SELECT, INSERT, UPDATE, DELETE ON devbook.users TO 'golang'@'localhost';
GRANT SELECT, INSERT, UPDATE, DELETE ON devbook.followers TO 'golang'@'localhost';

FLUSH PRIVILEGES;