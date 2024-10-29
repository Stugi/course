-- Мой вариант
CREATE OR REPLACE PROCEDURE populate()
-- язык, на котором написана процедура
LANGUAGE plpgsql AS $$
DECLARE
    vendor_id integer;
BEGIN
    FOR i IN 1..1000 LOOP
        INSERT INTO vendors (name, website) VALUES ('Поставщик ' || i, 'site' || i||'.com') RETURNING id INTO vendor_id;
        INSERT INTO products (name, vendor_id) VALUES ('Товар ' || i, vendor_id);
    END LOOP;
END $$;

-- Другой вариант
-- создание или замена процедуры
CREATE OR REPLACE PROCEDURE populate()
-- язык, на котором написана процедура
LANGUAGE plpgsql
AS $$
DECLARE
	new_id int;
-- начало транзакции
BEGIN
	-- вставка нового поставщика и сохранение его id в переменную
	INSERT INTO vendors (name) VALUES ('Поставщик')
	RETURNING id INTO new_id;
	FOR i IN 1..1000 LOOP
 		INSERT INTO products(name, vendor_id)
 		VALUES ('Товар ' || i, new_id);
	END LOOP;
-- завершение процедуры и транзакции
END;
$$;

-- вызов процедуры
CALL populate ();

-- отображение результатов
SELECT * FROM products LIMIT 20;