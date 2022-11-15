package javaframedemos;


import java.util.*;

public class FrameEvent {
    FrameEvent(String message, Map<String, Object> parameters) {
        this._message = message;
        this._parameters = parameters;
        this._return = null;
        return;
    }

    public FrameEvent() {
    }

    String _message;
    Map<String, Object> _parameters;
    Object _return;


}
