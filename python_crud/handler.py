from utils.helpers import setup_db, create_row, read_rows, update_row, delete_row

# from utils.conf import PSQLERROR


def main():

    setup_db()

    create_row()
    read_rows()
    update_row()
    delete_row()


if __name__ == "__main__":
    main()
