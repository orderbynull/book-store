    CREATE TABLE authors
    (
        id INTEGER PRIMARY KEY NOT NULL,
        name VARCHAR(255)
    );
    
    INSERT INTO public.authors (id, name) VALUES (1, 'Charles Blum');
    INSERT INTO public.authors (id, name) VALUES (2, 'Larry Ulman');
    INSERT INTO public.authors (id, name) VALUES (3, 'Dave Cheney');
    INSERT INTO public.authors (id, name) VALUES (4, 'Aamir Khan');
    INSERT INTO public.authors (id, name) VALUES (5, 'Aaron Torres');
--

    CREATE TABLE books
    (
        id INTEGER PRIMARY KEY NOT NULL,
        title VARCHAR(255) NOT NULL,
        author_id INTEGER NOT NULL,
        trash BOOLEAN DEFAULT false NOT NULL
    );

--

    CREATE TABLE books_bookshelfs
    (
        book_id INTEGER NOT NULL,
        bookshelf_id INTEGER NOT NULL
    );
    

--

    CREATE TABLE bookshelfs
    (
        id INTEGER PRIMARY KEY NOT NULL,
        title VARCHAR(255) NOT NULL
    );