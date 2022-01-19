CREATE TABLE aliases (
    paymail VARCHAR PRIMARY KEY
    ,user_id INTEGER
    ,UNIQUE(paymail)
);

INSERT INTO 
    aliases(paymail, user_id)
    VALUES("epic@nchain.com", 1);
