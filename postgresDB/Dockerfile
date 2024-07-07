# Use the official PostgreSQL image
FROM postgres:13

# Set environment variables
ENV POSTGRES_USER=${DB_USER}
ENV POSTGRES_PASSWORD=${DB_PASSWORD}
ENV POSTGRES_DB=${DB_NAME}

# Copy the SQL script into the container
COPY users.sql /docker-entrypoint-initdb.d/
