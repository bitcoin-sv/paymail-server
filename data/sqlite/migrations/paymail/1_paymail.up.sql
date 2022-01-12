CREATE TABLE aliases (
    paymail VARCHAR PRIMARY KEY
    ,user_id INTEGER
    ,UNIQUE(paymail)
)
