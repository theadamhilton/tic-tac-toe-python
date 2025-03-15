import asyncio
import websockets
import json

# A dictionary to store connected clients and game sessions
connected_clients = {}

# Function to broadcast a message to all clients in the same session
async def broadcast_message(session, message):
    if session in connected_clients:
        for websocket in connected_clients[session]:
            await websocket.send(message)

# Function to handle messages from clients
async def handle_message(websocket, session, message):
    try:
        data = json.loads(message)
        # Example: If the message is a move, broadcast it to the session
        if data["type"] == "move":
            await broadcast_message(session, message)
        elif data["type"] == "join":
            # Handle a new client joining the session
            if session not in connected_clients:
                connected_clients[session] = []
            connected_clients[session].append(websocket)
            print(f"Client joined session: {session}")
    except json.JSONDecodeError:
        await websocket.send(json.dumps({"type": "error", "message": "Invalid message format"}))

# Main WebSocket connection handler
async def connection_handler(websocket, path):
    session = None  # Track the session the client joins
    try:
        # Wait for a "join" message to determine the session
        join_message = await websocket.recv()
        join_data = json.loads(join_message)
        if join_data["type"] == "join":
            session = join_data["session"]
            if session not in connected_clients:
                connected_clients[session] = []
            connected_clients[session].append(websocket)
            print(f"New client connected to session: {session}")

        # Handle incoming messages from the client
        async for message in websocket:
            await handle_message(websocket, session, message)

    except websockets.exceptions.ConnectionClosed:
        print(f"Client disconnected from session: {session}")
    finally:
        # Remove the websocket from the session on disconnect
        if session in connected_clients and websocket in connected_clients[session]:
            connected_clients[session].remove(websocket)
            if not connected_clients[session]:  # Clean up empty sessions
                del connected_clients[session]

async def main():
    # Start the WebSocket server
    start_server = websockets.serve(connection_handler, "0.0.0.0", 6789)
    print("WebSocket server started on ws://0.0.0.0:6789")

    # Run the WebSocket server forever
    await start_server

# Explicitly start the asyncio event loop
if __name__ == "__main__":
    asyncio.run(main())