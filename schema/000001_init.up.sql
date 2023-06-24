CREATE TABLE users
(
    id serial not null unique,
    username varchar(255) not null,
    login varchar(255) not null unique,
    password_hash varchar(255) not null,
    Primary key (id)
);

CREATE TABLE notes_lists
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255),
    Primary key (id)
);

CREATE TABLE users_lists
(
    id_user int references users (id) on delete cascade not null, 
    id_list int references notes_lists (id) on delete cascade not null,
    Primary key (id_user, id_list)
);

CREATE TABLE notes_items
(
    id serial not null unique,
    title varchar(255),
    body varchar(255),
    Primary key (id)
);

CREATE TABLE lists_items
(
    id_list int references notes_lists (id) on delete cascade not null,
    id_item int references notes_items (id) on delete cascade not null,
    Primary key (id_list, id_item)
);
