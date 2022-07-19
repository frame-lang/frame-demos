
import threading
import time

from trafficlight.trafficlightMom import TrafficLightManagerController

# current Time in seconds
start_time = time.time()

manager = TrafficLightManagerController()

event = threading.Event()

while not event.wait(1):
    #Exit the loop after 10 second
    if(time.time() - start_time  >= 5):
        manager.stop()
        print("Finished")
        event.clear()
        break
    print("Tick")
    manager.tick()
    


