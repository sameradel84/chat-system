CREATE KEYSPACE IF NOT EXISTS chatsystem WITH replication = {
  'class': 'SimpleStrategy',
  'replication_factor': '1'
};

CREATE TABLE IF NOT EXISTS chatsystem.users (
    username text PRIMARY KEY,
    password text
);

CREATE TABLE IF NOT EXISTS chatsystem.messages (
    id UUID PRIMARY KEY,
    sender text,
    recipient text,
    content text,
    timestamp timestamp
);
