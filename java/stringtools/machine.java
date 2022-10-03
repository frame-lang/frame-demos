import java.util.*;

//emitted from framec_v0.10.0
//get include files at https://github.com/frame-lang/frame-ancillary-files

class StringTools {
    
    private StringToolsCompartment _compartment_;
    private StringToolsCompartment _nextCompartment_;
    
    StringTools() {
        
        // Create and intialize start state compartment.
        this._compartment_ = new StringToolsCompartment(this._state_);
        this._nextCompartment_ = null;
        
        
        // Send system start event
        FrameEvent frame_event = new FrameEvent(">", null);
        this._mux_(frame_event);
    
    }
    
    // states enum
    private enum StringToolsState {
        
        ROUTER(0),
        REVERSE(1),
        MAKEPALINDROME(2);
        
        public final int value;
        
        private StringToolsState(int value) {
            this.value=value;
        }
        
        public int getValue() {
            return value;
        }
    }
    
    //====================== Multiplexer ====================//
    
    private void _mux_(FrameEvent e) {
        if(this._compartment_.state == StringToolsState.ROUTER.getValue()) {
            this._sRouter_(e);
        }else if(this._compartment_.state == StringToolsState.REVERSE.getValue()) {
            this._sReverse_(e);
        }else if(this._compartment_.state == StringToolsState.MAKEPALINDROME.getValue()) {
            this._sMakePalindrome_(e);
        }
        
        
        if(this._nextCompartment_ != null) {
            StringToolsCompartment nextCompartment = this._nextCompartment_;
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
    
    public String reverse(String str) {
        HashMap<String,Object> parameters = new HashMap<String,Object>();
        parameters.put("str", str);

        FrameEvent e = new FrameEvent("reverse", parameters);
        this._mux_(e);
        return (String) e._return;
    }
    
    public String makePalindrome(String str) {
        HashMap<String,Object> parameters = new HashMap<String,Object>();
        parameters.put("str", str);

        FrameEvent e = new FrameEvent("makePalindrome", parameters);
        this._mux_(e);
        return (String) e._return;
    }
    
    
    //===================== Machine Block ===================//
    
    private void _sRouter_(FrameEvent e) {
        if(e._message == "makePalindrome") {
            // make\npalindrome
            StringToolsCompartment compartment =  new StringToolsCompartment(StringToolsState.MAKEPALINDROME.getValue());
            
            compartment._forwardEvent = e;
            
            this._transition_(compartment);
            return;
        }
        else if(e._message == "reverse") {
            // reverse
            StringToolsCompartment compartment =  new StringToolsCompartment(StringToolsState.REVERSE.getValue());
            
            compartment._forwardEvent = e;
            
            this._transition_(compartment);
            return;
        }
    }
    
    private void _sReverse_(FrameEvent e) {
        if(e._message == "reverse") {
            e._return = reverse_str_do(String.valueOf(((String) e._parameters.get("str"))));
            // ready
            StringToolsCompartment compartment =  new StringToolsCompartment(StringToolsState.ROUTER.getValue());
            
            this._transition_(compartment);
            return;
        }
    }
    
    private void _sMakePalindrome_(FrameEvent e) {
        if(e._message == "makePalindrome") {
            e._return = String.valueOf(((String) e._parameters.get("str"))) + reverse_str_do(String.valueOf(((String) e._parameters.get("str"))));
            // ready
            StringToolsCompartment compartment =  new StringToolsCompartment(StringToolsState.ROUTER.getValue());
            
            this._transition_(compartment);
            return;
        }
    }
    
    //===================== Actions Block ===================//
    
    protected String reverse_str_do(String str) {
        
	return str;

    }
    
    // Unimplemented Actions
    
    protected void ccvcv_do(String str) { throw new UnsupportedOperationException(); }
    
    //=============== Machinery and Mechanisms ==============//
    
    private int _state_;
    
    private void _transition_(StringToolsCompartment compartment) {
        this._nextCompartment_ = compartment;
    }
    
    private void _doTransition_(StringToolsCompartment nextCompartment) {
        this._mux_(new FrameEvent("<", this._compartment_.exitArgs));
        this._compartment_ = nextCompartment;
        this._mux_(new FrameEvent(">", this._compartment_.enterArgs));
    }
    
}

//=============== Compartment ==============//

class StringToolsCompartment {

    int state;
    
    StringToolsCompartment(int state) {
        this.state = state;
    }
    
    HashMap<String, Object> stateArgs = new HashMap<String, Object>();
    HashMap<String, Object> stateVars = new HashMap<String, Object>();
    HashMap<String, Object> enterArgs = new HashMap<String, Object>();
    HashMap<String, Object> exitArgs = new HashMap<String, Object>();
    FrameEvent _forwardEvent = new FrameEvent();
}

class StringToolsController extends StringTools {
	StringToolsController() {
	  super();
	}
    
    protected String reverse_str_do(String str) {
		String result = "";
		for(char c: str.toCharArray()) {
			result = c+result;
		}
		return result;
	}
}

