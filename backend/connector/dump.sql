-- Active: 1707650866017@@127.0.0.1@3306@bookrzn
CREATE TABLE Favorites(  
    id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    token VARCHAR(255),
    target_hash VARCHAR(255),
    count VARCHAR(255),
    date VARCHAR(255),
    id_order VARCHAR(255),
    status_order VARCHAR(255)
);


CREATE TABLE Orders(  
    id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    token VARCHAR(255),
    target_hash VARCHAR(255),
    count VARCHAR(255),
    date VARCHAR(255),
    id_order VARCHAR(255),
    status_order VARCHAR(255)
);

CREATE TABLE OrdersHistory(  
    id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    token VARCHAR(255),
    target_hash VARCHAR(255),
    count VARCHAR(255),
    date VARCHAR(255),
    id_order VARCHAR(255),
    status_order VARCHAR(255)
);

CREATE TABLE Targets(  
    id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    target_hash VARCHAR(255),
    autor VARCHAR(255),
    title VARCHAR(255),
    price VARCHAR(255),
    image VARCHAR(255),
    comment VARCHAR(255),
    url_source VARCHAR(255),
    target_type VARCHAR(255)

);

CREATE TABLE Users(  
    id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    login VARCHAR(255),
    password VARCHAR(255),
    type VARCHAR(255),
    token VARCHAR(255),
    name VARCHAR(255),
    family VARCHAR(255),
    phone VARCHAR(255),
    email VARCHAR(255)
);
