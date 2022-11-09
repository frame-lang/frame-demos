// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files
import java.util.*;

class Countdown {
    
    private CountdownCompartment _compartment_;
    private CountdownCompartment _nextCompartment_;
    
    Countdown(int i) {
        
        // Create and intialize start state compartment.
        this._compartment_ = new CountdownCompartment(this._state_);
        this._nextCompartment_ = null;
        this._compartment_.stateArgs.put("i", i);
        this._compartment_.stateVars.put("dec", 1);
        
        
        // Send system start event
        FrameEvent frame_event = new FrameEvent(">", null);
        this._mux_(frame_event);
    
    }
    
    // states enum
    private enum CountdownState {
        
        S0(0),
        S1(1),
        STOP(2);
        
        public final int value;
        
        private CountdownState(int value) {
            this.value=value;
        }
        
        public int getValue() {
            return value;
        }
    }
    
    //====================== Multiplexer ====================//
    
    private void _mux_(FrameEvent e) {
        if(this._compartment_.state == CountdownState.S0.getValue()) {
            this._sS0_(e);
        }else if(this._compartment_.state == CountdownState.S1.getValue()) {
            this._sS1_(e);
        }else if(this._compartment_.state == CountdownState.STOP.getValue()) {
            this._sStop_(e);
        }
        
        
        if(this._nextCompartment_ != null) {
            CountdownCompartment nextCompartment = this._nextCompartment_;
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
    
    private void _sS0_(FrameEvent e) {
        if(e._message == ">") {
            
            int value = (int) this._compartment_.stateArgs.get("i") - (int) this._compartment_.stateVars.get("dec");
            this._compartment_.stateArgs.put("i", value);
            print_do(this._compartment_.stateArgs.get("i").toString());
            if ((int) this._compartment_.stateArgs.get("i") == 0) {
                CountdownCompartment compartment =  new CountdownCompartment(CountdownState.STOP.getValue());
                
                this._transition_(compartment);
                return;
            }
            CountdownCompartment compartment =  new CountdownCompartment(CountdownState.S1.getValue());
            compartment.enterArgs.put("i", (int) this._compartment_.stateArgs.get("i"));
            
            this._transition_(compartment);
            return;
        }
    }
    
    private void _sS1_(FrameEvent e) {
        if(e._message == ">") {
            CountdownCompartment compartment =  new CountdownCompartment(CountdownState.S0.getValue());
            compartment.stateArgs.put("i", (int) e._parameters.get("i"));
            compartment.stateVars.put("dec", 1);
            
            this._transition_(compartment);
            return;
        }
    }
    
    private void _sStop_(FrameEvent e) {
        if(e._message == ">") {
            print_do("done");
            return;
        }
    }
    
    //===================== Actions Block ===================//
    
    protected void print_do(String s) {
        System.out.println(s);
    }
    
    // Unimplemented Actions
    
    
    //=============== Machinery and Mechanisms ==============//
    
    private int _state_;
    
    private void _transition_(CountdownCompartment compartment) {
        this._nextCompartment_ = compartment;
    }
    
    private void _doTransition_(CountdownCompartment nextCompartment) {
        this._mux_(new FrameEvent("<", this._compartment_.exitArgs));
        this._compartment_ = nextCompartment;
        this._mux_(new FrameEvent(">", this._compartment_.enterArgs));
    }
    
}

//=============== Compartment ==============//

class CountdownCompartment {

    int state;
    
    CountdownCompartment(int state) {
        this.state = state;
    }
    
    HashMap<String, Object> stateArgs = new HashMap<String, Object>();
    HashMap<String, Object> stateVars = new HashMap<String, Object>();
    HashMap<String, Object> enterArgs = new HashMap<String, Object>();
    HashMap<String, Object> exitArgs = new HashMap<String, Object>();
    FrameEvent _forwardEvent = new FrameEvent();
}



class CountdownController extends Countdown {
    CountdownController(int i) {
	  super(i);
	}
}


