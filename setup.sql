DROP TABLE IF EXISTS admins;
CREATE TABLE admins (
  user varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  role int NOT NULL,
  UNIQUE KEY user (user)
);

INSERT INTO admins VALUES ('alice','pass1234',1);
INSERT INTO admins VALUES ('bob','pass2345',2);
INSERT INTO admins VALUES ('craig','pass3456',4);
INSERT INTO admins VALUES ('dan','pass4567',8);

DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id int NOT NULL AUTO_INCREMENT,
  username varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  age int NOT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY username (username),
  UNIQUE KEY email (email)
);
