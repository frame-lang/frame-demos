import java.util.*;

class HistoryBasic {
    
    private HistoryBasicCompartment _compartment_;
    private HistoryBasicCompartment _nextCompartment_;
    
    HistoryBasic() {
        
        // Create state stack.
        
        this._stateStack_ = new Stack<HistoryBasicCompartment>();
        
        // Create and intialize start state compartment.
        this._compartment_ = new HistoryBasicCompartment(this._state_);
        this._nextCompartment_ = null;
        
        
        // Send system start event
        FrameEvent frame_event = new FrameEvent(">", null);
        this._mux_(frame_event);
    
    }
    
    // states enum
    private enum HistoryBasicState {
        
        START(0),
        S0(1),
        S1(2),
        DEADEND(3);
        
        public final int value;
        
        private HistoryBasicState(int value) {
            this.value=value;
        }
        
        public int getValue() {
            return value;
        }
    }
    
    //====================== Multiplexer ====================//
    
    private void _mux_(FrameEvent e) {
        if(this._compartment_.state == HistoryBasicState.START.getValue()) {
            this._sStart_(e);
        }else if(this._compartment_.state == HistoryBasicState.S0.getValue()) {
            this._sS0_(e);
        }else if(this._compartment_.state == HistoryBasicState.S1.getValue()) {
            this._sS1_(e);
        }else if(this._compartment_.state == HistoryBasicState.DEADEND.getValue()) {
            this._sDeadEnd_(e);
        }
        
        
        if(this._nextCompartment_ != null) {
            HistoryBasicCompartment nextCompartment = this._nextCompartment_;
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
    
    //===================== Interface Block ===================//
    
    public void start() {
        FrameEvent e = new FrameEvent(">>", null);
        this._mux_(e);
    }
    
    public void switchState() {
        FrameEvent e = new FrameEvent("switchState", null);
        this._mux_(e);
    }
    
    public void gotoDeadEnd() {
        FrameEvent e = new FrameEvent("gotoDeadEnd", null);
        this._mux_(e);
    }
    
    public void back() {
        FrameEvent e = new FrameEvent("back", null);
        this._mux_(e);
    }
    
    
    //===================== Machine Block ===================//
    
    private void _sStart_(FrameEvent e) {
        if(e._message == ">>") {
            HistoryBasicCompartment compartment =  new HistoryBasicCompartment(HistoryBasicState.S0.getValue());
            
            this._transition_(compartment);
            return;
        }
    }
    
    private void _sS0_(FrameEvent e) {
        if(e._message == ">") {
            print_do("Enter S0");
            return;
        }
        else if(e._message == "switchState") {
            // Switch\nState
            HistoryBasicCompartment compartment =  new HistoryBasicCompartment(HistoryBasicState.S1.getValue());
            
            this._transition_(compartment);
            return;
        }
        else if(e._message == "gotoDeadEnd") {
            _stateStack_push_(this._compartment_);
            // Goto\nDead End
            HistoryBasicCompartment compartment =  new HistoryBasicCompartment(HistoryBasicState.DEADEND.getValue());
            
            this._transition_(compartment);
            return;
        }
    }
    
    private void _sS1_(FrameEvent e) {
        if(e._message == ">") {
            print_do("Enter S1");
            return;
        }
        else if(e._message == "switchState") {
            // Switch\nState
            HistoryBasicCompartment compartment =  new HistoryBasicCompartment(HistoryBasicState.S0.getValue());
            
            this._transition_(compartment);
            return;
        }
        else if(e._message == "gotoDeadEnd") {
            _stateStack_push_(this._compartment_);
            // Goto\nDead End
            HistoryBasicCompartment compartment =  new HistoryBasicCompartment(HistoryBasicState.DEADEND.getValue());
            
            this._transition_(compartment);
            return;
        }
    }
    
    private void _sDeadEnd_(FrameEvent e) {
        if(e._message == ">") {
            print_do("Enter $DeadEnd");
            return;
        }
        else if(e._message == "back") {
            // Go Back
            HistoryBasicCompartment compartment = this._stateStack_pop_();
            this._transition_(compartment);
            return;
        }
    }
    
    //===================== Actions Block ===================//
    
    // Unimplemented Actions
    
    protected void print_do(String msg) { throw new UnsupportedOperationException(); }
    
    //=============== Machinery and Mechanisms ==============//
    
    private int _state_;
    
    private void _transition_(HistoryBasicCompartment compartment) {
        this._nextCompartment_ = compartment;
    }
    
    private void _doTransition_(HistoryBasicCompartment nextCompartment) {
        this._mux_(new FrameEvent("<", this._compartment_.exitArgs));
        this._compartment_ = nextCompartment;
        this._mux_(new FrameEvent(">", this._compartment_.enterArgs));
    }
    
    private Stack<HistoryBasicCompartment> _stateStack_ = new Stack<>();
    
    private void _stateStack_push_(HistoryBasicCompartment compartment) {
        _stateStack_.push(compartment);
    }
    
    private HistoryBasicCompartment _stateStack_pop_() {
        return _stateStack_.pop();
    }
    
}

//=============== Compartment ==============//

class HistoryBasicCompartment {

    int state;
    
    HistoryBasicCompartment(int state) {
        this.state = state;
    }
    
    HashMap<String, Object> stateArgs = new HashMap<String, Object>();
    HashMap<String, Object> stateVars = new HashMap<String, Object>();
    HashMap<String, Object> enterArgs = new HashMap<String, Object>();
    HashMap<String, Object> exitArgs = new HashMap<String, Object>();
    FrameEvent _forwardEvent = new FrameEvent();
}


class HistoryBasicController extends HistoryBasic {
	HistoryBasicController() {
	  super();
	}

    protected void print_do(String msg) {}
}
