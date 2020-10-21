import logging
import psycopg2 as psy


def connect_to_db():
    conn = psy.connect(host="localhost", dbname="employee", user="postgres", password="parinay")
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


def read_rows():
    conn = connect_to_db()
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM info")
    cursor.execute("SELECT * FROM account")
    items = cursor.fetchall()

    if items:
        print(items)
    else:
        print("Its empty")
    conn.commit()
    close_db(conn, cursor)


def create_row():
    conn = connect_to_db()
    cursor = conn.cursor()
    info_record = (4, "Steve w", "steve@email.com")
    acc_record = (3, 1236, "savings")
    try:
        cursor.execute("INSERT INTO info (empid, name, email) VALUES (%s, %s, %s)", info_record)
    except Exception as err:
        logging.exception("Failed to insert record in table info: {str(err)}")

    try:
        cursor.execute("INSERT INTO account (empid, number, type) VALUES (%s, %s, %s)", acc_record)
    except Exception as err:
        logging.exception("Failed to insert record in table info: {str(err)}")
    conn.commit()
    close_db(conn, cursor)


def delete_row():
    print("deleting row")


def update_row():
    print("updating row")


def close_db(connection, cursor):
    if connection:
        cursor.close()
        connection.close()
