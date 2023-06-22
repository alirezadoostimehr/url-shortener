CREATE TABLE IF NOT EXISTS urls (
    id int PRIMARY KEY AUTO_INCREMENT ,
    srt varchar(255) NOT NULL ,
    lng varchar(255) NOT NULL ,

    created_at datetime DEFAULT NOW(),
    updated_at datetime DEFAULT NOW() ON UPDATE NOW()
);