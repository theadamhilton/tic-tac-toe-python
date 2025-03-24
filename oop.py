import random

class Board:
    def __init__(self, width=3, height=3):
        self.width = width
        self.height = height
        self.grid = [[None for _ in range(self.width)] for _ in range(self.height)]

    def initialize(self):
        """Resets the board to its initial empty state."""
        self.grid = [[None for _ in range(self.width)] for _ in range(self.height)]

    def render(self):
        """Displays the board visually."""
        print('  0 1 2 ')
        print('  ------')
        for y, row in enumerate(self.grid):
            output_row = f"{y}|"
            for cell in row:
                output_row += cell if cell is not None else ' '
                output_row += " "
            print(output_row + "|")
        print('  ------')

    def make_move(self, x, y, marker):
        """Places a marker on the board."""
        if self.grid[y][x] is not None:
            raise Exception("Invalid move, square is already occupied.")
        self.grid[y][x] = marker

    def get_empty_spots(self):
        """Returns a list of empty spots on the board."""
        return [(x, y) for y in range(self.height) for x in range(self.width) if self.grid[y][x] is None]

    def check_winner(self):
        """Checks for a winner on the board."""
        lines = self.get_all_lines()
        for line in lines:
            markers = [self.grid[y][x] for (x, y) in line]
            if len(set(markers)) == 1 and markers[0] is not None:
                return markers[0]
        return None

    def is_tie(self):
        """Checks if the board is in a tie state."""
        return self.check_winner() is None and all(cell is not None for row in self.grid for cell in row)

    def get_all_lines(self):
        """Returns all possible winning lines (rows, columns, diagonals)."""
        rows = [[(x, y) for x in range(self.width)] for y in range(self.height)]
        columns = [[(x, y) for y in range(self.height)] for x in range(self.width)]
        diagonals = [[(i, i) for i in range(self.width)], [(i, self.width - i - 1) for i in range(self.width)]]
        return rows + columns + diagonals
    
    '''def find_winning_and_losing_moves(self, player, board):
        opponent = 'X' if player == 'O' else 'O'

        # Check for winning move
        for x, y in self.get_empty_spots():
            self.grid[y][x] = player
            if self.check_winner() == player:
                self.grid[y][x] = None
                return (x, y)
            self.grid[y][x] = None

        # Check for blocking move
        for x, y in self.get_empty_spots():
            self.grid[y][x] = opponent
            if self.check_winner() == opponent:
                self.grid[y][x] = None
                return (x. y)
            self.grid[y][x] = None

        # If no winning or blocking move, return a random move
        return random.choice(self.get_empty_spots())

    def minimax(self, depth, is_maximizing, player, opponent):
        winner = self.check_winner()
        if winner == player:
            return 10 - depth
        elif winner == opponent:
            return depth - 10
        elif self.is_tie():
            return 0

        if is_maximizing:
            max_eval = float('-inf')
            for x, y in self.get_empty_spots():
                self.grid[y][x] = player
                eval = self.minimax(depth + 1, False, player, opponent)
                self.grid[y][x] = None
                max_eval = max(max_eval, eval)
            return max_eval
        else:
            min_eval = float('inf')
            for x, y in self.get_empty_spots():
                self.grid[y][x] = opponent
                eval = self.minimax(depth + 1, True, player, opponent)
                self.grid[y][x] = None
                min_eval = min(min_eval, eval)
            return min_eval
    
    def minimax_move(self, player, board):
        print("Thinking...")
        opponent = 'X' if player == 'O' else 'O'
        best_score = float('-inf')
        best_move = None

        for x, y in self.get_empty_spots():
            self.grid[y][x] = player
            score = self.minimax(0, False, player, opponent)
            self.grid[y][x] = None
            if score > best_score:
                best_score = score
                best_move = (x, y)
        
        print(f"Player {player} selects move {best_move} with score {best_score}")
        return best_move'''
    
class Player:
    def __init__(self, name, marker, strategy=None):
        self.name = name
        self.marker = marker
        self.strategy = strategy

    def get_move(self, board):
        if self.strategy:  # If a strategy is provided, use it (AI player)
            return self.strategy(self.marker, board)
        else:  # Otherwise, it's a human player
            try:
                x = int(input(f"Player ({self.marker}), enter the x-coordinate of your move (0, 1, or 2): "))
                y = int(input(f"Player ({self.marker}), enter the y-coordinate of your move (0, 1, or 2): "))
                
                if x in range(3) and y in range(3):
                    return (x, y)
                else:
                    print("Invalid input. Please enter coordinates between 0 and 2.")
                    return self.get_move(board)
            except ValueError:
                print("Invalid input. Please enter numerical values.")
                return self.get_move(board)
            
def play_game(player1, player2):
    # Initialize the game board
    board = Board()
    board.initialize()
    players = [player1, player2]
    turn = 0

    print("Welcome to Tic-Tac-Toe!")
    while True:
        # Determine the current player
        current_player = players[turn % 2]

        try:
            # Render the current state of the board
            board.render()

            # Get and make the move
            move_coords = current_player.get_move(board.grid)
            board.make_move(move_coords[0], move_coords[1], current_player.marker)
        except Exception as e:
            print(e)
            continue  # Retry the turn if an error occurs

        # Check for a winner
        winner = board.check_winner()
        if winner is not None:
            board.render()
            print(f"THE WINNER IS {winner}!")
            break

        # Check for a tie
        if board.is_tie():
            board.render()
            print("IT'S A DRAW!")
            break

        # Switch to the next player
        turn += 1

    print("Game Over. Thanks for playing!")

    # Define a simple random AI strategy
def random_ai(player, board):
    empty_spots = [(x, y) for y in range(3) for x in range(3) if board[y][x] is None]
    if not empty_spots:
        raise Exception("No more moves available")
    return random.choice(empty_spots)

# Create the players
player1 = Player(name="Human", marker="X")  # Human player
player2 = Player(name="AI", marker="O", strategy=random_ai)  # AI player

# Test the game
play_game(player1, player2)