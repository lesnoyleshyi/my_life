CREATE TABLE IF NOT EXISTS users (
    id              SERIAL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS lists (
    id              SERIAL  PRIMARY KEY,             -- pimary key можно определять и на varchar. мб стоит использовать более "говорящие" названия
    UId             integer REFERENCES users ON DELETE CASCADE ON UPDATE CASCADE,
    emoji           varchar,
    title           varchar,
    order_          integer,
    relevanceTime   timestamp
                                 );

CREATE TABLE IF NOT EXISTS sections (
    id              SERIAL  PRIMARY KEY,
    UId             integer REFERENCES  users ON DELETE CASCADE,
    listId          integer REFERENCES  lists ON DELETE CASCADE,
    title           varchar,
    order_          integer,
    relevanceTime   timestamp
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
  repeatType        string,
  daysOfWeek        string,     -- это типа "среда", "wed" и т.д.? Нам наверное, надо в enum это засунуть. Почему в оригинале это массив?
  daysOfMonth       string,     -- аналогично предыдущему
  concreteDate      date,
  dateStart         date,
  dateEnd           date,
  dateReminder      date,
  relevanceTime     timestamp
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
  relevanceTime     timestamp
);