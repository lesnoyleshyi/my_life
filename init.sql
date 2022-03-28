CREATE TABLE IF NOT EXISTS lists (
    id              integer PRIMARY KEY,
    emoji           varchar,
    title           varchar,
    order           integer,
    relevance_time  timestamp
                                 );
