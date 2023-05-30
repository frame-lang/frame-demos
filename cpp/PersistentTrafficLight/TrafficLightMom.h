// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files

#ifndef TRAFFICLIGHTMOM_H
#define TRAFFICLIGHTMOM_H

#pragma once
#include <iostream>
#include <unordered_map>
#include <string>
#include <any>
#include "TrafficLight.h"

class TrafficLight;
class TrafficLightController;
class TrafficLightMomController;
class TrafficLightMom;


//=============== Compartment ==============//

class TrafficLightMomCompartment
{
public:
    int state;

    TrafficLightMomCompartment(int state)
    {
        this->state = state;
    }

    std::unordered_map<std::string, std::any> stateArgs;
    std::unordered_map<std::string, std::any> stateVars;
    std::unordered_map<std::string, std::any> enterArgs;
    std::unordered_map<std::string, std::any> exitArgs;
    FrameEvent *_forwardEvent = nullptr;
};

class TrafficLightMom
{
private:
    TrafficLightMomCompartment *_compartment_;
    TrafficLightMomCompartment *_nextCompartment_;

public:
    TrafficLightMom()
    {
        // Create and intialize start state compartment.
        _state_ = static_cast<int>(TrafficLightMomState::NEW);

        _compartment_ = new TrafficLightMomCompartment(this->_state_);
        _nextCompartment_ = nullptr;

        // Send system start event
        FrameEvent frame_event(">", std::unordered_map<std::string, std::any>());
        _mux_(frame_event);
    }

    // states enum
private:
    enum class TrafficLightMomState
    {
        NEW = 0,
        SAVING = 1,
        PERSISTED = 2,
        WORKING = 3,
        TRAFFICLIGHTAPI = 4,
        END = 5
    };

    //====================== Multiplexer ====================//

private:
    void _mux_(FrameEvent e)
    {
        if (this->_compartment_->state == static_cast<int>(TrafficLightMomState::NEW))
        {
            this->_sNew_(e);
        }
        else if (this->_compartment_->state == static_cast<int>(TrafficLightMomState::SAVING))
        {
            this->_sSaving_(e);
        }
        else if (this->_compartment_->state == static_cast<int>(TrafficLightMomState::PERSISTED))
        {
            this->_sPersisted_(e);
        }
        else if (this->_compartment_->state == static_cast<int>(TrafficLightMomState::WORKING))
        {
            this->_sWorking_(e);
        }
        else if (this->_compartment_->state == static_cast<int>(TrafficLightMomState::TRAFFICLIGHTAPI))
        {
            this->_sTrafficLightApi_(e);
        }
        else if (this->_compartment_->state == static_cast<int>(TrafficLightMomState::END))
        {
            this->_sEnd_(e);
        }

        if (this->_nextCompartment_ != nullptr)
        {
            TrafficLightMomCompartment *nextCompartment = this->_nextCompartment_;
            this->_nextCompartment_ = nullptr;
            if (nextCompartment->_forwardEvent != nullptr &&
                nextCompartment->_forwardEvent->_message == ">")
            {
                this->_mux_(FrameEvent("<", this->_compartment_->exitArgs));
                this->_compartment_ = nextCompartment;
                this->_mux_(*nextCompartment->_forwardEvent);
            }
            else
            {
                this->_doTransition_(nextCompartment);
                if (nextCompartment->_forwardEvent != nullptr)
                {
                    this->_mux_(*nextCompartment->_forwardEvent);
                }
            }
            nextCompartment->_forwardEvent = nullptr;
        }
    }

    //===================== Interface Block ===================//
public:
    void stop()
    {
        FrameEvent e("stop", std::unordered_map<std::string, std::any>());
        this->_mux_(e);
    }

    void tick()
    {
        FrameEvent e("tick", std::unordered_map<std::string, std::any>());
        this->_mux_(e);
    }

    void enterRed()
    {
        FrameEvent e("enterRed", std::unordered_map<std::string, std::any>());
        this->_mux_(e);
    }

    void enterGreen()
    {
        FrameEvent e("enterGreen", std::unordered_map<std::string, std::any>());
        this->_mux_(e);
    }

    void enterYellow()
    {
        FrameEvent e("enterYellow", std::unordered_map<std::string, std::any>());
        this->_mux_(e);
    }

    void enterFlashingRed()
    {
        FrameEvent e("enterFlashingRed", std::unordered_map<std::string, std::any>());
        this->_mux_(e);
    }

    void exitFlashingRed()
    {
        FrameEvent e("exitFlashingRed", std::unordered_map<std::string, std::any>());
        this->_mux_(e);
    }

    void startWorkingTimer()
    {
        FrameEvent e("startWorkingTimer", std::unordered_map<std::string, std::any>());
        this->_mux_(e);
    }

    void stopWorkingTimer()
    {
        FrameEvent e("stopWorkingTimer", std::unordered_map<std::string, std::any>());
        this->_mux_(e);
    }

    void startFlashingTimer()
    {
        FrameEvent e("startFlashingTimer", std::unordered_map<std::string, std::any>());
        this->_mux_(e);
    }

    void stopFlashingTimer()
    {
        FrameEvent e("stopFlashingTimer", std::unordered_map<std::string, std::any>());
        this->_mux_(e);
    }

    void changeColor(std::string color)
    {
        std::unordered_map<std::string, std::any> parameters;
        parameters["color"] = color;

        FrameEvent e("changeColor", parameters);
        this->_mux_(e);
    }

    void startFlashing()
    {
        FrameEvent e("startFlashing", std::unordered_map<std::string, std::any>());
        this->_mux_(e);
    }

    void stopFlashing()
    {
        FrameEvent e("stopFlashing", std::unordered_map<std::string, std::any>());
        this->_mux_(e);
    }

    void changeFlashingAnimation()
    {
        FrameEvent e("changeFlashingAnimation", std::unordered_map<std::string, std::any>());
        this->_mux_(e);
    }

    void systemError()
    {
        FrameEvent e("systemError", std::unordered_map<std::string, std::any>());
        this->_mux_(e);
    }

    void systemRestart()
    {
        FrameEvent e("systemRestart", std::unordered_map<std::string, std::any>());
        this->_mux_(e);
    }

    void log(std::string msg)
    {
        std::unordered_map<std::string, std::any> parameters;
        parameters["msg"] = msg;

        FrameEvent e("log", parameters);
        this->_mux_(e);
    }

    //===================== Machine Block ===================//

private:
    void _sNew_(FrameEvent e)
    {
        if (e._message == ">")
        {
            this->_trafficLight = new TrafficLightController(this);
            // Traffic Light\nStarted
            TrafficLightMomCompartment *compartment = new TrafficLightMomCompartment(static_cast<int>(TrafficLightMomState::SAVING));

            this->_transition_(compartment);
            return;
        }
        _sTrafficLightApi_(e);
    }

    void _sSaving_(FrameEvent e)
    {
        if (e._message == ">")
        {
            this->_data = this->_trafficLight->marshal();
            this->_trafficLight = nullptr;
            // Saved
            TrafficLightMomCompartment *compartment = new TrafficLightMomCompartment(static_cast<int>(TrafficLightMomState::PERSISTED));

            this->_transition_(compartment);
            return;
        }
    }

    void _sPersisted_(FrameEvent e)
    {
        if (e._message == "tick")
        {
            // Tick
            TrafficLightMomCompartment *compartment = new TrafficLightMomCompartment(static_cast<int>(TrafficLightMomState::WORKING));

            compartment->_forwardEvent = &e;

            this->_transition_(compartment);
            return;
        }
        else if (e._message == "systemError")
        {
            // System Error
            TrafficLightMomCompartment *compartment = new TrafficLightMomCompartment(static_cast<int>(TrafficLightMomState::WORKING));

            compartment->_forwardEvent = &e;

            this->_transition_(compartment);
            return;
        }
        else if (e._message == "stop")
        {
            // Stop
            TrafficLightMomCompartment *compartment = new TrafficLightMomCompartment(static_cast<int>(TrafficLightMomState::END));

            this->_transition_(compartment);
            return;
        }
    }

    void _sWorking_(FrameEvent e)
    {
        if (e._message == ">")
        {
            TrafficLightController* controller;
            this->_trafficLight = controller->loadTrafficLight((TrafficLightMomController*)this, (TrafficLightController*)this->_data);
            return;
        }
        else if (e._message == "tick")
        {
            this->trafficLight->tick();
            // Done
            TrafficLightMomCompartment *compartment = new TrafficLightMomCompartment((int)TrafficLightMomState::SAVING);

            this->_transition_(compartment);
            return;
        }
        else if (e._message == "systemError")
        {
            this->trafficLight->systemError();
            // Done
            TrafficLightMomCompartment *compartment = new TrafficLightMomCompartment((int)TrafficLightMomState::SAVING);

            this->_transition_(compartment);
            return;
        }
        _sTrafficLightApi_(e);
    }

    void _sTrafficLightApi_(FrameEvent e)
    {
        if (e._message == "enterRed")
        {
            enterRed_do();
            return;
        }
        else if (e._message == "enterGreen")
        {
            enterGreen_do();
            return;
        }
        else if (e._message == "enterYellow")
        {
            enterYellow_do();
            return;
        }
        else if (e._message == "enterFlashingRed")
        {
            enterFlashingRed_do();
            return;
        }
        else if (e._message == "exitFlashingRed")
        {
            exitFlashingRed_do();
            return;
        }
        else if (e._message == "startWorkingTimer")
        {
            startWorkingTimer_do();
            return;
        }
        else if (e._message == "stopWorkingTimer")
        {
            stopWorkingTimer_do();
            return;
        }
        else if (e._message == "startFlashingTimer")
        {
            startFlashingTimer_do();
            return;
        }
        else if (e._message == "stopFlashingTimer")
        {
            stopFlashingTimer_do();
            return;
        }
        else if (e._message == "changeColor")
        {
            changeColor_do(std::any_cast<std::string>(e._parameters["color"]));
            return;
        }
        else if (e._message == "startFlashing")
        {
            startFlashing_do();
            return;
        }
        else if (e._message == "stopFlashing")
        {
            stopFlashing_do();
            return;
        }
        else if (e._message == "changeFlashingAnimation")
        {
            changeFlashingAnimation_do();
            return;
        }
        else if (e._message == "systemError")
        {
            systemError_do();
            return;
        }
        else if (e._message == "systemRestart")
        {
            systemRestart_do();
            return;
        }
        else if (e._message == "log")
        {
            log_do(std::any_cast<std::string>(e._parameters["msg"]));
            return;
        }
    }

    void _sEnd_(FrameEvent e)
    {
        if (e._message == ">")
        {
            TrafficLightController* controller;
            this->_trafficLight = controller->loadTrafficLight((TrafficLightMomController*)this, (TrafficLightController*)this->_data);
            this->trafficLight->stop();
            this->trafficLight = nullptr;
            return;
        }
        _sTrafficLightApi_(e);
    };

    //===================== Actions Block ===================//

    // Unimplemented Actions
public:
    void enterRed_do();
    void enterGreen_do();
    void enterYellow_do();
   
    void enterFlashingRed_do() { throw std::logic_error("Not implemented"); }
    void exitFlashingRed_do() { throw std::logic_error("Not implemented"); }
    void startWorkingTimer_do() { throw std::logic_error("Not implemented"); }
    void stopWorkingTimer_do() { throw std::logic_error("Not implemented"); }
    void startFlashingTimer_do() { throw std::logic_error("Not implemented"); }
    void stopFlashingTimer_do() { throw std::logic_error("Not implemented"); }
    void changeColor_do(std::string color) { throw std::logic_error("Not implemented"); }
    void startFlashing_do() { throw std::logic_error("Not implemented"); }
    void stopFlashing_do() { throw std::logic_error("Not implemented"); }
    void changeFlashingAnimation_do() { throw std::logic_error("Not implemented"); }
    void systemError_do() { throw std::logic_error("Not implemented"); }
    void systemRestart_do() { throw std::logic_error("Not implemented"); }
    void log_do(std::string msg) { throw std::logic_error("Not implemented"); }

    //===================== Domain Block ===================//

public:
    TrafficLightMomController *trafficLight = nullptr;

    TrafficLight *_data = nullptr;

    TrafficLight *_trafficLight = nullptr;

    //=============== Machinery and Mechanisms ==============//

private:
    int _state_;

    void _transition_(TrafficLightMomCompartment *compartment)
    {
        _nextCompartment_ = compartment;
    }

    void _doTransition_(TrafficLightMomCompartment *nextCompartment)
    {
        this->_mux_(FrameEvent("<", _compartment_->exitArgs));
        this->_compartment_ = nextCompartment;
        this->_mux_(FrameEvent(">", this->_compartment_->enterArgs));
    }
};

class TrafficLightMomController : public TrafficLightMom
{
public:
    TrafficLightMomController() : TrafficLightMom() {}

public: 
    void enterRed_do();

    void enterGreen_do();

    void enterYellow_do();
    
    void enterFlashingRed_do() {}

    void exitFlashingRed_do() {}

    void startWorkingTimer_do() {}

    void stopWorkingTimer_do() {}

    void startFlashingTimer_do() {}

    void stopFlashingTimer_do() {}

    void changeColor_do(const std::string &color) {}

    void startFlashing_do() {}

    void stopFlashing_do() {}

    void changeFlashingAnimation_do() {}

    void log_do(const std::string &msg) {}
};

void enterRed_do() {
        std::cout << "Red" << std::endl;
    }

    void enterGreen_do() {
        std::cout << "Green" << std::endl;
    }

    void enterYellow_do() {
        std::cout << "Yellow" << std::endl;
    }
 #endif


