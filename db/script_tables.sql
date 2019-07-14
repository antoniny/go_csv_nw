-- Comando para criar a base de dados do app
CREATE DATABASE nw_app;
-- Comando para criar o usu?rio e senha padr?o para o app
CREATE USER user_app WITH ENCRYPTED PASSWORD 'user_app';
-- concede previlegios na base de dados para o usuario da app
GRANT ALL PRIVILEGES ON	DATABASE nw_app TO user_app;
GRANT CONNECT ON DATABASE nw_app TO user_app;
--GRANT ALL PRIVILEGES ON	DATABASE nw_app TO postgres;

DROP TABLE IF EXISTS file_csv;

CREATE TABLE file_csv (
  id                 serial    PRIMARY KEY, 
  cpf                character varying(20)              NULL, 
  private            character varying(20)              NULL, 
  incomplete         character varying(20)              NULL, 
  last_purchase      character varying(20)              NULL, 
  ticket_avg_amount  character varying(20)              NULL, 
  ticket_last_amount character varying(20)              NULL, 
  cnpj_max_frequency character varying(20)              NULL, 
  cnpj_last_purchase character varying(20)              NULL, 
	  last_change_time   timestamp(6) without time zone     NULL DEFAULT now(), 
  cpf_ok             character varying(1)               NULL, 
  cnpj_ok            character varying(1)               NULL
)


COMMENT ON COLUMN file_csv.id IS 'ID Incremental';
COMMENT ON COLUMN file_csv.cpf IS 'CPF do Cliente';
COMMENT ON COLUMN file_csv.private IS ' ';
COMMENT ON COLUMN file_csv.incomplete IS ' ';
COMMENT ON COLUMN file_csv.last_purchase IS 'Data da ultima compra';
COMMENT ON COLUMN file_csv.ticket_avg_amount IS 'Valor do Ticket médio';
COMMENT ON COLUMN file_csv.ticket_last_amount IS 'Valor do ultimo ticket';
COMMENT ON COLUMN file_csv.cnpj_max_frequency IS 'Loja com maior frequencia';
COMMENT ON COLUMN file_csv.last_change_time IS 'Ultima alteraçao no registro';

commit;

