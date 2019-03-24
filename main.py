from flask import Flask, redirect
from flask_cors import CORS, cross_origin
from flask import request

import threading
import requests
import json
import time
import socket
#from pi2go import pi2go as pi


app = Flask(__name__)
cors = CORS(app, resources={r"/*": {"origins": "*", "headers": "content-type", "methods": ["POST", "GET"]}})


@app.route("/connect", methods=['POST'])
@cross_origin()
def conn():
    if request.method == 'POST' :
        print('connected')
        return 'hello world'


@app.route("/start/<file>", methods=['POST'])
@cross_origin()
def start(file):
    print(request)
    if request.method == 'POST':
        threading._start_new_thread(blocks, (file,))
        print('running')
        return 'hello world'

@app.route("/controller/<dir>+<speed>", methods=['POST'])
def controller(dir, speed):
    if dir == 'forward':
        #pi.forward(speed)
        pass
    elif dir == 'backward':
        #pi.reverse(speed)
        pass
    elif dir == 'left':
        #pi.spinLeft(speed)
        pass
    elif dir == 'backward':
        #pi.spinRight(speed)
        pass
    elif dir == 'stop':
        #pi.stop()
        pass
    
    print(dir, speed)
    return 'hello world'

@app.route("/stop", methods=['POST'])
def stop():
    print(request)
    if request.method == 'POST':
        #pi.stop()
        #pi.cleanup()
        #pi.init()
        return 'hello world'

@app.route("/code", methods=['GET'])
def code():
    threading._start_new_thread(python, ())
    print('runing')
    return 'hello world'


def forward(n, speed):
    #pi.forward(speed)
    time.sleep(n)
    #pi.stop()

def backward(n, speed):
    #pi.backward(speed)
    time.sleep(n)
    #pi.stop()

def left(n, speed):
    #pi.spinLeft(speed)
    time.sleep(n)
    #pi.stop()

def right(n, speed):
    #pi.spinRight(speed)
    time.sleep(n)
    #pi.stop()
    

def blocks(file):
    moves = json.loads(requests.get('http://192.168.1.25:3000/file/{' + file + '}').text)

    for move in moves:
        if move['Diraction'] == 'forward':
            forward(move['Duration'], move['Speed'])
            
        if move['Diraction'] == 'backward':
            backward(move['Duration'], move['Speed'])
            
        if move['Diraction'] == 'right':
            right(move['Duration'], move['Speed'])

        if move['Diraction'] == 'left':
            left(move['Duration'], move['Speed'])
            
        print(move)

def python():
    source = json.loads(requests.get('http://192.168.1.25:8080/view'))
    compiled = compile(source)
    exec(compiled)

if __name__ == '__main__':
    #pi.init()

    app.run(host='')
