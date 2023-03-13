using fizzbuzz;
#nullable disable
namespace csharp.fizzbuzz
{

class FizzBuzz
{

    private FizzBuzzCompartment _compartment_;
    private FizzBuzzCompartment _nextCompartment_;


    public FizzBuzz()
    {


        // Create and intialize start state compartment.


        this._state_ = (int)FizzBuzzState.BEGIN;
        this._compartment_ = new FizzBuzzCompartment(this._state_);
        this._nextCompartment_ = null;


        // Send system start event
        FrameEvent frameEvent = new FrameEvent(">", null);
        this._mux_(frameEvent);

    }

    // states enum
    private enum FizzBuzzState
    {

        BEGIN = 0,
        FIZZ = 1,
        BUZZ = 2,
        DIGIT = 3,
        END = 4
    }
    //====================== Multiplexer ====================//

    private void _mux_(FrameEvent e)
    {
        switch (this._compartment_.state)
        {
            case (int)FizzBuzzState.BEGIN:
                this._sBegin_(e);
                break;
            case (int)FizzBuzzState.FIZZ:
                this._sFizz_(e);
                break;
            case (int)FizzBuzzState.BUZZ:
                this._sBuzz_(e);
                break;
            case (int)FizzBuzzState.DIGIT:
                this._sDigit_(e);
                break;
            case (int)FizzBuzzState.END:
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

    public void start() {
        FrameEvent e = new FrameEvent(">>",null);
        this._mux_(e);
    }


    //===================== Machine Block ===================//

    FizzBuzzCompartment compartment;


    private void _sBegin_(FrameEvent e)
    {
        if (e._message == ">>")
        {

            // start

            compartment =  new FizzBuzzCompartment((int)FizzBuzzState.FIZZ);

            compartment.EnterArgs["i"]  = 1;

            this._transition_(compartment);
            return;
        }
    }

    private void _sFizz_(FrameEvent e)
    {
        if (e._message == ">")
        {
            if (this.gt_100_do(((int)e._parameters["i"]))) {

                // i > 100

                compartment =  new FizzBuzzCompartment((int)FizzBuzzState.END);


                this._transition_(compartment);
                return;
            }
            if (this.mod3_eq0_do(((int)e._parameters["i"]))) {
                this.print_do("Fizz");

                // i % 3 == 0

                compartment =  new FizzBuzzCompartment((int)FizzBuzzState.BUZZ);

                compartment.EnterArgs["i"]  = (int)e._parameters["i"];
                compartment.EnterArgs["fizzed"]  = true;

                this._transition_(compartment);
            } else {

                // i % 3 != 0

                compartment =  new FizzBuzzCompartment((int)FizzBuzzState.BUZZ);

                compartment.EnterArgs["i"]  = (int)e._parameters["i"];
                compartment.EnterArgs["fizzed"]  = false;

                this._transition_(compartment);
            }
            return;
        }
    }

    private void _sBuzz_(FrameEvent e)
    {
        if (e._message == ">")
        {
            if (this.mod5_eq0_do(((int)e._parameters["i"]))) {
                this.print_do("Buzz");

                // i % 5 == 0
                this._compartment_.ExitArgs["output"] = " ";

                compartment =  new FizzBuzzCompartment((int)FizzBuzzState.FIZZ);

                compartment.EnterArgs["i"]  = this.plus_1_do((int)e._parameters["i"]);

                this._transition_(compartment);
                return;
            }
            if (((bool)e._parameters["fizzed"])) {

                // fizzed
                this._compartment_.ExitArgs["output"] = " ";

                compartment =  new FizzBuzzCompartment((int)FizzBuzzState.FIZZ);

                compartment.EnterArgs["i"]  = this.plus_1_do((int)e._parameters["i"]);

                this._transition_(compartment);
                return;
            }

            // ! mod3 | mod5
            this._compartment_.ExitArgs["output"] = "";

            compartment =  new FizzBuzzCompartment((int)FizzBuzzState.DIGIT);

            compartment.EnterArgs["i"]  = (int)e._parameters["i"];

            this._transition_(compartment);
            return;
        }
        else if (e._message == "<")
        {
            this.print_do(((string)e._parameters["output"]));
            return;
        }
    }

    private void _sDigit_(FrameEvent e)
    {
        if (e._message == ">")
        {
            this.print_do(((int)e._parameters["i"]).ToString());
            this.print_do(" ");

            // loop

            compartment =  new FizzBuzzCompartment((int)FizzBuzzState.FIZZ);

            compartment.EnterArgs["i"]  = this.plus_1_do((int)e._parameters["i"]);

            this._transition_(compartment);
            return;
        }
    }

    private void _sEnd_(FrameEvent e)
    {
    }

    //===================== Actions Block ===================//

    public void print_do(string msg)
    {
        Console.WriteLine(msg);
    }

    public bool gt_100_do(int i)
    {
        return i > 100;
    }

    public bool mod3_eq0_do(int i)
    {
        return i % 3 == 0;
    }

    public bool mod5_eq0_do(int i)
    {
        return i % 5 == 0;
    }

    public int plus_1_do(int i)
    {
        return i + 1;
    }

    // Unimplemented Actions



    //=============== Machinery and Mechanisms ==============//

    private int _state_;

    private void _transition_(FizzBuzzCompartment compartment)
    {
        this._nextCompartment_ = compartment;
    }

    private void _doTransition_(FizzBuzzCompartment nextCompartment)
    {
        this._mux_(new FrameEvent("<", this._compartment_.ExitArgs));
        this._compartment_ = nextCompartment;
        this._mux_(new FrameEvent(">", this._compartment_.EnterArgs));
    }

}

//=============== Compartment ==============//

class FizzBuzzCompartment
{

    public int state;

    public FizzBuzzCompartment(int state)
    {
        this.state = state;
    }

    public Dictionary<string, object> StateArgs { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> StateVars { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> EnterArgs { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> ExitArgs { get; set; } = new Dictionary<string, object>();
    public FrameEvent _forwardEvent = new FrameEvent();
}


class FizzBuzzController : FizzBuzz
{
        public FizzBuzzController() : base()
        {
        }
}

}