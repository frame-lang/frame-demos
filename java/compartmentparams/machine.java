import java.util.*;

// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files

enum CompartmentParamsState {
	S0(0), S1(1);

	public final int value;

	private CompartmentParamsState(int value) {
		this.value = value;
	}

	public int getValue() {
		return value;
	}
}

class CompartmentParams {
	private CompartmentParamsCompartment _compartment_;
	private CompartmentParamsCompartment _nextCompartment_;
	    
    CompartmentParams(int state_param, int enter_param) {

        // Create and intialize start state compartment.

        this._compartment_=new CompartmentParamsCompartment(this._state_);
        this._nextCompartment_=null;
        this._compartment_.stateArgs.put("state_param", state_param);
        this._compartment_.stateVars.put("state_var", 100);
        this._compartment_.enterArgs.put("enter_param", enter_param);

        // Initialize domain

        // Send system start event
        FrameEvent frameEvent = new FrameEvent(">", this._compartment_.enterArgs);
        this._mux_(frameEvent);
    }
    
    //====================== Multiplexer ====================//
    
    private void _mux_(FrameEvent e){
		
        if(this._compartment_.state == CompartmentParamsState.S0.getValue()) {
            this._sS0_(e);
        } else if(this._compartment_.state == CompartmentParamsState.S1.getValue()) {
        	this._sS1_(e);
        }

        if(this._nextCompartment_ != null) {
            CompartmentParamsCompartment nextCompartment = this._nextCompartment_;
            this._nextCompartment_ = null;
            if(nextCompartment._forwardEvent != null &&
              nextCompartment._forwardEvent._message == ">") {
                this._mux_(new FrameEvent("<", this._compartment_.exitArgs));
                this._compartment_ = nextCompartment;
                this._mux_(nextCompartment._forwardEvent);
            } else {
                this._doTransition_(nextCompartment);
                if (nextCompartment._forwardEvent != null) {
                    this._mux_(nextCompartment._forwardEvent);
                }
            }
            nextCompartment._forwardEvent = null;
        }
    }
    
    //===================== Machine Block ===================//  //  see working demo at codepen.io: https://codepen.io/frame-lang/pen/e96568889939d3bc112b16528900c352
      //  try out on http://framepiler.frame-lang.org
       
    private void _sS0_(FrameEvent e) {
        if (e._message == ">") {
            this.print_do((this._compartment_.stateArgs.get("state_param")) + " " + (this._compartment_.stateVars.get("state_var")) + " " + (e._parameters.get("enter_param")));
            CompartmentParamsCompartment compartment = new CompartmentParamsCompartment(CompartmentParamsState.S1.getValue());
            
            compartment._forwardEvent=e;
            compartment.stateArgs.put("state_param", (int) this._compartment_.stateArgs.get("state_param")+20);
            compartment.stateVars.put("state_var", 200);
            
            this._transition_(compartment);
            return;
        }
    }
    
    private void _sS1_(FrameEvent e) {
    	if (e._message == ">") {
            this.print_do((this._compartment_.stateArgs.get("state_param")) + " " + (this._compartment_.stateVars.get("state_var")) + " " + (e._parameters.get("enter_param")));
            return;
        }
    }
    
    //===================== Actions Block ===================//
    
    protected void print_do(String s) { throw new UnsupportedOperationException(); }
    
    
    //=============== Machinery and Mechanisms ==============//
    
    private int _state_;
    
    private void _transition_(CompartmentParamsCompartment compartment) {
        this._nextCompartment_=compartment;
    }

    private void _doTransition_(CompartmentParamsCompartment nextCompartment) {
        this._mux_(new FrameEvent("<", this._compartment_.exitArgs));
        this._compartment_ = _nextCompartment_;
        this._mux_(new FrameEvent(">", this._compartment_.enterArgs));
    } 
    
}

//====================Compartment====================//

class CompartmentParamsCompartment {

	int state;

	CompartmentParamsCompartment(int state) {
		this.state = state;
	}

	HashMap<String, Object> stateArgs = new HashMap<String, Object>();
	HashMap<String, Object> stateVars = new HashMap<String, Object>();
	HashMap<String, Object> enterArgs = new HashMap<String, Object>();
	HashMap<String, Object> exitArgs = new HashMap<String, Object>();

	FrameEvent _forwardEvent = new FrameEvent();
}


class CompartmentParamsController extends CompartmentParams {

	CompartmentParamsController(int state_param, int enter_param) {
		super(state_param, enter_param);
	}
	
    protected void print_do(String s) {
    	System.out.println(s);
    }
}