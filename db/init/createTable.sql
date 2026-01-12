USE goapi;

create table if not exists articles (
  article_id integer unsigned auto_increment primary key,
  title varchar(100) not null,
  contents text not null,
  username varchar(100) not null,
  nice integer not null,
  created_at datetime
);

create table if not exists comments (
  comment_id integer unsigned auto_increment primary key,
  article_id integer unsigned not null,
  message text not null,
  created_at datetime,
  foreign key (article_id) references articles(article_id)
);

create table if not exists users (
  id char(36) not null,
  name varchar(255) not null,
  email varchar(255),
  created_at datetime not null default current_timestamp,
  PRIMARY KEY (id),
  UNIQUE KEY uk_email (email)
);

create table if not exists user_identities (
  id bigint auto_increment primary key,
  user_id char(36) not null,
  provider varchar(50) not null,
  provider_sub varchar(255) not null,
  created_at datetime not null default current_timestamp,
  unique key uk_provider_sub (provider, provider_sub),
  foreign key (user_id) references users(id)
);

create table if not exists user_credentials (
  user_id char(36) not null,
  password_hash varchar(255) not null,
  created_at datetime not null default current_timestamp,
  primary key (user_id),
  foreign key (user_id) references users(id)
);
