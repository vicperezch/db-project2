-- Tabla de autores
CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    names VARCHAR(100) NOT NULL CHECK (LENGTH(names) >= 2),
    last_names VARCHAR(100) NOT NULL CHECK (LENGTH(last_names) >= 2),
    nationality VARCHAR(50) NOT NULL
);

-- Tabla de los tipos de publicaciones
CREATE TABLE publications_type (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE CHECK (LENGTH(name) >= 3)
);

-- Tabla para los géneros de las publicaciones
CREATE TABLE genres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE CHECK (LENGTH(name) >= 3),
    description TEXT
);

-- Tabla de publicaciones
CREATE TABLE publications (
    id SERIAL PRIMARY KEY,
    id_publication_type INTEGER NOT NULL REFERENCES publications_type(id),
    id_genre INTEGER NOT NULL REFERENCES genres(id),
    title VARCHAR(100) NOT NULL CHECK (LENGTH(title) >= 2),
    publication_date DATE NOT NULL CHECK (publication_date <= CURRENT_DATE),
    isbn VARCHAR(20) UNIQUE NOT NULL CHECK (LENGTH(isbn) >= 8),
    total_editions INTEGER NOT NULL DEFAULT 1 CHECK (total_editions >= 0)
);

-- Tabla de ediciones
CREATE TABLE editions (
    id SERIAL PRIMARY KEY,
    id_publication INTEGER NOT NULL REFERENCES publications(id),
    edition_number INTEGER NOT NULL CHECK (edition_number >= 1),
    edition_date DATE NOT NULL CHECK (edition_date <= CURRENT_DATE),
    description TEXT,
    UNIQUE (id_publication, edition_number)
);

-- Tabla de clientes
CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    names VARCHAR(100) NOT NULL CHECK (LENGTH(names) >= 2),
    last_names VARCHAR(100) NOT NULL CHECK (LENGTH(last_names) >= 2),
    address TEXT NOT NULL CHECK (LENGTH(address) >= 5),
    phone VARCHAR(10) NOT NULL CHECK (phone ~ '^[0-9]{10}$'),
    email VARCHAR(100) NOT NULL CHECK (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
    total_loans INTEGER DEFAULT 0 CHECK (total_loans >= 0),
    total_fines INTEGER DEFAULT 0 CHECK (total_fines >= 0)
);

-- Tabla de empleados
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    names VARCHAR(100) NOT NULL CHECK (LENGTH(names) >= 2),
    last_names VARCHAR(100) NOT NULL CHECK (LENGTH(last_names) >= 2),
    hiring_date DATE NOT NULL CHECK (hiring_date <= CURRENT_DATE)
);

-- Tabla de préstamos
CREATE TABLE loans (
    id SERIAL PRIMARY KEY,
    id_client INTEGER NOT NULL REFERENCES clients(id),
    id_employee INTEGER NOT NULL REFERENCES employees(id),
    id_edition INTEGER NOT NULL REFERENCES editions(id),
    loan_date DATE NOT NULL DEFAULT CURRENT_DATE CHECK (loan_date <= CURRENT_DATE),
    return_date DATE NOT NULL CHECK (return_date > loan_date)
);

-- Tabla de los autores de las publicaciones
CREATE TABLE authors_publications (
    id SERIAL PRIMARY KEY,
    id_author INTEGER NOT NULL REFERENCES authors(id),
    id_publication INTEGER NOT NULL REFERENCES publications(id),
    is_main_author BOOLEAN DEFAULT FALSE,
    UNIQUE (id_author, id_publication)
);

-- Tabla de multas
CREATE TABLE fines (
    id SERIAL PRIMARY KEY,
    id_loan INTEGER NOT NULL REFERENCES loans(id),
    amount DECIMAL(10,2) NOT NULL CHECK (amount > 0),
    reason TEXT,
    date_payment DATE NOT NULL CHECK (date_payment <= CURRENT_DATE)
);

-- Tabla reviews
CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    id_client INTEGER NOT NULL REFERENCES clients(id),
    id_edition INTEGER NOT NULL REFERENCES editions(id),
    comment TEXT CHECK (comment IS NULL OR LENGTH(comment) >= 10),
    rating INTEGER NOT NULL CHECK (rating BETWEEN 1 AND 5),
    review_date DATE DEFAULT CURRENT_DATE,
    UNIQUE (id_client, id_edition)
);

-- Triggers
-- 1. Trigger para actualizar el total de préstamos de un cliente
CREATE OR REPLACE FUNCTION update_client_loans()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE clients
    SET total_loans = total_loans + 1
    WHERE id = NEW.id_client;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_after_insert_loan
AFTER INSERT ON loans
FOR EACH ROW
EXECUTE FUNCTION update_client_loans();

-- 2. Trigger para actualizar el total de ediciones de una publicación
CREATE OR REPLACE FUNCTION update_publication_editions()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE publications
    SET total_editions = total_editions + 1
    WHERE id = NEW.id_publication;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_after_insert_edition
AFTER INSERT ON editions
FOR EACH ROW
EXECUTE FUNCTION update_publication_editions();

-- 3. Trigger para actualizar el total de multas de un cliente
CREATE OR REPLACE FUNCTION update_client_fines()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE clients
    SET total_fines = total_fines + 1
    WHERE id = (SELECT id_client FROM loans WHERE id = NEW.id_loan);

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_after_insert_fine
AFTER INSERT ON fines
FOR EACH ROW
EXECUTE FUNCTION update_client_fines();
