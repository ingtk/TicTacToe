CREATE TABLE games (
    id VARCHAR(100) PRIMARY KEY,
    host_user_id VARCHAR(100) NOT NULL,
    guest_user_id VARCHAR(100) NOT NULL,
    turn VARCHAR(100),
    winner VARCHAR(100) NOT NULL,
    board TEXT NOT NULL,
    started_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);

