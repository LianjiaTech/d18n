CREATE TABLE actor (
  actor_id int NOT NULL IDENTITY ,
  first_name VARCHAR(45) NOT NULL,
  last_name VARCHAR(45) NOT NULL,
  last_update DATETIME NOT NULL,
  PRIMARY KEY NONCLUSTERED (actor_id)
  )
GO
 ALTER TABLE actor ADD CONSTRAINT [DF_actor_last_update] DEFAULT (getdate()) FOR last_update
GO
 CREATE  INDEX idx_actor_last_name ON actor(last_name) 
GO
