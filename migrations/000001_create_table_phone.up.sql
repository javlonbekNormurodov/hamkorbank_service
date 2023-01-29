create table if not exists phone (
    "id" uuid primary key not null gen_random_uuid(),
    "phone_number" varchar(255) not null,
    "created_at" timestamp default current_timestamp not null
);

insert into phone(phone_number) values
    ("+99890909090"),
    ("+99890909092"),
    ("+99893909090"),
    ("+99890909321"),
    ("+99890905614");
