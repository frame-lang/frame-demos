using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

using SystemParams;
#nullable disable
namespace csharp.systemparams
{

class SystemParams
{

    private SystemParamsCompartment _compartment_;
    private SystemParamsCompartment _nextCompartment_;


    public SystemParams(string stateMsg, string enterMsg)
    {


        // Create and intialize start state compartment.


        this._state_ = (int)SystemParamsState.BEGIN;
        this._compartment_ = new SystemParamsCompartment(this._state_);
        this._nextCompartment_ = null;
        this._compartment_.StateArgs["stateMsg"] = stateMsg;
        this._compartment_.EnterArgs["enterMsg"] = enterMsg;


        // Send system start event
        FrameEvent frameEvent = new FrameEvent(">", this._compartment_.EnterArgs);
        this._mux_(frameEvent);

    }

    // states enum
    private enum SystemParamsState
    {

        BEGIN = 0
    }
    //====================== Multiplexer ====================//

    private void _mux_(FrameEvent e)
    {
        switch (this._compartment_.state)
        {
            case (int)SystemParamsState.BEGIN:
                this._sBegin_(e);
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

    SystemParamsCompartment compartment;


    private void _sBegin_(FrameEvent e)
    {
        if (e._message == ">")
        {
            this.print_do(((string)this._compartment_.StateArgs["stateMsg"]) + " " + ((string)e._parameters["enterMsg"]));
            return;
        }
    }

    //===================== Actions Block ===================//

    public void print_do(string msg)
    {

        Console.WriteLine(msg);

    }

    // Unimplemented Actions



    //=============== Machinery and Mechanisms ==============//

    private int _state_;

    private void _transition_(SystemParamsCompartment compartment)
    {
        this._nextCompartment_ = compartment;
    }

    private void _doTransition_(SystemParamsCompartment nextCompartment)
    {
        this._mux_(new FrameEvent("<", this._compartment_.ExitArgs));
        this._compartment_ = nextCompartment;
        this._mux_(new FrameEvent(">", this._compartment_.EnterArgs));
    }

}

//=============== Compartment ==============//

class SystemParamsCompartment
{

    public int state;

    public SystemParamsCompartment(int state)
    {
        this.state = state;
    }

    public Dictionary<string, object> StateArgs { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> StateVars { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> EnterArgs { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> ExitArgs { get; set; } = new Dictionary<string, object>();
    public FrameEvent _forwardEvent = new FrameEvent();
}

class SystemParamsController : SystemParams
{
        public SystemParamsController(string stateMsg, string enterMsg) : base(stateMsg, enterMsg)
        {
        }
}

}