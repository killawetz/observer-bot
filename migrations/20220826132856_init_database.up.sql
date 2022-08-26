create table chat (
                      id bigint primary key,
                      chat_name varchar(128)
);

create table users (
                       id bigint primary key,
                       username varchar(32) not null unique,
                       firstname varchar(64),
                       lastname varchar(64)
);

create table chat_users (
                            chat_id bigint references chat (id) on update cascade on delete cascade,
                            member_id bigint references users (id) on update cascade on delete cascade,
                            text_message int default 0,
                            sticker int default 0,
                            voice int default 0,
                            audio int default 0,
                            video int default 0,
                            video_note int default 0,
                            doc int default 0,
                            animation int default 0,
                            photo int default 0,


                            PRIMARY KEY (chat_id, member_id)
);