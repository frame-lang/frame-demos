package javaframedemos.PersistentTrafficLight;

import java.util.*;

public class TrafficLightMom {

        // creating private properties
        private int __state;
        private TrafficLightMomCompartment _compartment_;
        private TrafficLightMomCompartment _nextCompartment_;
        private TrafficLight _trafficLight;
        private TrafficLight _data;
        // protected static TrafficLight tfObj;
        FrameEvent e;

        TrafficLightMom tmom;

        public TrafficLightMom() {
                // Create and intialize start state compartment.
                e = new FrameEvent(">", null);
                this.__state = 0;
                this._compartment_ = new TrafficLightMomCompartment(this.__state);
                this._nextCompartment_ = null;

                // Initialize domain
                this._trafficLight = null;
                this._data = null;

                // Send system start event
                try {
                        this._mux_(this.e);
                } catch (Exception e) {
                       //e.printStackTrace();
                }
        }

        // state enum
        private enum TrafficLightState {

                NEW(0),
                SAVING(1),
                PERSISTED(2),
                WORKING(3),
                TRAFFICLIGHTAPI(4),
                END(5);

                private final int value;

                private TrafficLightState(int value) {
                        this.value = value;
                }

                public int getValue() {
                        return value;
                }

        }

        // ===================Interface Block===================//
        public void stop() throws Exception {
                e = new FrameEvent("stop", null);
                this._mux_(e);
        }

        public void tick() throws Exception {
                e = new FrameEvent("tick", null);
                this._mux_(e);
        }

        public void enterRed() throws Exception {
                e = new FrameEvent("enterRed", null);
                this._mux_(e);
        }

        public void enterGreen() throws Exception {
                e = new FrameEvent("enterGreen", null);
                this._mux_(e);
        }

        public void enterYellow() throws Exception {
                e = new FrameEvent("enterYellow", null);
                this._mux_(e);
        }

        public void enterFlashingRed() throws Exception {
                e = new FrameEvent("enterFlashingRed", null);
                this._mux_(e);
        }

        public void exitFlashingRed() throws Exception {
                e = new FrameEvent("exitFlashingRed", null);
                this._mux_(e);
        }

        public void startWorkingTimer() throws Exception {
                e = new FrameEvent("startWorkingTimer", null);
                this._mux_(e);
        }

        public void stopWorkingTimer() throws Exception {
                e = new FrameEvent("stopWorkingTimer", null);
                this._mux_(e);
        }

        public void startFlashingTimer() throws Exception {
                e = new FrameEvent("startFlashingTimer", null);
                this._mux_(e);
        }

        public void stopFlashingTimer() throws Exception {
                e = new FrameEvent("stopFlashingTimer", null);
                this._mux_(e);
        }

        public void changeColor(String color) throws Exception {
                HashMap<String, Object> parameters = new HashMap<String, Object>();
                parameters.put("color", color);

                FrameEvent e = new FrameEvent("changeColor", parameters);
                this._mux_(e);
        }

        public void startFlashing() throws Exception {
                e = new FrameEvent("startFlashing", null);
                this._mux_(e);
        }

        public void stopFlashing() throws Exception {
                e = new FrameEvent("stopFlashing", null);
                this._mux_(e);
        }

        public void changeFlashingAnimation() throws Exception {
                e = new FrameEvent("changeFlashingAnimation", null);
                this._mux_(e);
        }

        public void systemError() throws Exception {
                e = new FrameEvent("systemError", null);
                this._mux_(e);
        }

        public void systemRestart() throws Exception {
                e = new FrameEvent("systemRestart", null);
                this._mux_(e);
        }

        public void log(String msg) throws Exception {
                HashMap<String, Object> parameters = new HashMap<String, Object>();
                parameters.put("msg", msg);

                FrameEvent e = new FrameEvent("log", parameters);
                this._mux_(e);
        }

        // ===================Multiplexer===================//
        public void _mux_(FrameEvent e) throws Exception {
                if (this._compartment_ != null) {
                        if (this._compartment_.state == TrafficLightState.NEW.getValue()) {
                                this._sNew_(e);
                        } else if (this._compartment_.state == TrafficLightState.SAVING.getValue()) {
                                this._sSaving_(e);
                        } else if (this._compartment_.state == TrafficLightState.PERSISTED.getValue()) {
                                this._sPersisted_(e);
                        } else if (this._compartment_.state == TrafficLightState.WORKING.getValue()) {
                                this._sWorking_(e);
                        } else if (this._compartment_.state == TrafficLightState.TRAFFICLIGHTAPI.getValue()) {
                                this._sTrafficLightApi_(e);
                        } else if (this._compartment_.state == TrafficLightState.END.getValue()) {
                                this._sEnd_(e);
                        }
                }
                if (this._nextCompartment_ != null) {
                        TrafficLightMomCompartment nextCompartment = this._nextCompartment_;
                        this._nextCompartment_ = null;
                        if (nextCompartment._forwardEvent != null && nextCompartment._forwardEvent._message != null &&
                                        nextCompartment._forwardEvent._message.equalsIgnoreCase(">")) {
                                this._mux_(new FrameEvent("<", this._compartment_.ExitArgs));
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

        // =================== Machine Block ===================//
        private void _sNew_(FrameEvent e) {
                if (e._message.equalsIgnoreCase(">")) {
                        this._trafficLight = new TrafficLightController(this);
                        // Traffic Light\nStarted
                        TrafficLightMomCompartment compartment = new TrafficLightMomCompartment(

                        TrafficLightState.SAVING.getValue());
                        this._transition_(compartment);
                        return;
                }

                this._sTrafficLightApi_(e);
                return;
        }

        private void _sSaving_(FrameEvent e) throws Exception {
                if (e._message.equals(">")) {

                        this._data = this._trafficLight.marshal();
                        this._trafficLight = null;
                        // Saved
                        TrafficLightMomCompartment compartment = new TrafficLightMomCompartment(
                                        TrafficLightState.PERSISTED.getValue());

                        this._transition_(compartment);
                        return;

                }
        }

        private void _sPersisted_(FrameEvent e) {
                if (e._message != null) {
                        if (e._message.equalsIgnoreCase("tick")) {

                                // Tick
                                TrafficLightMomCompartment compartment = new TrafficLightMomCompartment(
                                                TrafficLightState.WORKING.getValue());

                                compartment._forwardEvent = e;

                                this._transition_(compartment);
                                return;
                        }

                        else if (e._message.equalsIgnoreCase("systemError")) {
                                // System Error
                                TrafficLightMomCompartment compartment = new TrafficLightMomCompartment(
                                                TrafficLightState.WORKING.getValue());

                                compartment._forwardEvent = e;

                                this._transition_(compartment);
                                return;
                        }

                        else if (e._message.equalsIgnoreCase("stop")) {
                                // Stop
                                TrafficLightMomCompartment compartment = new TrafficLightMomCompartment(
                                                TrafficLightState.END.getValue());

                                this._transition_(compartment);
                                return;
                        }

                }
        }

        private void _sWorking_(FrameEvent e) throws Exception {
                if (e._message.equalsIgnoreCase(">")) {

                        this._trafficLight = TrafficLightController.loadTrafficLight((TrafficLightMomController) this,
                                        (TrafficLightController) this._data);
                        return;
                }

                else if (e._message.equalsIgnoreCase("tick")) {
                        this._trafficLight.tick();
                        // Done
                        TrafficLightMomCompartment compartment = new TrafficLightMomCompartment(
                                        TrafficLightState.SAVING.getValue());

                        this._transition_(compartment);
                        return;
                }

                else if (e._message.equalsIgnoreCase("systemError")) {
                        this._trafficLight.systemError();
                        // Done
                        TrafficLightMomCompartment compartment = new TrafficLightMomCompartment(
                                        TrafficLightState.SAVING.getValue());

                        this._transition_(compartment);
                        return;
                }

                this._sTrafficLightApi_(e);

        }

        private void _sTrafficLightApi_(FrameEvent e) {
                if (e._message.equalsIgnoreCase("enterRed")) {
                        this.enterRed_do();
                        return;
                }

                else if (e._message.equalsIgnoreCase("enterGreen")) {
                        this.enterGreen_do();
                        return;
                }

                else if (e._message.equalsIgnoreCase("enterYellow")) {
                        this.enterYellow_do();
                        return;
                }

                else if (e._message.equalsIgnoreCase("enterFlashingRed")) {
                        this.enterFlashingRed_do();
                        return;
                }

                else if (e._message.equalsIgnoreCase("exitFlashingRed")) {
                        this.exitFlashingRed_do();
                        return;
                }

                else if (e._message.equalsIgnoreCase("startWorkingTimer")) {
                        this.startWorkingTimer_do();
                        return;
                }

                else if (e._message.equalsIgnoreCase("stopWorkingTimer")) {
                        this.stopWorkingTimer_do();
                        return;
                }

                else if (e._message.equalsIgnoreCase("startFlashingTimer")) {
                        this.startFlashingTimer_do();
                        return;
                }

                else if (e._message.equalsIgnoreCase("stopFlashingTimer")) {
                        this.stopFlashingTimer_do();
                        return;
                }

                else if (e._message.equalsIgnoreCase("changeColor")) {
                        this.changeColor_do(String.valueOf(((String) e._parameters.get("color"))));
                        return;
                }

                else if (e._message.equalsIgnoreCase("startFlashing")) {
                        this.startFlashing_do();
                        return;
                }

                else if (e._message.equalsIgnoreCase("stopFlashing")) {
                        this.stopFlashing_do();
                        return;
                }

                else if (e._message.equalsIgnoreCase("changeFlashingAnimation")) {
                        this.changeFlashingAnimation_do();
                        return;
                }

                else if (e._message.equalsIgnoreCase("systemError")) {
                        this.systemError_do();
                        return;
                }

                else if (e._message.equalsIgnoreCase("systemRestart")) {
                        this.systemRestart_do();
                        return;
                }

                else if (e._message.equalsIgnoreCase("log")) {
                        this.log_do(String.valueOf(((String) e._parameters.get("msg"))));
                        return;
                }

        }

        private void _sEnd_(FrameEvent e) throws Exception {
                if (e._message.equalsIgnoreCase(">")) {

                        this._trafficLight = TrafficLightController.loadTrafficLight((TrafficLightMomController) this,
                                        (TrafficLightController) this._data);
                        this._trafficLight.stop();
                        this._trafficLight = null;
                        return;

                }
                this._sTrafficLightApi_(e);
        }
        // =================== Action Block ==================//
        // Unimplemented Actions

        protected void enterRed_do() {
                throw new UnsupportedOperationException();
        }

        protected void enterGreen_do() {
                throw new UnsupportedOperationException();
        }

        protected void enterYellow_do() {
                throw new UnsupportedOperationException();
        }

        protected void enterFlashingRed_do() {
                throw new UnsupportedOperationException();
        }

        protected void exitFlashingRed_do() {
                throw new UnsupportedOperationException();
        }

        protected void startWorkingTimer_do() {
                throw new UnsupportedOperationException();
        }

        protected void stopWorkingTimer_do() {
                throw new UnsupportedOperationException();
        }

        protected void startFlashingTimer_do() {
                throw new UnsupportedOperationException();
        }

        protected void stopFlashingTimer_do() {
                throw new UnsupportedOperationException();
        }

        protected void changeColor_do(String color) {
                throw new UnsupportedOperationException();
        }

        protected void startFlashing_do() {
                throw new UnsupportedOperationException();
        }

        protected void stopFlashing_do() {
                throw new UnsupportedOperationException();
        }

        protected void changeFlashingAnimation_do() {
                throw new UnsupportedOperationException();
        }

        protected void systemError_do() {
                throw new UnsupportedOperationException();
        }

        protected void systemRestart_do() {
                throw new UnsupportedOperationException();
        }

        protected void log_do(String msg) {
                throw new UnsupportedOperationException();
        }

        // =============== Machinery and Mechanisms ==============//
        private int _transition_(TrafficLightMomCompartment compartment) {
                this._nextCompartment_ = compartment;
                return this._nextCompartment_.state;
        }

        private void _doTransition_(TrafficLightMomCompartment nextCompartment) throws Exception {
                this._mux_(new FrameEvent("<", this._compartment_.ExitArgs));
                this._compartment_ = nextCompartment;
                this._mux_(new FrameEvent(">", this._compartment_.EnterArgs));
        }

}

// =============== Compartment ==============//
class TrafficLightMomCompartment {
        int state;

        TrafficLightMomCompartment(int state) {
                this.state = state;
        }

        HashMap<String, Object> StateArgs = new HashMap<String, Object>();
        HashMap<String, Object> StateVars = new HashMap<String, Object>();
        HashMap<String, Object> EnterArgs = new HashMap<String, Object>();
        HashMap<String, Object> ExitArgs = new HashMap<String, Object>();
        FrameEvent _forwardEvent = new FrameEvent();
}

class TrafficLightMomController extends TrafficLightMom {
        TrafficLightMomController() {
                super();
        }

        @Override
        protected void enterRed_do() {
                System.out.println("Red");
        }

        @Override
        protected void enterGreen_do() {
                System.out.println("Green");
        }

        @Override
        protected void enterYellow_do() {
                System.out.println("Yellow");
        }

        protected void enterFlashingRed_do() {
        }

        protected void exitFlashingRed_do() {
        }

        protected void startWorkingTimer_do() {
        }

        protected void stopWorkingTimer_do() {
        }

        protected void startFlashingTimer_do() {
        }

        protected void stopFlashingTimer_do() {
        }

        protected void changeColor_do(String color) {
        }

        protected void startFlashing_do() {
        }

        protected void stopFlashing_do() {
        }

        protected void changeFlashingAnimation_do() {
        }

        protected void log_do(String msg) {
        }

}
