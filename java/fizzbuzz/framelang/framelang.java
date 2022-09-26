import java.util.*;

interface FrameState { 
    void call(FrameEvent e); 
}
    
class FrameEvent {
    FrameEvent(String message, HashMap<String,Object> parameters) {
        this._message = message;
        this._parameters = parameters;
    }

    FrameEvent() {}
    
    String  _message;
    HashMap<String,Object> _parameters;
    Object  _return;
}