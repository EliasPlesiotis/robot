from flask import Flask, redirect
from flask import request

import threading
import requests
import json
import time
import socket
#from pi2go import pi2go as pi


app = Flask(__name__)


@app.route("/connect", methods=['POST'])
def conn():
    if request.method == 'POST' :
        print('connected')
        return 'Hello world'


@app.route("/start", methods=['POST'])
def start():
    print(request)
    if request.method == 'POST':
        threading._start_new_thread(main, ())
        print('runing')
        return 'hello world'

@app.route("/stop", methods=['POST'])
def stop():
    print(request)
    if request.method == 'POST':
        #pi.cleanup()
        threading._start_new_thread(exit, ())
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
    

def main():
    moves = json.loads(requests.get('http://localhost:8080/files').text)

    for move in moves:
        if move['Dir'] == 'forward':
            forward(move['Duration'], move['Speed'])
            
        if move['Dir'] == 'backward':
            backward(move['Duration'], move['Speed'])
            
        if move['Dir'] == 'right':
            right(move['Duration'], move['Speed'])

        if move['Dir'] == 'left':
            left(move['Duration'], move['Speed'])
            
        print(move)


if __name__ == '__main__':
    #pi.init()

    app.run()
