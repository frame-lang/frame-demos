import java.util.*;

// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files

class SystemParams {
    
    private SystemParamsCompartment _compartment_;
    private SystemParamsCompartment _nextCompartment_;
    
    SystemParams(String stateMsg, String enterMsg) {
        
        // Create and intialize start state compartment.
        this._compartment_ = new SystemParamsCompartment(this._state_);
        this._nextCompartment_ = null;
        this._compartment_.stateArgs.put("stateMsg", stateMsg);
        this._compartment_.enterArgs.put("enterMsg", enterMsg);
        
        
        // Send system start event
        FrameEvent frame_event = new FrameEvent(">", this._compartment_.enterArgs);
        this._mux_(frame_event);
    
    }
    
    // states enum
    private enum SystemParamsState {
        
        BEGIN(0);
        
        public final int value;
        
        private SystemParamsState(int value) {
            this.value=value;
        }
        
        public int getValue() {
            return value;
        }
    }
    
    //====================== Multiplexer ====================//
    
    private void _mux_(FrameEvent e) {
        if(this._compartment_.state == SystemParamsState.BEGIN.getValue()) {
            this._sBegin_(e);
        }
        
        
        if(this._nextCompartment_ != null) {
            SystemParamsCompartment nextCompartment = this._nextCompartment_;
            this._nextCompartment_ = null;
            if(nextCompartment._forwardEvent != null && 
             nextCompartment._forwardEvent._message == ">") {
                this._mux_(new FrameEvent( "<", this._compartment_.exitArgs));
                this._compartment_ = nextCompartment;
                this._mux_(nextCompartment._forwardEvent);
            } else {
                this._doTransition_(nextCompartment);
                if(nextCompartment._forwardEvent != null) {
                    this._mux_(nextCompartment._forwardEvent);
                }
            }
            nextCompartment._forwardEvent = null;
        }
    }
    
    
    //===================== Machine Block ===================//
    
    private void _sBegin_(FrameEvent e) {
        if(e._message == ">") {
            print_do(((String) this._compartment_.stateArgs.get("stateMsg")) + " " + String.valueOf(((String) e._parameters.get("enterMsg"))));
            return;
        }
    }
    
    //===================== Actions Block ===================//
    
    protected void print_do(String msg) {
        
        System.out.println(msg);
    
    }
    
    // Unimplemented Actions
    
    
    //=============== Machinery and Mechanisms ==============//
    
    private int _state_;
    
    private void _transition_(SystemParamsCompartment compartment) {
        this._nextCompartment_ = compartment;
    }
    
    private void _doTransition_(SystemParamsCompartment nextCompartment) {
        this._mux_(new FrameEvent("<", this._compartment_.exitArgs));
        this._compartment_ = nextCompartment;
        this._mux_(new FrameEvent(">", this._compartment_.enterArgs));
    }
    
}

//=============== Compartment ==============//

class SystemParamsCompartment {

    int state;
    
    SystemParamsCompartment(int state) {
        this.state = state;
    }
    
    HashMap<String, Object> stateArgs = new HashMap<String, Object>();
    HashMap<String, Object> stateVars = new HashMap<String, Object>();
    HashMap<String, Object> enterArgs = new HashMap<String, Object>();
    HashMap<String, Object> exitArgs = new HashMap<String, Object>();
    FrameEvent _forwardEvent = new FrameEvent();
}


class SystemParamsController extends SystemParams {
	SystemParamsController(String stateMsg, String enterMsg) {
	  super(stateMsg, enterMsg);
	}
}
