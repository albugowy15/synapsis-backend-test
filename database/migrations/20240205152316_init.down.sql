DO $$ 
DECLARE 
    r RECORD;
BEGIN 
    -- Drop functions
    FOR r IN (
        SELECT proname, oid 
        FROM pg_proc 
        WHERE pronamespace = 'public'::regnamespace
    ) LOOP
        EXECUTE 'DROP FUNCTION IF EXISTS ' || r.proname || ' CASCADE';
    END LOOP;

    -- Drop triggers
    FOR r IN (
        SELECT trigger_name 
        FROM information_schema.triggers 
        WHERE trigger_schema = 'public'
    ) LOOP
        EXECUTE 'DROP TRIGGER IF EXISTS ' || r.trigger_name || ' ON ' || r.table_name || ' CASCADE';
    END LOOP;
END $$;

DROP TABLE IF EXISTS "payment_transactions" CASCADE;
DROP TABLE IF EXISTS "payment_methods" CASCADE;
DROP TABLE IF EXISTS "products_categories" CASCADE;
DROP TABLE IF EXISTS "categories" CASCADE;
DROP TABLE IF EXISTS "products" CASCADE;
DROP TABLE IF EXISTS "items" CASCADE;
DROP TABLE IF EXISTS "shopping_carts" CASCADE;
DROP TABLE IF EXISTS "users" CASCADE;
