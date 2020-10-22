import logging
from flask import Flask, jsonify, request
from utils.helpers import setup_db, connect_to_db, close_db

# from utils.conf import PSQLERROR


app = Flask(__name__)


def main():

    setup_db()
    app.run(debug=True)


@app.route("/apiv1/employees/<int:empid>", methods=["GET"])
def read_row(empid):
    conn = connect_to_db()
    cursor = conn.cursor()
    query_info = "SELECT * FROM info WHERE empid = %s"
    cursor.execute(query_info, (empid,))
    items = cursor.fetchall()
    if items:
        return jsonify(items)
    else:
        print("Its empty")

    query_acc = "SELECT * FROM account WHERE empid = %s"
    cursor.execute(query_acc, (empid,))

    items = cursor.fetchall()
    if items:
        return jsonify(items)
    else:
        print("Its empty")
    conn.commit()
    close_db(conn, cursor)


@app.route("/apiv1/employees", methods=["GET"])
def read_rows():
    conn = connect_to_db()
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM info")
    items = cursor.fetchall()

    if items:
        return jsonify(items)
    else:
        print("Its empty")
    
    cursor.execute("SELECT * FROM account")
    items = cursor.fetchall()

    if items:
        return jsonify(items)
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
    conn = connect_to_db()
    cursor = conn.cursor()

    # code to delete the row in question

    conn.commit()
    close_db(conn, cursor)


def update_row():
    print("updating row")
    conn = connect_to_db()
    cursor = conn.cursor()

    # code to update the row in question

    conn.commit()
    close_db(conn, cursor)


if __name__ == "__main__":
    main()
