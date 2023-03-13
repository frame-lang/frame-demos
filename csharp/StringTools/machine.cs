using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

using StringTools;
#nullable disable
namespace csharp.stringtools
{

class StringTools
{

    private StringToolsCompartment _compartment_;
    private StringToolsCompartment _nextCompartment_;


    public StringTools()
    {


        // Create and intialize start state compartment.


        this._state_ = (int)StringToolsState.ROUTER;
        this._compartment_ = new StringToolsCompartment(this._state_);
        this._nextCompartment_ = null;


        // Send system start event
        FrameEvent frameEvent = new FrameEvent(">", null);
        this._mux_(frameEvent);

    }

    // states enum
    private enum StringToolsState
    {

        ROUTER = 0,
        REVERSE = 1,
        MAKEPALINDROME = 2
    }
    //====================== Multiplexer ====================//

    private void _mux_(FrameEvent e)
    {
        switch (this._compartment_.state)
        {
            case (int)StringToolsState.ROUTER:
                this._sRouter_(e);
                break;
            case (int)StringToolsState.REVERSE:
                this._sReverse_(e);
                break;
            case (int)StringToolsState.MAKEPALINDROME:
                this._sMakePalindrome_(e);
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

    public string reverse(string str) {
        Dictionary<string,object> parameters = new Dictionary<string,object>();
        parameters["str"] = str;

        FrameEvent e = new FrameEvent("reverse",parameters);
        this._mux_(e);
        return (string) e._return;
    }

    public string makePalindrome(string str) {
        Dictionary<string,object> parameters = new Dictionary<string,object>();
        parameters["str"] = str;

        FrameEvent e = new FrameEvent("makePalindrome",parameters);
        this._mux_(e);
        return (string) e._return;
    }


    //===================== Machine Block ===================//

    StringToolsCompartment compartment;


    private void _sRouter_(FrameEvent e)
    {
        if (e._message == "makePalindrome")
        {

            // make\npalindrome

            compartment =  new StringToolsCompartment((int)StringToolsState.MAKEPALINDROME);

            compartment._forwardEvent = e;

            this._transition_(compartment);
            return;
        }
        else if (e._message == "reverse")
        {

            // reverse

            compartment =  new StringToolsCompartment((int)StringToolsState.REVERSE);

            compartment._forwardEvent = e;

            this._transition_(compartment);
            return;
        }
    }

    private void _sReverse_(FrameEvent e)
    {
        if (e._message == "reverse")
        {
            e._return = this.reverse_str_do(((string)e._parameters["str"]));

            // ready

            compartment =  new StringToolsCompartment((int)StringToolsState.ROUTER);


            this._transition_(compartment);
            return;
        }
    }

    private void _sMakePalindrome_(FrameEvent e)
    {
        if (e._message == "makePalindrome")
        {
            e._return = ((string)e._parameters["str"]) + this.reverse_str_do(((string)e._parameters["str"]));

            // ready

            compartment =  new StringToolsCompartment((int)StringToolsState.ROUTER);


            this._transition_(compartment);
            return;
        }
    }

    //===================== Actions Block ===================//

    public string reverse_str_do(string str)
    {

    string result = "";
    foreach (char c in str)
    {
        result = c + result;
    }
    return result;

    }

    // Unimplemented Actions



    //=============== Machinery and Mechanisms ==============//

    private int _state_;

    private void _transition_(StringToolsCompartment compartment)
    {
        this._nextCompartment_ = compartment;
    }

    private void _doTransition_(StringToolsCompartment nextCompartment)
    {
        this._mux_(new FrameEvent("<", this._compartment_.ExitArgs));
        this._compartment_ = nextCompartment;
        this._mux_(new FrameEvent(">", this._compartment_.EnterArgs));
    }

}

//=============== Compartment ==============//

class StringToolsCompartment
{

    public int state;

    public StringToolsCompartment(int state)
    {
        this.state = state;
    }

    public Dictionary<string, object> StateArgs { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> StateVars { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> EnterArgs { get; set; } = new Dictionary<string, object>();
    public Dictionary<string, object> ExitArgs { get; set; } = new Dictionary<string, object>();
    public FrameEvent _forwardEvent = new FrameEvent();
}


class StringToolsController : StringTools
{
        public StringToolsController() : base()
        {
        }
}

}