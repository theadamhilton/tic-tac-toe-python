from test import get_winner, is_board_tie

if __name__ == "__main__":
    board_1 = [
        ['X', 'X', 'O'],
        ['O', 'X', None],
        ['O', 'O', 'X']
    ]
    get_winner(board_1)
    is_board_tie(board_1)

    board_2 = [
        ['X', 'X', 'O'],
        ['O', None, 'X'],
        ['O', 'O', 'X']
    ]
    get_winner(board_2)
    is_board_tie(board_2)

    board_3 = [
        ['X', 'X', 'O'],
        ['O', None, 'X'],
        ['O', 'O', 'O']
    ]
    get_winner(board_3)
    is_board_tie(board_3)

    board_4 = [
        ['X', 'X', 'O'],
        ['O', 'O', 'X'],
        ['O', 'O', 'X']
    ]
    get_winner(board_4)
    is_board_tie(board_4)

    board_5 = [
        ['O', 'X', 'O'],
        ['X', 'X', 'O'],
        ['O', 'O', 'X']
    ]
    get_winner(board_5)
    is_board_tie(board_5)

    board_6 = [
        ['X', 'O', None],
        ['O', 'O', None],
        ['X', None, None]
    ]
    get_winner(board_6)
    is_board_tie(board_6)