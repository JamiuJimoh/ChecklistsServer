-- run `psql -U jamiu -f setup.sql -d checklist `
drop table checklists cascade;
drop table items;

create table checklists(
   id serial primary key,
   title varchar(250),
   created_at timestamp
);

insert into checklists (title, created_at)
values ('Football', '2024-05-31'),
       ('Bed', '2024-05-01');

create table items(
   id serial primary key,
   description varchar(250),
   checklist_id integer references checklists (id) on delete cascade,
   created_at timestamp,
   expire_at timestamp
);

insert into items(description, checklist_id, created_at, expire_at)
values ('Going to the field', 1, '2024-05-31',  '2024-06-11'),
       ('Going to bed', 2, '2024-05-01', '2024-06-21');

select * from checklists;
select * from items;
