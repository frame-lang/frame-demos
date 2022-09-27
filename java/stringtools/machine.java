import java.util.*;

//emitted from framec_v0.10.0
//get include files at https://github.com/frame-lang/frame-ancillary-files

enum StringToolsState {
	ROUTER(0), REVERSE(1), MAKEPALINDROME(2);

	public final int value;

	private StringToolsState(int value) {
		this.value = value;
	}

	public int getValue() {
		return value;
	}
}

class StringTools {
	private StringToolsCompartment _compartment_;
	private StringToolsCompartment _nextCompartment_;

	StringTools() {
		//Create and initialize start state compartment.
		
		this._compartment_ = new StringToolsCompartment(this._state_);
		this._nextCompartment_ = null;
		
		//Initialize domain
		
		//Send system start event
		FrameEvent e = new FrameEvent(">", null);
		this._mux_(e);
	}

	// ===================== Interface Block ===================//

	public String reverse(String str) {
		HashMap<String, Object> map = new HashMap<>();
		map.put("str", str);
		FrameEvent e = new FrameEvent("reverse", map);
		this._mux_(e);
		return e._return.toString();
	}

	public String makePalindrome(String str) {
		HashMap<String, Object> map = new HashMap<>();
		map.put("str", str);
		FrameEvent e = new FrameEvent("makePalindrome", map);
		this._mux_(e);
		return e._return.toString();
	}
	
    //====================== Multiplexer ====================//

	private void _mux_(FrameEvent e) {
		if(this._compartment_.state == StringToolsState.ROUTER.getValue()) {
			this._sRouter_(e);
		} else if(this._compartment_.state == StringToolsState.REVERSE.getValue()) {
			this._sReverse_(e);
		} else if(this._compartment_.state == StringToolsState.MAKEPALINDROME.getValue()) {
			this._sMakePalindrome_(e);
		}
		
		if (this._nextCompartment_ != null) {
			StringToolsCompartment nextCompartment = this._nextCompartment_;
			this._nextCompartment_ = null;
			if (nextCompartment._forwardEvent != null && nextCompartment._forwardEvent._message == ">") {
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
	

	// ===================== Machine Block ===================// // see working demo
	// at codepen.io:
	// https://codepen.io/frame-lang/pen/e96568889939d3bc112b16528900c352
	// try out on http://framepiler.frame-lang.org

	private void _sRouter_(FrameEvent e) {
		if (e._message == "makePalindrome") {
            // make\npalindrome
            StringToolsCompartment compartment = new StringToolsCompartment(StringToolsState.MAKEPALINDROME.getValue()); 
            compartment._forwardEvent = e;
            this._transition_(compartment);
            return;
        }
        else if (e._message == "reverse") {
            // reverse
        	StringToolsCompartment compartment = new StringToolsCompartment(StringToolsState.REVERSE.getValue());
        	compartment._forwardEvent = e;
        	this._transition_(compartment);
        	return;
        }
	}

	private void _sReverse_(FrameEvent e) {
		if(e._message == "reverse") {
			e._return = this.reverse_str_do((String) e._parameters.get("str"));
			StringToolsCompartment compartment = new StringToolsCompartment(StringToolsState.ROUTER.getValue());
			this._transition_(compartment);
			return;
		}
	}

	private void _sMakePalindrome_(FrameEvent e) {
		if(e._message == "makePalindrome") {
			e._return = (e._parameters.get("str")) + this.reverse_str_do((String) (e._parameters.get("str")));
			StringToolsCompartment compartment = new StringToolsCompartment(StringToolsState.ROUTER.getValue());
			this._transition_(compartment);
			return;
		}
	}

	// ===================== Actions Block ===================//

	protected String reverse_str_do(String str) {
		throw new UnsupportedOperationException();
	}

	// =============== Machinery and Mechanisms ==============//

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

//====================Compartment====================//

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
	protected String reverse_str_do(String str) {
		String result = "";
		for(char c: str.toCharArray()) {
			result = c+result;
		}
		return result;
	}
}
