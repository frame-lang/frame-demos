using Countdown;
#nullable disable
namespace csharp.countdown
{

class Countdown
{

    private CountdownCompartment _compartment_;
    private CountdownCompartment _nextCompartment_;


    public Countdown(int i)
    {


        // Create and intialize start state compartment.


        this._state_ = (int)CountdownState.S0;
        this._compartment_ = new CountdownCompartment(this._state_);
        this._nextCompartment_ = null;
        this._compartment_.StateArgs["i"] = i;
        this._compartment_.StateVars["dec"] = 1;


        // Send system start event
        FrameEvent frameEvent = new FrameEvent(">", null);
        this._mux_(frameEvent);

    }

    // states enum
    private enum CountdownState
    {

        S0 = 0,
        S1 = 1,
        STOP = 2
    }
    //====================== Multiplexer ====================//

    private void _mux_(FrameEvent e)
    {
        switch (this._compartment_.state)
        {
            case (int)CountdownState.S0:
                this._sS0_(e);
                break;
            case (int)CountdownState.S1:
                this._sS1_(e);
                break;
            case (int)CountdownState.STOP:
                this._sStop_(e);
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


    //===================== Machine Block ===================//

    CountdownCompartment compartment;


    private void _sS0_(FrameEvent e)
    {
        if (e._message == ">")
        {
            (this._compartment_.StateArgs["i"]) = ((int)this._compartment_.StateArgs["i"]) - ((int)this._compartment_.StateVars["dec"]);
            this.print_do(((int)this._compartment_.StateArgs["i"]).ToString());
            if (((int)this._compartment_.StateArgs["i"]) == 0) {


                compartment =  new CountdownCompartment((int)CountdownState.STOP);


                this._transition_(compartment);
                return;
            }


            compartment =  new CountdownCompartment((int)CountdownState.S1);

            compartment.EnterArgs["i"]  = (int)this._compartment_.StateArgs["i"];

            this._transition_(compartment);
            return;
        }
    }

    private void _sS1_(FrameEvent e)
    {
        if (e._message == ">")
        {


            compartment =  new CountdownCompartment((int)CountdownState.S0);

            compartment.StateArgs["i"] = (int)e._parameters["i"];

            compartment.StateVars["dec"] = 1;

            this._transition_(compartment);
            return;
        }
    }

    private void _sStop_(FrameEvent e)
    {
        if (e._message == ">")
        {
            this.print_do("done");
            return;
        }
    }

    //===================== Actions Block ===================//

    public void print_do(String s)
    {
        Console.WriteLine(s);
    }

    // Unimplemented Actions



    //=============== Machinery and Mechanisms ==============//

    private int _state_;

    private void _transition_(CountdownCompartment compartment)
    {
        this._nextCompartment_ = compartment;
    }

    private void _doTransition_(CountdownCompartment nextCompartment)
    {
        this._mux_(new FrameEvent("<", this._compartment_.ExitArgs));
        this._compartment_ = nextCompartment;
        this._mux_(new FrameEvent(">", this._compartment_.EnterArgs));
    }

}

//=============== Compartment ==============//

class CountdownCompartment
{

    public int state;

    public CountdownCompartment(int state)
    {
        this.state = state;
    }

    public Dictionary<string, object> StateArgs { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> StateVars { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> EnterArgs { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> ExitArgs { get; set; } = new Dictionary<string, object>();
    public FrameEvent _forwardEvent = new FrameEvent();
}


class CountdownController : Countdown
{
        public CountdownController(int i) : base(i)
        {
        }
}

}