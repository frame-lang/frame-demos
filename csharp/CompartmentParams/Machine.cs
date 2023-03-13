// See https://aka.ms/new-console-template for more information
using CompartmentParams;
#nullable disable
namespace csharp.compartmentparams
{

class CompartmentParams
{

    private CompartmentParamsCompartment _compartment_;
    private CompartmentParamsCompartment _nextCompartment_;


    public CompartmentParams(int state_param, int enter_param)
    {


        // Create and intialize start state compartment.


        this._state_ = (int)CompartmentParamsState.S0;
        this._compartment_ = new CompartmentParamsCompartment(this._state_);
        this._nextCompartment_ = null;
        this._compartment_.StateArgs["state_param"] = state_param;
        this._compartment_.StateVars["state_var"] = 100;
        this._compartment_.EnterArgs["enter_param"] = enter_param;


        // Send system start event
        FrameEvent frameEvent = new FrameEvent(">", this._compartment_.EnterArgs);
        this._mux_(frameEvent);

    }

    // states enum
    private enum CompartmentParamsState
    {

        S0 = 0,
        S1 = 1
    }
    //====================== Multiplexer ====================//

    private void _mux_(FrameEvent e)
    {
        switch (this._compartment_.state)
        {
            case (int)CompartmentParamsState.S0:
                this._sS0_(e);
                break;
            case (int)CompartmentParamsState.S1:
                this._sS1_(e);
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

    CompartmentParamsCompartment compartment;


    private void _sS0_(FrameEvent e)
    {
        if (e._message == ">")
        {
            this.print_do(((int)this._compartment_.StateArgs["state_param"]) + " " + ((int)this._compartment_.StateVars["state_var"]) + " " + ((int)e._parameters["enter_param"]));


            compartment =  new CompartmentParamsCompartment((int)CompartmentParamsState.S1);

            compartment._forwardEvent = e;
            compartment.StateArgs["state_param"] = (int)this._compartment_.StateArgs["state_param"] + 20;

            compartment.StateVars["state_var"] = 200;

            this._transition_(compartment);
            return;
        }
    }

    private void _sS1_(FrameEvent e)
    {
        if (e._message == ">")
        {
            this.print_do(((int)this._compartment_.StateArgs["state_param"]) + " " + ((int)this._compartment_.StateVars["state_var"]) + " " + ((int)e._parameters["enter_param"]));
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

    private void _transition_(CompartmentParamsCompartment compartment)
    {
        this._nextCompartment_ = compartment;
    }

    private void _doTransition_(CompartmentParamsCompartment nextCompartment)
    {
        this._mux_(new FrameEvent("<", this._compartment_.ExitArgs));
        this._compartment_ = nextCompartment;
        this._mux_(new FrameEvent(">", this._compartment_.EnterArgs));
    }

}

//=============== Compartment ==============//

class CompartmentParamsCompartment
{

    public int state;

    public CompartmentParamsCompartment(int state)
    {
        this.state = state;
    }

    public Dictionary<string, object> StateArgs { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> StateVars { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> EnterArgs { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> ExitArgs { get; set; } = new Dictionary<string, object>();
    public FrameEvent _forwardEvent = new FrameEvent();
}


class CompartmentParamsController : CompartmentParams
{
        public CompartmentParamsController(int state_param, int enter_param) : base(state_param, enter_param)
        {
        }
}
}

