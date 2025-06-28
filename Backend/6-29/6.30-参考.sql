-- 参考结构体
type Storage struct{
	ID int
	Stock int
	BookID string 
}
type PersonBookRelationship struct{
    ID int 
    BookID string 
    PersonID int 
}

-- 参考SQL语句
-- 1
CREATE TABLE book (
    id VARCHAR(50) PRIMARY KEY,
    title VARCHAR(100),
    author VARCHAR(100)
);

CREATE TABLE storage (
    book_id VARCHAR(50) PRIMARY KEY,
    stock INT NOT NULL
);

CREATE TABLE person (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) UNIQUE
);

CREATE TABLE person_book_relationship (
    person_id INT,
    book_id VARCHAR(50),
    id INT AUTO_INCREMENT PRIMARY KEY,
    UNIQUE (person_id, book_id)
);

-- 2
INSERT INTO book (id, title, author) VALUES
('go-away', 'the way to go', 'Ivo'),
('go-lang', 'Go语言圣经', 'Alan, Brian'),
('go-web', 'Go Web编程', 'Anonymous'),
('con-cur', 'Concurrency in Go', 'Katherine');

INSERT INTO storage (book_id, stock) VALUES
('go-away', 20),
('go-lang', 17),
('go-web', 4),
('con-cur', 9);

INSERT INTO person (name) VALUES
('小明'),
('张三'),
('翟曙');

INSERT INTO person_book_relationship (person_id, book_id) VALUES
(1, 'go-away'),
(1, 'go-web'),
(1, 'con-cur'),
(2, 'go-away'),
(3, 'go-web'),
(3, 'con-cur');

-- 3

-- 3.1
SELECT p.name
FROM person p
JOIN person_book_relationship pb ON p.id = pb.person_id
WHERE pb.book_id = 'go-web';

-- 3.2
SELECT b.id, b.author, b.title, s.stock
FROM book b
JOIN storage s ON b.id = s.book_id
LEFT JOIN person_book_relationship pb ON b.id = pb.book_id
WHERE pb.book_id IS NULL;

-- 3.3
SELECT p.name AS person_name, b.title AS book_title
FROM person_book_relationship pb
JOIN person p ON pb.person_id = p.id
JOIN book b ON pb.book_id = b.id
ORDER BY p.name;
