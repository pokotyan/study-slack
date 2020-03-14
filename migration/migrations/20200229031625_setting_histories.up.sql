CREATE TABLE IF NOT EXISTS setting_histories (
  id INT NOT NULL AUTO_INCREMENT,
  search_range INT NOT NULL,
  num_of_people INT NOT NULL,
  word varchar(255),
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL,
  PRIMARY KEY (`id`)
);
