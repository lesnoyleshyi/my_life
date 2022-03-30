

CREATE TABLE IF NOT EXISTS users (
    id              SERIAL PRIMARY KEY
);

// unique - это значит "данные в данном столбце уникальны в рамках этой таблицы": https://www.postgrespro.ru/docs/postgresql/14/ddl-constraints#DDL-CONSTRAINTS-UNIQUE-CONSTRAINTS
// unique + not null = PRIMARY KEY


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
    UId             integer REFERENCES  users ON DELETE CASCADE ON UPDATE CASCADE,
    listId          integer REFERENCES  lists ON DELETE CASCADE ON UPDATE CASCADE,
    title           varchar,
    order_          integer,
    relevanceTime   timestamp
);

CREATE TABLE IF NOT EXISTS tasks (
  id                SERIAL  PRIMARY KEY,
  UId               integer REFERENCES  users ON DELETE CASCADE ON UPDATE CASCADE,
  listId            integer REFERENCES  lists ON DELETE CASCADE ON UPDATE CASCADE,
  sectionId         integer REFERENCES  sections ON DELETE CASCADE ON UPDATE CASCADE,
  title             varchar,
  isCompleted       boolean,
  completedDays     varchar,    -- почему в оригинале это строка, а не число?
  note              varchar,
  order_            integer,
  repeatType        string,
  daysOfWeek        string,     -- это типа "среда", "wed" и т.д.? Нам наверное, надо в enum это засунуть
  daysOfMonth       string,     -- аналогично предыдущему
  concreteDate      date,
  dateStart         date,
  dateEnd           date,
  dateReminder      date,
  relevanceTime     timestamp
);

CREATE TABLE IF NOT EXISTS subtasks (
  id                SERIAL  PRIMARY KEY,
  UId               integer REFERENCES  users ON DELETE CASCADE ON UPDATE CASCADE,
  listId            integer REFERENCES  lists ON DELETE CASCADE ON UPDATE CASCADE,
  sectionId         integer REFERENCES  sections ON DELETE CASCADE ON UPDATE CASCADE,
  taskId            integer REFERENCES  tasks ON DELETE CASCADE ON UPDATE CASCADE,
  title             varchar,
  isCompleted       boolean,
  order_            integer,
  relevanceTime     timestamp
);