CREATE TABLE IF NOT EXISTS users (
    id              SERIAL  PRIMARY KEY,
    username        varchar NOT NULL UNIQUE,
    phone           varchar,
    email           varchar,
    passwdHash      varchar NOT NULL,
    relevanceTime   timestamptz DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS lists (
    id              SERIAL  PRIMARY KEY,
    UId             integer REFERENCES users ON DELETE CASCADE,
    emoji           varchar,
    title           varchar,
    order_          integer,
    relevanceTime   timestamptz  DEFAULT CURRENT_TIMESTAMP
                                 );

CREATE TABLE IF NOT EXISTS sections (
    id              SERIAL  PRIMARY KEY,
    UId             integer REFERENCES  users ON DELETE CASCADE,
    listId          integer REFERENCES  lists ON DELETE CASCADE,
    title           varchar,
    order_          integer,
    relevanceTime   timestamptz  DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tasks (
  id                SERIAL  PRIMARY KEY,
  UId               integer REFERENCES  users ON DELETE CASCADE,
  listId            integer REFERENCES  lists ON DELETE CASCADE,
  sectionId         integer REFERENCES  sections ON DELETE CASCADE,
  title             varchar,
  isCompleted       boolean,
  completedDays     varchar,    -- почему в оригинале это массив строк, а не число?
  note              varchar,
  order_            integer,
  repeatType        varchar,
  daysOfWeek        varchar,     -- это типа "среда", "wed" и т.д.? Нам наверное, надо в enum это засунуть. Почему в оригинале это массив?
  daysOfMonth       varchar,     -- аналогично предыдущему
  concreteDate      date,
  dateStart         date,
  dateEnd           date,
  dateReminder      date,
  relevanceTime     timestamptz  DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS subtasks (
  id                SERIAL  PRIMARY KEY,
  UId               integer REFERENCES  users ON DELETE CASCADE,
  listId            integer REFERENCES  lists ON DELETE CASCADE,
  sectionId         integer REFERENCES  sections ON DELETE CASCADE,
  taskId            integer REFERENCES  tasks ON DELETE CASCADE,
  title             varchar,
  isCompleted       boolean,
  order_            integer,
  relevanceTime     timestamptz  DEFAULT CURRENT_TIMESTAMP
);