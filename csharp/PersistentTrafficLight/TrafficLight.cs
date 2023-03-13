using PersistentTrafficLight;
#nullable disable
namespace csharp.persistenttrafficlight
{

class TrafficLight
{

    private TrafficLightCompartment _compartment_;
    private TrafficLightCompartment _nextCompartment_;

    protected TrafficLightMom _manager;

    public TrafficLight(TrafficLightMom manager){


        // Create and intialize start state compartment.
        this._manager = manager;


        this._state_ = (int)TrafficLightState.BEGIN;
        this._compartment_ = new TrafficLightCompartment(this._state_);
        this._nextCompartment_ = null;


        // Send system start event
        FrameEvent frameEvent = new FrameEvent(">", null);
        this._mux_(frameEvent);

    }

    // states enum
    private enum TrafficLightState
    {

        BEGIN = 0,
        RED = 1,
        GREEN = 2,
        YELLOW = 3,
        FLASHINGRED = 4,
        END = 5,
        WORKING = 6
    }

    public static TrafficLight loadTrafficLight(TrafficLightMomController manager, TrafficLightController data)
    {

        data._manager = manager;
        return data;

    }

    public TrafficLight marshal()
    {

        TrafficLight data = this;
        return data;

    }

    //====================== Multiplexer ====================//

    private void _mux_(FrameEvent e)
    {
        switch (this._compartment_.state)
        {
            case (int)TrafficLightState.BEGIN:
                this._sBegin_(e);
                break;
            case (int)TrafficLightState.RED:
                this._sRed_(e);
                break;
            case (int)TrafficLightState.GREEN:
                this._sGreen_(e);
                break;
            case (int)TrafficLightState.YELLOW:
                this._sYellow_(e);
                break;
            case (int)TrafficLightState.FLASHINGRED:
                this._sFlashingRed_(e);
                break;
            case (int)TrafficLightState.END:
                this._sEnd_(e);
                break;
            case (int)TrafficLightState.WORKING:
                this._sWorking_(e);
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

    public void systemError() {
        FrameEvent e = new FrameEvent("systemError",null);
        this._mux_(e);
    }

    public void systemRestart() {
        FrameEvent e = new FrameEvent("systemRestart",null);
        this._mux_(e);
    }


    //===================== Machine Block ===================//

    TrafficLightCompartment compartment;


    private void _sBegin_(FrameEvent e)
    {
        if (e._message == ">")
        {
            this.startWorkingTimer_do();


            compartment =  new TrafficLightCompartment((int)TrafficLightState.RED);

            
            this._transition_(compartment);
            return;
        }
    }

    private void _sRed_(FrameEvent e)
    {
        if (e._message == ">")
        {
            this.enterRed_do();
            return;
        }
        else if (e._message == "tick")
        {


            compartment =  new TrafficLightCompartment((int)TrafficLightState.GREEN);


            this._transition_(compartment);
            return;
        }
        _sWorking_(e);

    }

    private void _sGreen_(FrameEvent e)
    {
        if (e._message == ">")
        {
            this.enterGreen_do();
            return;
        }
        else if (e._message == "tick")
        {


            compartment =  new TrafficLightCompartment((int)TrafficLightState.YELLOW);


            this._transition_(compartment);
            return;
        }
        _sWorking_(e);

    }

    private void _sYellow_(FrameEvent e)
    {
        if (e._message == ">")
        {
            this.enterYellow_do();
            return;
        }
        else if (e._message == "tick")
        {


            compartment =  new TrafficLightCompartment((int)TrafficLightState.RED);


            this._transition_(compartment);
            return;
        }
        _sWorking_(e);

    }

    private void _sFlashingRed_(FrameEvent e)
    {
        if (e._message == ">")
        {
            this.enterFlashingRed_do();
            this.stopWorkingTimer_do();
            this.startFlashingTimer_do();
            return;
        }
        else if (e._message == "<")
        {
            this.exitFlashingRed_do();
            this.stopFlashingTimer_do();
            this.startWorkingTimer_do();
            return;
        }
        else if (e._message == "tick")
        {
            this.changeFlashingAnimation_do();
            return;
        }
        else if (e._message == "systemRestart")
        {


            compartment =  new TrafficLightCompartment((int)TrafficLightState.RED);


            this._transition_(compartment);
            return;
        }
        else if (e._message == "stop")
        {


            compartment =  new TrafficLightCompartment((int)TrafficLightState.END);


            this._transition_(compartment);
            return;
        }
    }

    private void _sEnd_(FrameEvent e)
    {
        if (e._message == ">")
        {
            this.stopWorkingTimer_do();
            return;
        }
    }

    private void _sWorking_(FrameEvent e)
    {
        if (e._message == "stop")
        {


            compartment =  new TrafficLightCompartment((int)TrafficLightState.END);


            this._transition_(compartment);
            return;
        }
        else if (e._message == "systemError")
        {


            compartment =  new TrafficLightCompartment((int)TrafficLightState.FLASHINGRED);


            this._transition_(compartment);
            return;
        }
    }

    //===================== Actions Block ===================//

    public void enterRed_do()
    {
        this._manager.enterRed_do();
    }

    public void enterGreen_do()
    {
        this._manager.enterGreen_do();
    }

    public void enterYellow_do()
    {
        this._manager.enterYellow_do();
    }

    public void startWorkingTimer_do()
    {
        return;
    }

    // Unimplemented Actions

    protected void enterFlashingRed_do() {  throw new NotImplementedException();  }
    protected void exitFlashingRed_do() {  throw new NotImplementedException();  }
    protected void stopWorkingTimer_do() {  throw new NotImplementedException();  }
    protected void startFlashingTimer_do() {  throw new NotImplementedException();  }
    protected void stopFlashingTimer_do() {  throw new NotImplementedException();  }
    protected void changeColor_do(string color) {  throw new NotImplementedException();  }
    protected void startFlashing_do() {  throw new NotImplementedException();  }
    protected void stopFlashing_do() {  throw new NotImplementedException();  }
    protected void changeFlashingAnimation_do() {  throw new NotImplementedException();  }
    protected void log_do(string msg) {  throw new NotImplementedException();  }

    //===================== Domain Block ===================//

    private string flashColor = "";


    //=============== Machinery and Mechanisms ==============//

    private int _state_;

    private void _transition_(TrafficLightCompartment compartment)
    {
        this._nextCompartment_ = compartment;
    }

    private void _doTransition_(TrafficLightCompartment nextCompartment)
    {
        this._mux_(new FrameEvent("<", this._compartment_.ExitArgs));
        this._compartment_ = nextCompartment;
        this._mux_(new FrameEvent(">", this._compartment_.EnterArgs));
    }

}

//=============== Compartment ==============//

class TrafficLightCompartment
{

    public int state;

    public TrafficLightCompartment(int state)
    {
        this.state = state;
    }

    public Dictionary<string, object> StateArgs { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> StateVars { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> EnterArgs { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> ExitArgs { get; set; } = new Dictionary<string, object>();
    public FrameEvent _forwardEvent = new FrameEvent();
}


class TrafficLightController : TrafficLight
{
        public TrafficLightController(TrafficLightMom manager) : base(manager)
        {
        }

    protected new void enterFlashingRed_do() {}

    protected new void exitFlashingRed_do() {}

    protected new void stopWorkingTimer_do() {}

    protected new void startFlashingTimer_do() {}

    protected new void stopFlashingTimer_do() {}

    protected new void changeColor_do(string color) {}

    protected new void startFlashing_do() {}

    protected new void stopFlashing_do() {}

    protected new void changeFlashingAnimation_do() {}

    protected new void log_do(string msg) {}
}

}