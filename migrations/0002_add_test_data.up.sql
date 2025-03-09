-- 0002_add_test_data.up.sql
INSERT INTO tasks (title, description, status) VALUES
                                                   ('Изучить Go', 'Изучить основы языка Go', 'new'),
                                                   ('Изучить Fiber', 'Изучить фреймворк Fiber', 'in_progress'),
                                                   ('Изучить PostgreSQL', 'Изучить работу с PostgreSQL', 'done');