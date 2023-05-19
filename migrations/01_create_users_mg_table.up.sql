create table if not exists interns (
    id_intern serial primary key,
    name_intern varchar(64),
    surname_intern varchar(64),
    email varchar(64),
    password varchar(64)
);

create table if not exists candidates (
    id_candidate serial primary key,
    name_candidate varchar(64),
    surname_candidate varchar(64),
    email varchar(64),
    password varchar(64)
);

create table if not exists curators (
    id_curator serial primary key,
    name_curator varchar(64),
    surname_curator varchar(64),
    email varchar(64),
    password varchar(64)
);

create table if not exists mentors (
    id_mentor SERIAL PRIMARY KEY,
    name_curator varchar(64),
    surname_curator varchar(64)
);

create table if not exists hrs (
    id_hr serial primary key,
    name_hr varchar(64),
    surname_hr varchar(64)
);