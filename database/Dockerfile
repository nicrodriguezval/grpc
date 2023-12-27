FROM postgres:latest
LABEL authors="nicro"

COPY up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]
