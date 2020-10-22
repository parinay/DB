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
        cursor.execute("CREATE TABLE info(empid serial PRIMARY KEY,  name VARCHAR(30), email VARCHAR(25))")
    except Exception as err:
        logging.exception("Unexpected err while creating table: {str(err)}")

    try:
        cursor.execute("CREATE TABLE account(empid serial PRIMARY KEY,  number INT, type VARCHAR(10))")
    except Exception as err:
        logging.exception("Unexpected error while creating table: {str(err)}")

    conn.commit()
    close_db(conn, cursor)


def close_db(connection, cursor):
    if connection:
        cursor.close()
        connection.close()
