'''new_board:
We’re going to start by making a “data structure” 
that stores the state of the game board. Tic-Tac-Toe 
is played on a 3x3 grid, and in code, grids are often 
represented using nested lists, or “lists of lists”. 
Each element in our list-of-lists will represent a 
square on the board. In each square, we will need to 
be able to represent the 3 possible states - O, X and
empty. I suggest representing O and X using the 
characters O and X, and an empty square using None 
or nil or however it is that your programming language 
represents “nothing”.

Your first task is to write a new_board function that
takes in 0 arguments and returns an empty, 3x3 grid. 
Test that your function works with a simple print-statement.
'''

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

# or use list comprehension
# def new_board():
#    return [[None for _ in range(BOARD_HEIGHT)] for _ in range(BOARD_WIDTH)]

board = new_board()
# print(board)


'''render:
Basic print-statements are fine for debugging, but our players 
deserve something more refined. In this milestone, we’re going 
to write a function whose only job is to pretty-print our board 
to the terminal in a form that is useful to humans.

Write a render function that takes in 1 argument - 
a Tic-Tac-Toe board - and prints it to the terminal. 
This function does not need to return anything. You 
can format the board however you like; I personally 
found it useful to add co-ordinate markers around the 
edge to make it easier for human players to reference 
specific squares.

Test your render function by using it to print some dummy 
boards. Construct these boards by generating an empty board 
using new_board and adding moves to it manually.'''

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

# or use list comprehension
'''def render(board):
    # Create a list of rows by transposing the board
    rows = [[board[x][y] for x in range(BOARD_WIDTH)] for y in range(BOARD_HEIGHT)]
    
    print('  0 1 2 ')
    print('  ------')

    # Print each row with appropriate formatting
    for row_num, row in enumerate(rows):
        output_row = ''.join(' ' if sq is None else sq for sq in row)
        print(f"{row_num}|{' '.join(output_row)}|")

    print('  ------')'''

# board[0][1] = 'X'
# board[1][1] = 'O'
render(board)


'''get_move:
Now that we’ve got a board and a way to display it, we’re ready to start playing.

In this milestone we’re going to work out how to ask our human players to input 
their moves. Because we understand the extreme power of focusing on one thing 
at a time, we’re not yet going to worry about using these moves to update the board. 
Once again, this temporary shortcut is not cheating.

We could ask our players to input their moves in many different formats. I like asking 
for X and Y co-ordinates, because this maps well onto our grid data structure. But you 
could also assign each square a number from 1 to 9, or a word, or a letter.

Once we have turned a player’s input into a move co-ordinate, the code that uses this 
co-ordinate to update the board shouldn’t have to know or concern itself with where the 
co-ordinate came from. This is a concept known as modularity.

Write a get_move function that takes 0 arguments, and returns the co-ordinates of the 
player’s chosen move as a 2-element tuple or array (one element each for the x and y 
co-ordinates), or whatever works best in your language.'''

def get_move():
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

# move_coords = get_move()
# print(move_coords)


'''make_move
Now that our players can input their moves, it’s time for us to use those moves to update the board. 
We’re still not going to worry about whether either of our players has won. You can probably guess why.

The make_move function will need 3 arguments:
- The board
- The co-ordinates of the move
- The player who is making the move (so that make_move knows whether to insert an O or an X into the board)'''

def make_move(player, board, move_coords):
    if board[move_coords[0]][move_coords[1]] is not None:
        raise Exception("Invalid move, try again")

    board[move_coords[0]][move_coords[1]] = player
    return board

# move_coords = (2, 0)

'''try:
    board = make_move("X", board, move_coords)
except Exception as e:
    print(e)

try:
    board = make_move("O", board, move_coords)
except Exception as e:
    print(e)

render(board)'''


'''get_winner
When I was thinking about checking for winners, I asked myself 
“how does a human work out whether either player has won a game 
of Tic Tac Toe?” Asking “how would a human do this?” can be a 
useful way to design programs. Often (not always), a computer 
will do things in the same way as a human would, only faster. 
In this case, the answer I came up with was “a human would look 
at every possible line of 3 grid squares, and check whether any 
of them contain all Os or all Xs.”

Write a get_winner function that takes take 1 argument - a board - 
and returns the ID of the player that has won (O or X), or None 
if no one has won yet. I’d suggest that you structure your function 
like this:

- Build a list of all the lines on the board - columns, rows and 
diagonals (eg. - [[O, X, None], [X, O, X], # etc...)
- For each line in this list, check whether it contains all Os or 
all Xs. If yes, return the character of the winning player (O or X). 
If no, continue onto the next line
- If none of the lines on the board are winners, return None or nil 
to indicate that no one has won yet

Test your function with some practice boards via unit tests.'''

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
            # print(line_values[0]) # For unit testing
            return line_values[0]

    # print(None) # For unit testing
    return None


'''is_board_full
Check for a draw since not all games of Tic-Tac-Toe end in a victory. 
If both players are evenly matched meaning the entire board is full of 
letters but no column/row/diagonal matches, then a game may end in a 
hard-fought draw.'''

'''def is_board_tie(board):
    winner = get_winner(board)

    for col in board:
        for sq in col:
            if sq is None:
                print(False) # For unit testing
                return False
    
    if winner is not None:
        print(False) # For unit testing
        return False
    
    print(True) # For unit testing
    return True'''

# or use refactored code
def is_board_tie(board):
    winner = get_winner(board)
    
    # Check if any square is None or if there is a winner
    for col in board:
        if any(sq is None for sq in col) or winner is not None:
            # print(False)  # For unit testing
            return False

    # print(True)  # For unit testing
    return True



'''We did not create pseudo-code for the following function 
so we will create its production code here. 

Here, we are combining the get_move and make_move functions 
to alternate player turns and automate the tested tasks to 
play the game.'''

def play():
    # Initial board setup
    board = [[None for _ in range(3)] for _ in range(3)] # List comprehension
    players = ["X", "O"]
    turn = 0

    while True:
        # Determine the current player
        current_player = players[turn % 2]

        try:
            # Get and make a move
            move_coords = get_move()
            board = make_move(current_player, board, move_coords)
        except Exception as e:
            print(e)
            continue # Retry the move if it was invalid
        
        # Render the board
        render(board)

        # Determine winner. Only add this logic after unit testing.
        winner = get_winner(board)
        if winner is not None:
            render(board)
            print("THE WINNER IS %s!" % winner)
            break

        # Check if there is a draw. Only add this logic after unit testing.
        if is_board_tie(board):
            render(board)
            print("IT'S A DRAW!")
            break

        # Switch to the next player
        turn += 1

play()