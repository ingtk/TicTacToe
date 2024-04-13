CREATE TABLE IF NOT EXISTS games (
    id VARCHAR(100) PRIMARY KEY,
    host_user_id VARCHAR(100) NOT NULL,
    guest_user_id VARCHAR(100) NOT NULL,
    turn VARCHAR(100),
    winner VARCHAR(100) NOT NULL,
    board TEXT NOT NULL,
    started_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL
);

SELECT * FROM games WHERE host_user_id = 'hoge' ORDER BY created_at DESC LIMIT 1;
