import os
import logging
import psycopg2 as psy


def connect_to_db():
    key = "PASSWORD"
    password = os.getenv(key)
    conn = psy.connect(host="localhost", dbname="employee", user="postgres", password=password)
    return conn


def setup_db():
    conn = connect_to_db()
    cursor = conn.cursor()
    try:
        cursor.execute(
            "CREATE TABLE IF NOT EXISTS info(empid serial PRIMARY KEY,  name VARCHAR(30), email VARCHAR(25))"
        )
    except psy.DatabaseError as err:
        logging.exception("Database error while creating table: {str(err)}")
        raise
    except psy.OperationalError as err:
        logging.exception("Operation error while creating table: {str(err)}")
        raise
    except Exception as err:
        logging.exception("Unexpected err while creating table: {str(err)}")
        raise

    try:
        cursor.execute("CREATE TABLE IF NOT EXISTS account(empid serial PRIMARY KEY,  number INT, type VARCHAR(10))")
    except psy.OperationalError as err:
        logging.exception("Operation error while creating table: {str(err)}")
        raise
    except Exception as err:
        logging.exception("Unexpected error while creating table: {str(err)}")
        raise

    conn.commit()
    close_db(conn, cursor)


def close_db(connection, cursor):
    if connection:
        cursor.close()
        connection.close()
