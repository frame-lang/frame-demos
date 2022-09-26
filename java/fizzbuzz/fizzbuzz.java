import java.util.*;

// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files

enum FizzBuzzState {
	BEGIN(0), FIZZ(1), BUZZ(2), DIGIT(3), END(4);

	public final int value;

	private FizzBuzzState(int value) {
		this.value = value;
	}

	public int getValue() {
		return value;
	}
}

class FizzBuzz {
	private FizzBuzzCompartment _compartment_;
	private FizzBuzzCompartment _nextCompartment_;

	FizzBuzz() {
		this._compartment_ = new FizzBuzzCompartment(this._state_);
		this._nextCompartment_ = null;
		
		FrameEvent e = new FrameEvent(">", null);
		this._mux_(e);
	}

	// ===================== Interface Block ===================//

	public void start() {
		FrameEvent e = new FrameEvent(">>", null);
		this._mux_(e);
	}

	// ====================== Multiplexer ====================//

	private void _mux_(FrameEvent e) {

		if (this._compartment_.state == FizzBuzzState.BEGIN.getValue()) {
			this._sBegin_(e);
		} else if(this._compartment_.state == FizzBuzzState.FIZZ.getValue()) {
			this._sFizz_(e);
		} else if(this._compartment_.state == FizzBuzzState.BUZZ.getValue()) {
			this._sBuzz_(e);
		} else if(this._compartment_.state == FizzBuzzState.DIGIT.getValue()) {
			this._sDigit_(e);
		} else if(this._compartment_.state == FizzBuzzState.END.getValue()) {
			this._sEnd_(e);
		}

		if (this._nextCompartment_ != null) {
			FizzBuzzCompartment nextCompartment = this._nextCompartment_;
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

	private void _doTransition_(FizzBuzzCompartment nextCompartment) {
		this._mux_(new FrameEvent("<", this._compartment_.exitArgs));
		this._compartment_ = nextCompartment;
		this._mux_(new FrameEvent(">", this._compartment_.enterArgs));
	}

	// ===================== Machine Block ===================// // see working demo
	// at codepen.io:
	// https://codepen.io/frame-lang/pen/e96568889939d3bc112b16528900c352
	// try out on http://framepiler.frame-lang.org

	private void _sBegin_(FrameEvent e) {
		if (e._message == ">>") {
			FizzBuzzCompartment compartment = new FizzBuzzCompartment(FizzBuzzState.FIZZ.getValue());
			compartment.enterArgs.put("i", 1);
			this._transition_(compartment);
			return;
		}
	}

	private void _sFizz_(FrameEvent e) {
		if (e._message == ">") {
			if (gt_100_do(((int) e._parameters.get("i")))) {
				// i > 100
				FizzBuzzCompartment compartment = new FizzBuzzCompartment(FizzBuzzState.END.getValue());
				this._transition_(compartment);
				return;
			}
			if (mod3_eq0_do(((int) e._parameters.get("i")))) {
				print_do("Fizz");
				// i % 3 == 0
				FizzBuzzCompartment compartment = new FizzBuzzCompartment(FizzBuzzState.BUZZ.getValue());
				compartment.enterArgs.put("i", e._parameters.get("i"));
                compartment.enterArgs.put("fizzed", true);
				this._transition_(compartment);
				return;
			} else {
				// i % 3 != 0
				FizzBuzzCompartment compartment = new FizzBuzzCompartment(FizzBuzzState.BUZZ.getValue());
				compartment.enterArgs.put("i", e._parameters.get("i"));
                compartment.enterArgs.put("fizzed", false);
				this._transition_(compartment);
			}
			return;
		}
	}

	private void _sBuzz_(FrameEvent e) {
		if (e._message == ">") {
			if (mod5_eq0_do(((int) e._parameters.get("i")))) {
				print_do("Buzz");
				// i % 5 == 0
				this._compartment_.exitArgs.put("output", " ");
				FizzBuzzCompartment compartment = new FizzBuzzCompartment(FizzBuzzState.FIZZ.getValue());
				compartment.enterArgs.put("i", this.plus_1_do((int) e._parameters.get("i")));
				this._transition_(compartment);
				return;
			}
			if (((boolean) e._parameters.get("fizzed"))) {
				// fizzed
				this._compartment_.exitArgs.put("output", " ");
				FizzBuzzCompartment compartment = new FizzBuzzCompartment(FizzBuzzState.FIZZ.getValue());
				compartment.enterArgs.put("i", this.plus_1_do((int) e._parameters.get("i")));
				this._transition_(compartment);
				return;
			}
			// ! mod3 | mod5
			this._compartment_.exitArgs.put("output", " ");
			FizzBuzzCompartment compartment = new FizzBuzzCompartment(FizzBuzzState.DIGIT.getValue());
			compartment.enterArgs.put("i", this.plus_1_do((int) e._parameters.get("i")));
			this._transition_(compartment);
			return;
		} else if (e._message == "<") {
			print_do(((String) e._parameters.get("output")));
			return;
		}
	}

	private void _sDigit_(FrameEvent e) {
		if (e._message == ">") {
			print_do((String.valueOf((int) e._parameters.get("i"))));
			print_do(" ");
			// loop
			FizzBuzzCompartment compartment = new FizzBuzzCompartment(FizzBuzzState.FIZZ.getValue());
			compartment.enterArgs.put("i", this.plus_1_do((int) e._parameters.get("i")));
			this._transition_(compartment);
			return;
		}
	}

	private void _sEnd_(FrameEvent e) {
	}

	// ===================== Actions Block ===================//

	protected void print_do(String msg) {
		throw new UnsupportedOperationException();
	}

	protected boolean gt_100_do(int i) {
		throw new UnsupportedOperationException();
	}

	protected boolean mod3_eq0_do(int i) {
		throw new UnsupportedOperationException();
	}

	protected boolean mod5_eq0_do(int i) {
		throw new UnsupportedOperationException();
	}

	protected int plus_1_do(int i) {
		throw new UnsupportedOperationException();
	}

	// =============== Machinery and Mechanisms ==============//

	private int _state_;

	private void _transition_(FizzBuzzCompartment compartment) {
		this._nextCompartment_ = compartment;
	}

}

//====================Compartment====================//

class FizzBuzzCompartment {

	int state;

	FizzBuzzCompartment(int state) {
		this.state = state;
	}

	HashMap<String, Object> stateArgs = new HashMap<String, Object>();
	HashMap<String, Object> stateVars = new HashMap<String, Object>();
	HashMap<String, Object> enterArgs = new HashMap<String, Object>();
	HashMap<String, Object> exitArgs = new HashMap<String, Object>();

	FrameEvent _forwardEvent = new FrameEvent();
}

class FizzBuzzController extends FizzBuzz {

	FizzBuzzController() {
		super();
	}

	protected void print_do(String msg) {
		System.out.print(msg + " ");
	}

	protected boolean gt_100_do(int i) {
		return i > 100;
	}

	protected boolean mod3_eq0_do(int i) {
		return i % 3 == 0;
	}

	protected boolean mod5_eq0_do(int i) {
		return i % 5 == 0;
	}

	protected int plus_1_do(int i) {
		return i + 1;
	}
}