CREATE TABLE actor (
  actor_id numeric NOT NULL ,
  first_name VARCHAR(45) NOT NULL,
  last_name VARCHAR(45) NOT NULL,
  last_update DATE NOT NULL,
  CONSTRAINT pk_actor PRIMARY KEY  (actor_id)
);
