using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using PersistentTrafficLight;
#nullable disable
namespace csharp.persistenttrafficlight

{

class TrafficLightMom
{

    private TrafficLightMomCompartment _compartment_;
    private TrafficLightMomCompartment _nextCompartment_;



    public TrafficLightMom(){


        // Create and intialize start state compartment.


        this._state_ = (int)TrafficLightMomState.NEW;
        this._compartment_ = new TrafficLightMomCompartment(this._state_);
        this._nextCompartment_ = null;


        // Send system start event
        FrameEvent frameEvent = new FrameEvent(">", null);
        this._mux_(frameEvent);

    }

    // states enum
    private enum TrafficLightMomState
    {

        NEW = 0,
        SAVING = 1,
        PERSISTED = 2,
        WORKING = 3,
        TRAFFICLIGHTAPI = 4,
        END = 5
    }
    //====================== Multiplexer ====================//

    private void _mux_(FrameEvent e)
    {
        switch (this._compartment_.state)
        {
            case (int)TrafficLightMomState.NEW:
                this._sNew_(e);
                break;
            case (int)TrafficLightMomState.SAVING:
                this._sSaving_(e);
                break;
            case (int)TrafficLightMomState.PERSISTED:
                this._sPersisted_(e);
                break;
            case (int)TrafficLightMomState.WORKING:
                this._sWorking_(e);
                break;
            case (int)TrafficLightMomState.TRAFFICLIGHTAPI:
                this._sTrafficLightApi_(e);
                break;
            case (int)TrafficLightMomState.END:
                this._sEnd_(e);
                break;
        }

        if( this._nextCompartment_ != null)
        {
            var nextCompartment = this._nextCompartment_;
            this._nextCompartment_ = null;
            if (nextCompartment._forwardEvent != null &&
               nextCompartment._forwardEvent._message == ">")
            {
                this._mux_(new FrameEvent( "<", this._compartment_.ExitArgs));
                this._compartment_ = nextCompartment;
                this._mux_(nextCompartment._forwardEvent);
            }
            else
            {
                this._doTransition_(nextCompartment);
                if (nextCompartment._forwardEvent != null)
                {
                    this._mux_(nextCompartment._forwardEvent);
                }
            }
            nextCompartment._forwardEvent = null;
        }
    }

    //===================== Interface Block ===================//

    public void stop() {
        FrameEvent e = new FrameEvent("stop",null);
        this._mux_(e);
    }

    public void tick() {
        FrameEvent e = new FrameEvent("tick",null);
        this._mux_(e);
    }

    public void enterRed() {
        FrameEvent e = new FrameEvent("enterRed",null);
        this._mux_(e);
    }

    public void enterGreen() {
        FrameEvent e = new FrameEvent("enterGreen",null);
        this._mux_(e);
    }

    public void enterYellow() {
        FrameEvent e = new FrameEvent("enterYellow",null);
        this._mux_(e);
    }

    public void enterFlashingRed() {
        FrameEvent e = new FrameEvent("enterFlashingRed",null);
        this._mux_(e);
    }

    public void exitFlashingRed() {
        FrameEvent e = new FrameEvent("exitFlashingRed",null);
        this._mux_(e);
    }

    public void startWorkingTimer() {
        FrameEvent e = new FrameEvent("startWorkingTimer",null);
        this._mux_(e);
    }

    public void stopWorkingTimer() {
        FrameEvent e = new FrameEvent("stopWorkingTimer",null);
        this._mux_(e);
    }

    public void startFlashingTimer() {
        FrameEvent e = new FrameEvent("startFlashingTimer",null);
        this._mux_(e);
    }

    public void stopFlashingTimer() {
        FrameEvent e = new FrameEvent("stopFlashingTimer",null);
        this._mux_(e);
    }

    public void changeColor(string color) {
        Dictionary<string,object> parameters = new Dictionary<string,object>();
        parameters["color"] = color;

        FrameEvent e = new FrameEvent("changeColor",parameters);
        this._mux_(e);
    }

    public void startFlashing() {
        FrameEvent e = new FrameEvent("startFlashing",null);
        this._mux_(e);
    }

    public void stopFlashing() {
        FrameEvent e = new FrameEvent("stopFlashing",null);
        this._mux_(e);
    }

    public void changeFlashingAnimation() {
        FrameEvent e = new FrameEvent("changeFlashingAnimation",null);
        this._mux_(e);
    }

    public void systemError() {
        FrameEvent e = new FrameEvent("systemError",null);
        this._mux_(e);
    }

    public void systemRestart() {
        FrameEvent e = new FrameEvent("systemRestart",null);
        this._mux_(e);
    }

    public void log(string msg) {
        Dictionary<string,object> parameters = new Dictionary<string,object>();
        parameters["msg"] = msg;

        FrameEvent e = new FrameEvent("log",parameters);
        this._mux_(e);
    }


    //===================== Machine Block ===================//

    TrafficLightMomCompartment compartment;


    private void _sNew_(FrameEvent e)
    {
        if (e._message == ">")
        {
            this._trafficLight = new TrafficLightController(this);

            // Traffic Light\nStarted

            compartment =  new TrafficLightMomCompartment((int)TrafficLightMomState.SAVING);


            this._transition_(compartment);
            return;
        }
        _sTrafficLightApi_(e);

    }

    private void _sSaving_(FrameEvent e)
    {
        if (e._message == ">")
        {
            this._data = this._trafficLight.marshal();
            this._trafficLight = null;

            // Saved

            compartment =  new TrafficLightMomCompartment((int)TrafficLightMomState.PERSISTED);


            this._transition_(compartment);
            return;
        }
    }

    private void _sPersisted_(FrameEvent e)
    {
        if (e._message == "tick")
        {

            // Tick

            compartment =  new TrafficLightMomCompartment((int)TrafficLightMomState.WORKING);

            compartment._forwardEvent = e;

            this._transition_(compartment);
            return;
        }
        else if (e._message == "systemError")
        {

            // System Error

            compartment =  new TrafficLightMomCompartment((int)TrafficLightMomState.WORKING);

            compartment._forwardEvent = e;

            this._transition_(compartment);
            return;
        }
        else if (e._message == "stop")
        {

            // Stop

            compartment =  new TrafficLightMomCompartment((int)TrafficLightMomState.END);


            this._transition_(compartment);
            return;
        }
    }

    private void _sWorking_(FrameEvent e)
    {
        if (e._message == ">")
        {
            this._trafficLight = TrafficLightController.loadTrafficLight((TrafficLightMomController) this, (TrafficLightController) this._data);
            return;
        }
        else if (e._message == "tick")
        {
            this._trafficLight.tick();

            // Done

            compartment =  new TrafficLightMomCompartment((int)TrafficLightMomState.SAVING);


            this._transition_(compartment);
            return;
        }
        else if (e._message == "systemError")
        {
            this.systemError_do();

            // Done

            compartment =  new TrafficLightMomCompartment((int)TrafficLightMomState.SAVING);


            this._transition_(compartment);
            return;
        }
        _sTrafficLightApi_(e);

    }

    private void _sTrafficLightApi_(FrameEvent e)
    {
        if (e._message == "enterRed")
        {
            this.enterRed_do();
            return;
        }
        else if (e._message == "enterGreen")
        {
            this.enterGreen_do();
            return;
        }
        else if (e._message == "enterYellow")
        {
            this.enterYellow_do();
            return;
        }
        else if (e._message == "enterFlashingRed")
        {
            this.enterFlashingRed_do();
            return;
        }
        else if (e._message == "exitFlashingRed")
        {
            this.exitFlashingRed_do();
            return;
        }
        else if (e._message == "startWorkingTimer")
        {
            this.startWorkingTimer_do();
            return;
        }
        else if (e._message == "stopWorkingTimer")
        {
            this.stopWorkingTimer_do();
            return;
        }
        else if (e._message == "startFlashingTimer")
        {
            this.startFlashingTimer_do();
            return;
        }
        else if (e._message == "stopFlashingTimer")
        {
            this.stopFlashingTimer_do();
            return;
        }
        else if (e._message == "changeColor")
        {
            this.changeColor_do(((string)e._parameters["color"]));
            return;
        }
        else if (e._message == "startFlashing")
        {
            this.startFlashing_do();
            return;
        }
        else if (e._message == "stopFlashing")
        {
            this.stopFlashing_do();
            return;
        }
        else if (e._message == "changeFlashingAnimation")
        {
            this.changeFlashingAnimation_do();
            return;
        }
        else if (e._message == "systemError")
        {
            this.systemError_do();
            return;
        }
        else if (e._message == "systemRestart")
        {
            this.systemRestart_do();
            return;
        }
        else if (e._message == "log")
        {
            this.log_do(((string)e._parameters["msg"]));
            return;
        }
    }

    private void _sEnd_(FrameEvent e)
    {
        if (e._message == ">")
        {
            this._trafficLight = TrafficLightController.loadTrafficLight((TrafficLightMomController) this, (TrafficLightController) this._data);
            this.trafficLight.stop();
            this.trafficLight = null;
            return;
        }
        _sTrafficLightApi_(e);

    }

    //===================== Actions Block ===================//

    public void enterRed_do()
    {
        Console.WriteLine("Red");
    }

    public void enterGreen_do()
    {
        Console.WriteLine("Green");
    }

    public void enterYellow_do()
    {
        Console.WriteLine("Yellow");
    }

    // Unimplemented Actions

    protected void enterFlashingRed_do() {  throw new NotImplementedException();  }
    protected void exitFlashingRed_do() {  throw new NotImplementedException();  }
    protected void startWorkingTimer_do() {  throw new NotImplementedException();  }
    protected void stopWorkingTimer_do() {  throw new NotImplementedException();  }
    protected void startFlashingTimer_do() {  throw new NotImplementedException();  }
    protected void stopFlashingTimer_do() {  throw new NotImplementedException();  }
    protected void changeColor_do(string color) {  throw new NotImplementedException();  }
    protected void startFlashing_do() {  throw new NotImplementedException();  }
    protected void stopFlashing_do() {  throw new NotImplementedException();  }
    protected void changeFlashingAnimation_do() {  throw new NotImplementedException();  }
    protected void systemError_do() {  throw new NotImplementedException();  }
    protected void systemRestart_do() {  throw new NotImplementedException();  }
    protected void log_do(string msg) {  throw new NotImplementedException();  }

    //===================== Domain Block ===================//

    private TrafficLightMomController trafficLight = null;
    private TrafficLight _data = null;
    private TrafficLight _trafficLight = null;


    //=============== Machinery and Mechanisms ==============//

    private int _state_;

    private void _transition_(TrafficLightMomCompartment compartment)
    {
        this._nextCompartment_ = compartment;
    }

    private void _doTransition_(TrafficLightMomCompartment nextCompartment)
    {
        this._mux_(new FrameEvent("<", this._compartment_.ExitArgs));
        this._compartment_ = nextCompartment;
        this._mux_(new FrameEvent(">", this._compartment_.EnterArgs));
    }

}

//=============== Compartment ==============//

class TrafficLightMomCompartment
{

    public int state;

    public TrafficLightMomCompartment(int state)
    {
        this.state = state;
    }

    public Dictionary<string, object> StateArgs { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> StateVars { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> EnterArgs { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> ExitArgs { get; set; } = new Dictionary<string, object>();
    public FrameEvent _forwardEvent = new FrameEvent();
}


class TrafficLightMomController : TrafficLightMom
{
        public TrafficLightMomController() : base()
        {
        }

    protected new void enterFlashingRed_do() {}

    protected new void exitFlashingRed_do() {}

    protected new void startWorkingTimer_do() {}

    protected new void stopWorkingTimer_do() {}

    protected new void startFlashingTimer_do() {}

    protected new void stopFlashingTimer_do() {}

    protected new void changeColor_do(string color) {}

    protected new void startFlashing_do() {}

    protected new void stopFlashing_do() {}

    protected new void changeFlashingAnimation_do() {}

    protected new void systemError_do() {}

    protected new void systemRestart_do() {}

    protected new void log_do(string msg) {}
}

}