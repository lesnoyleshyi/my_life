CREATE TABLE IF NOT EXISTS lists (
    id              SERIAL PRIMARY KEY,
    emoji           varchar,
    title           varchar,
    order_          integer,
    relevance_time  timestamp
                                 );
