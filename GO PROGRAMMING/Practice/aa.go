import RPi.GPIO as GPIO
import time

# Set up GPIO pins
GPIO.setmode(GPIO.BCM)
GPIO.setwarnings(False)
GPIO.setup(18, GPIO.OUT)
GPIO.setup(23, GPIO.OUT)

# Toggle the two LEDs
while True:
    GPIO.output(18, GPIO.HIGH)
    GPIO.output(23, GPIO.LOW)
    time.sleep(1)
    GPIO.output(18, GPIO.LOW)
    GPIO.output(23, GPIO.HIGH)
    time.sleep(1)
