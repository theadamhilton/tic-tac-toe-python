import random

BOARD_WIDTH = 3
BOARD_HEIGHT = 3

def new_board():
    board = []
    for x in range(0, BOARD_WIDTH):
        column = []
        for y in range(0, BOARD_HEIGHT):
            column.append(None)
        board.append(column)

    return board

board = new_board()

def render(board):
    rows = []
    for y in range(0, BOARD_HEIGHT):
        row = []
        for x in range(0, BOARD_WIDTH):
            row.append(board[x][y])
        rows.append(row)

    row_num = 0
    print ('  0 1 2 ')
    print ('  ------')
    for row in rows:
        output_row = ''
        for sq in row:
            if sq is None:
                output_row += ' '
            else:
                output_row += sq
        print("%d|%s|" % (row_num, ' '.join(output_row)))
        row_num += 1
    print ('  ------')

render(board)

def random_ai(board, player):
    empty_spots = [(x, y) for x in range(BOARD_WIDTH) for y in range(BOARD_HEIGHT) if board[x][y] is None]

    if not empty_spots:
        raise Exception("No more moves available")
    
    return random.choice(empty_spots)

def get_move(player, board):
    if player == 'O': # if find_winning_and_losing_moves is not None: # uncomment code to switch to AI vs AI or comment code to switch to Human vs Human
        return find_winning_and_losing_moves(board, player) # comment code to switch to Human vs Human
    else: # comment code to switch to Human vs Human
        try:
            x = int(input("Enter the x-coordinate of your move (0, 1, or 2): "))
            y = int(input("Enter the y-coordinate of your move (0, 1, or 2): "))
            
            # Ensure the inputs are within the valid range
            if x in range(3) and y in range(3):
                return (x, y)
            else:
                print("Invalid input. Please enter coordinates between 0 and 2.")
                return get_move()
        except ValueError:
            print("Invalid input. Please enter numerical values.")
            return get_move()
        
def make_move(player, board, move_coords):
    if board[move_coords[0]][move_coords[1]] is not None:
        raise Exception("Invalid move, try again")

    board[move_coords[0]][move_coords[1]] = player
    return board

def get_all_line_coords():
    cols = []
    for x in range(0, BOARD_WIDTH):
        col = []
        for y in range(0, BOARD_HEIGHT):
            col.append((x, y))
        cols.append(col)

    rows = []
    for y in range(0, BOARD_HEIGHT):
        row = []
        for x in range(0, BOARD_WIDTH):
            row.append((x, y))
        rows.append(row)

    diagonals = [
        [(0, 0), (1, 1), (2, 2)],
        [(0, 2), (1, 1), (2, 0)]
    ]
    return cols + rows + diagonals

def get_winner(board):
    all_line_coords = get_all_line_coords()

    for line in all_line_coords:
        line_values = [board[x][y] for (x, y) in line]
        if len(set(line_values)) == 1 and line_values[0] is not None:
            return line_values[0]

    return None

def is_board_tie(board):
    winner = get_winner(board)
    
    for col in board:
        if any(sq is None for sq in col) or winner is not None:
            return False

    return True

def find_winning_moves_ai(board, player):
    # Check for a winning move
    for x in range(BOARD_WIDTH):
        for y in range(BOARD_HEIGHT):
            if board[x][y] is None:
                # Simulate making the move
                board[x][y] = player
                if get_winner(board) == player:
                    # Undo the move
                    board[x][y] = None
                    return (x, y)
                # Undo the move
                board[x][y] = None
    
    # If no winning move is found, return a random move
    return random_ai(board, player)

def find_winning_and_losing_moves(board, player):
    opponent = 'X' if player == 'O' else 'O'
    
    # Check for a winning move for the AI
    for x in range(BOARD_WIDTH):
        for y in range(BOARD_HEIGHT):
            if board[x][y] is None:
                board[x][y] = player
                if get_winner(board) == player:
                    board[x][y] = None
                    return (x, y)
                board[x][y] = None
    
    # Check for a move that blocks the opponent from winning
    for x in range(BOARD_WIDTH):
        for y in range(BOARD_HEIGHT):
            if board[x][y] is None:
                board[x][y] = opponent
                if get_winner(board) == opponent:
                    board[x][y] = None
                    return (x, y)
                board[x][y] = None
    
    # If no winning or blocking move is found, return a random move
    return random_ai(board, player)

def play():
    # Initial board setup
    board = [[None for _ in range(BOARD_WIDTH)] for _ in range(BOARD_HEIGHT)]
    players = ["X", "O"]
    turn = 0

    while True:
        # Determine the current player
        current_player = players[turn % 2]

        try:
            # Get and make a move
            move_coords = get_move(current_player, board)
            board = make_move(current_player, board, move_coords)
        except Exception as e:
            if "missing" not in str(e): # Hide specific error message
                print(e)
            continue # Retry the move if it was invalid
        
        # Render the board
        render(board)

        # Determine winner
        winner = get_winner(board)
        if winner is not None:
            render(board)
            print("THE WINNER IS %s!" % winner)
            break

        # Check if there is a draw
        if is_board_tie(board):
            render(board)
            print("IT'S A DRAW!")
            break

        # Switch to the next player
        turn += 1

play()