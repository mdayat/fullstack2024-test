CREATE TABLE my_client (
  id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
  name CHAR(250) NOT NULL,
  slug CHAR(100) NOT NULL,
  is_project VARCHAR(30) NOT NULL DEFAULT '0' CHECK (is_project IN ('0', '1')),
  self_capture CHAR(1) NOT NULL DEFAULT '1',
  client_prefix CHAR(4) NOT NULL,
  client_logo CHAR(255) NOT NULL DEFAULT 'no-image.jpg',
  address TEXT DEFAULT NULL,
  phone_number CHAR(50) DEFAULT NULL,
  city CHAR(50) DEFAULT NULL,
  created_at TIMESTAMP(0) NULL,
  updated_at TIMESTAMP(0) NULL,
  deleted_at TIMESTAMP(0) NULL,

  PRIMARY KEY (id)
);