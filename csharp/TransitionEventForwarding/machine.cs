using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

using TransitionEventForwarding;
#nullable disable
namespace csharp.transitioneventforwarding
{

class TransitionEventForwarding
{

    private TransitionEventForwardingCompartment _compartment_;
    private TransitionEventForwardingCompartment _nextCompartment_;


    public TransitionEventForwarding(int cycles)
    {


        // Create and intialize start state compartment.


        this._state_ = (int)TransitionEventForwardingState.START;
        this._compartment_ = new TransitionEventForwardingCompartment(this._state_);
        this._nextCompartment_ = null;
        this._compartment_.EnterArgs["cycles"] = cycles;


        // Send system start event
        FrameEvent frameEvent = new FrameEvent(">", this._compartment_.EnterArgs);
        this._mux_(frameEvent);

    }

    // states enum
    private enum TransitionEventForwardingState
    {

        START = 0,
        FORWARDEVENTAGAIN = 1,
        DECREMENT = 2,
        STOP = 3
    }
    //====================== Multiplexer ====================//

    private void _mux_(FrameEvent e)
    {
        switch (this._compartment_.state)
        {
            case (int)TransitionEventForwardingState.START:
                this._sStart_(e);
                break;
            case (int)TransitionEventForwardingState.FORWARDEVENTAGAIN:
                this._sForwardEventAgain_(e);
                break;
            case (int)TransitionEventForwardingState.DECREMENT:
                this._sDecrement_(e);
                break;
            case (int)TransitionEventForwardingState.STOP:
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

    TransitionEventForwardingCompartment compartment;


    private void _sStart_(FrameEvent e)
    {
        if (e._message == ">")
        {
            if (((int)e._parameters["cycles"]) == 0) {

                this._compartment_.ExitArgs["msg"] = "stopping";

                compartment =  new TransitionEventForwardingCompartment((int)TransitionEventForwardingState.STOP);

                compartment._forwardEvent = e;

                this._transition_(compartment);
                return;
            } else {

                this._compartment_.ExitArgs["msg"] = "keep going";

                compartment =  new TransitionEventForwardingCompartment((int)TransitionEventForwardingState.FORWARDEVENTAGAIN);

                compartment._forwardEvent = e;

                this._transition_(compartment);
            }
            return;
        }
        else if (e._message == "<")
        {
            this.print_do(((string)e._parameters["msg"]));
            return;
        }
    }

    private void _sForwardEventAgain_(FrameEvent e)
    {
        if (e._message == ">")
        {


            compartment =  new TransitionEventForwardingCompartment((int)TransitionEventForwardingState.DECREMENT);

            compartment._forwardEvent = e;

            this._transition_(compartment);
            return;
        }
    }

    private void _sDecrement_(FrameEvent e)
    {
        if (e._message == ">")
        {
            this.print_do(((int)e._parameters["cycles"]).ToString());


            compartment =  new TransitionEventForwardingCompartment((int)TransitionEventForwardingState.START);

            compartment.EnterArgs["cycles"]  = ((int)e._parameters["cycles"] - 1);

            this._transition_(compartment);
            return;
        }
    }

    private void _sStop_(FrameEvent e)
    {
        if (e._message == ">")
        {
            this.print_do(((int)e._parameters["cycles"]).ToString());
            this.print_do("done");
            return;
        }
    }

    //===================== Actions Block ===================//

    // Unimplemented Actions

    protected void print_do(string msg) {  
        Console.WriteLine(msg);
      }


    //=============== Machinery and Mechanisms ==============//

    private int _state_;

    private void _transition_(TransitionEventForwardingCompartment compartment)
    {
        this._nextCompartment_ = compartment;
    }

    private void _doTransition_(TransitionEventForwardingCompartment nextCompartment)
    {
        this._mux_(new FrameEvent("<", this._compartment_.ExitArgs));
        this._compartment_ = nextCompartment;
        this._mux_(new FrameEvent(">", this._compartment_.EnterArgs));
    }

}

//=============== Compartment ==============//

class TransitionEventForwardingCompartment
{

    public int state;

    public TransitionEventForwardingCompartment(int state)
    {
        this.state = state;
    }

    public Dictionary<string, object> StateArgs { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> StateVars { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> EnterArgs { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> ExitArgs { get; set; } = new Dictionary<string, object>();
    public FrameEvent _forwardEvent = new FrameEvent();
}


class TransitionEventForwardingController : TransitionEventForwarding
{
        public TransitionEventForwardingController(int cycles) : base(cycles)
        {
        }

    
}

}