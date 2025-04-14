DO
$$
BEGIN
   IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'mooddb') THEN
      CREATE DATABASE mooddb;
   END IF;
END
$$;

CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   name TEXT NOT NULL,
   email TEXT NOT NULL UNIQUE,
   password TEXT NOT NULL,
   role TEXT NOT NULL CHECK (role IN ('paciente', 'psicologo'))
);

CREATE TABLE mood_entries (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  level INTEGER NOT NULL CHECK (level BETWEEN 1 AND 5),
  timestamp TIMESTAMP NOT NULL DEFAULT NOW()
);
