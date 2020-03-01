CREATE TABLE IF NOT EXISTS setting_histories (
  id INT NOT NULL AUTO_INCREMENT,
  search_range INT,
  num_of_people INT,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL,
  PRIMARY KEY (`id`)
);
