FROM migrate/migrate:4.8.0

COPY *sql migrations/

# TODO: fix sslmode
CMD ["-path", "/migrations", "-database", "postgres://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable", "-verbose", "up"]
