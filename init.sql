-- Comando para criar a base de dados do app
CREATE DATABASE nw_app;
-- Comando para criar o usu?rio e senha padr?o para o app
CREATE USER user_app WITH ENCRYPTED PASSWORD 'user_app';
-- concede previlegios na base de dados para o usuario da app
GRANT ALL PRIVILEGES ON	DATABASE nw_app TO user_app;
--GRANT ALL PRIVILEGES ON	DATABASE nw_app TO postgres;
GRANT CONNECT ON DATABASE nw_app TO user_app;
