insert into articles (
  title, contents, username, nice, created_at
) values (
'first post', 'This is a test blog post', 'test user', 1, now()
);

insert into articles (
  title, contents, username, nice, created_at
) values (
'2nd post', '2nd blog post', 'test user', 1, now()
);

insert into comments (
  article_id, message, created_at
) values (
1, 'This is a test comment', now()
);


insert into comments (
  article_id, message
) values (
1, 'hello'
);

