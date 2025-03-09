-- 0002_add_test_data.down.sql
DELETE FROM tasks WHERE title IN (
                                  'Изучить Go',
                                  'Изучить Fiber',
                                  'Изучить PostgreSQL'
    );