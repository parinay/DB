from utils.helpers import setup_db, create_row, read_rows, update_row, delete_row
#from flask import Flask

# from utils.conf import PSQLERROR

#app = Flask(__name__)

#@app.route("/")
def main():

    setup_db()

    create_row()
    read_rows()
    update_row()
    delete_row()


if __name__ == "__main__":
    main()
