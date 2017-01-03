CREATE SCHEMA postgres_template;

CREATE SEQUENCE postgres_template.data_row_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE postgres_template.data_rows (
    id integer DEFAULT nextval('meal_allowance.data_row_seq'::regclass) NOT NULL,
    data text NOT NULL
);

INSERT INTO postgres_template.data_rows (data) VALUES ('Chicken');
INSERT INTO postgres_template.data_rows (data) VALUES ('Dog');
INSERT INTO postgres_template.data_rows (data) VALUES ('Cat');
INSERT INTO postgres_template.data_rows (data) VALUES ('Eagle');
INSERT INTO postgres_template.data_rows (data) VALUES ('Dragon');