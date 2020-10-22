import logging
from flask import Flask, jsonify, abort, make_response, request
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
        abort(404)

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
        abort(404)

    cursor.execute("SELECT * FROM account")
    items = cursor.fetchall()

    if items:
        return jsonify(items)
    else:
        print("Its empty")
    conn.commit()
    close_db(conn, cursor)


@app.errorhandler(404)
def not_found(error):
    return make_response(jsonify({"error": "Not Found"}), 404)


@app.route("/apiv1/employees", methods=["POST"])
def create_row():
    if not request.json:
        abort(400)

    conn = connect_to_db()
    cursor = conn.cursor()
    logging.info("The request is {requesti.json}")
    employee_id = (request.json["empid"],)
    employee_name = request.json["name"]
    employee_email = request.json["email"]
    info_record = (employee_id, employee_name, employee_email)
    acc_record = (5, 1238, "savings")
    try:
        cursor.execute("INSERT INTO info (empid, name, email) VALUES (%s, %s, %s)", info_record)
    except Exception as err:
        logging.exception("Failed to insert record in table info: {str(err)}")

    try:
        cursor.execute("INSERT INTO account (empid, number, type) VALUES (%s, %s, %s)", acc_record)
    except Exception as err:
        logging.exception("Failed to insert record in table info: {str(err)}")

    conn.commit()
    return jsonify(info_record), 201
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
