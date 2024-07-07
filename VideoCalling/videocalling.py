from flask import Flask, request
from flask_socketio import SocketIO, emit, join_room

app = Flask(__name__)
app.secret_key = 'random secret key!'
socketio = SocketIO(app, cors_allowed_origins="*")


@socketio.on('join')
def join(message):
    username = message['username']
    room = message['room']
    join_room(room)
    print('RoomEvent: {} has joined the room {}\n'.format(username, room))
    emit('ready', {'username': username}, room=room, skip_sid=request.sid)


@socketio.on('data')
def transfer_data(message):
    print("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
    print(message)
    print("###################################")
    username = message['username']
    room = message['room']
    data = message['data']
    print('DataEvent: {} has sent the data:\n {}\n'.format(username, data))
    emit('data', data, room=room, skip_sid=request.sid)


if __name__ == '__main__':
    socketio.run(app, host="0.0.0.0", port=9000, ssl_context=('cert.pem', 'key.pem'))