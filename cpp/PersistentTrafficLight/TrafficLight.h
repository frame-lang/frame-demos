// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files
// #ifndef TRAFFICLIGHT_H
// #define TRAFFICLIGHT_H
#ifndef TRAFFICLIGHT_H
#define TRAFFICLIGHT_H

#pragma once
#include <iostream>
#include <unordered_map>
#include <string>
#include <any>
#include "TrafficLightMom.h"
#include "frameEvent.h"


class TrafficLight;
class TrafficLightMom;
class TrafficLightController;
class TrafficLightMomController;
//=============== Compartment ==============//

class TrafficLightCompartment
{
public:
    int state;

    TrafficLightCompartment(int state)
    {
        this->state = state;
    }

    std::unordered_map<std::string, std::any> stateArgs;
    std::unordered_map<std::string, std::any> stateVars;
    std::unordered_map<std::string, std::any> enterArgs;
    std::unordered_map<std::string, std::any> exitArgs;
    FrameEvent *_forwardEvent = nullptr;
};

class TrafficLight
{
private:
    TrafficLightCompartment *_compartment_;
    TrafficLightCompartment *_nextCompartment_;
public:
    TrafficLightMom *_manager;

public:
    TrafficLight(TrafficLightMom *manager)
    {
        // Create and intialize start state compartment.
        _state_ = static_cast<int>(TrafficLightState::BEGIN);

        _manager = manager;
        _compartment_ = new TrafficLightCompartment(this->_state_);
        _nextCompartment_ = nullptr;

        // Send system start event
        FrameEvent frame_event(">", std::unordered_map<std::string, std::any>());
        _mux_(frame_event);
    }

    // states enum
private:
    enum class TrafficLightState
    {
        BEGIN = 0,
        RED = 1,
        GREEN = 2,
        YELLOW = 3,
        FLASHINGRED = 4,
        END = 5,
        WORKING = 6
    };

public:
    static TrafficLightController *loadTrafficLight(TrafficLightMomController *manager, TrafficLightController *data);

public:
    TrafficLight *marshal()
    {

        TrafficLight *data = this;
        return data;
    }

    //====================== Multiplexer ====================//

private:
    void _mux_(FrameEvent e)
    {
        if (this->_compartment_->state == static_cast<int>(TrafficLightState::BEGIN))
        {
            this->_sBegin_(e);
        }
        else if (this->_compartment_->state == static_cast<int>(TrafficLightState::RED))
        {
            this->_sRed_(e);
        }
        else if (this->_compartment_->state == static_cast<int>(TrafficLightState::GREEN))
        {
            this->_sGreen_(e);
        }
        else if (this->_compartment_->state == static_cast<int>(TrafficLightState::YELLOW))
        {
            this->_sYellow_(e);
        }
        else if (this->_compartment_->state == static_cast<int>(TrafficLightState::FLASHINGRED))
        {
            this->_sFlashingRed_(e);
        }
        else if (this->_compartment_->state == static_cast<int>(TrafficLightState::END))
        {
            this->_sEnd_(e);
        }
        else if (this->_compartment_->state == static_cast<int>(TrafficLightState::WORKING))
        {
            this->_sWorking_(e);
        }

        if (this->_nextCompartment_ != nullptr)
        {
            TrafficLightCompartment *nextCompartment = this->_nextCompartment_;
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

public:
    //===================== Interface Block ===================//

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

    //===================== Machine Block ===================//

private:
    void _sBegin_(FrameEvent e)
    {
        if (e._message == ">")
        {
            startWorkingTimer_do();
            TrafficLightCompartment *compartment = new TrafficLightCompartment(static_cast<int>(TrafficLightState::RED));

            this->_transition_(compartment);
            return;
        }
    }

    void _sRed_(FrameEvent e)
    {
        if (e._message == ">")
        {
            enterRed_do();
            return;
        }
        else if (e._message == "tick")
        {
            TrafficLightCompartment *compartment = new TrafficLightCompartment(static_cast<int>(TrafficLightState::GREEN));

            this->_transition_(compartment);
            return;
        }
        _sWorking_(e);
    }

    void _sGreen_(FrameEvent e)
    {
        if (e._message == ">")
        {
            enterGreen_do();
            return;
        }
        else if (e._message == "tick")
        {
            TrafficLightCompartment *compartment = new TrafficLightCompartment(static_cast<int>(TrafficLightState::YELLOW));

            this->_transition_(compartment);
            return;
        }
        _sWorking_(e);
    }

    void _sYellow_(FrameEvent e)
    {
        if (e._message == ">")
        {
            enterYellow_do();
            return;
        }
        else if (e._message == "tick")
        {
            TrafficLightCompartment *compartment = new TrafficLightCompartment(static_cast<int>(TrafficLightState::RED));

            this->_transition_(compartment);
            return;
        }
        _sWorking_(e);
    }

    void _sFlashingRed_(FrameEvent e)
    {
        if (e._message == ">")
        {
            enterFlashingRed_do();
            stopWorkingTimer_do();
            startFlashingTimer_do();
            return;
        }
        else if (e._message == "<")
        {
            exitFlashingRed_do();
            stopFlashingTimer_do();
            startWorkingTimer_do();
            return;
        }
        else if (e._message == "tick")
        {
            changeFlashingAnimation_do();
            return;
        }
        else if (e._message == "systemRestart")
        {
            TrafficLightCompartment *compartment = new TrafficLightCompartment(static_cast<int>(TrafficLightState::RED));

            this->_transition_(compartment);
            return;
        }
        else if (e._message == "stop")
        {
            TrafficLightCompartment *compartment = new TrafficLightCompartment(static_cast<int>(TrafficLightState::END));

            this->_transition_(compartment);
            return;
        }
    }

    void _sEnd_(FrameEvent e)
    {
        if (e._message == ">")
        {
            stopWorkingTimer_do();
            return;
        }
    }

    void _sWorking_(FrameEvent e)
    {
        if (e._message == "stop")
        {
            TrafficLightCompartment *compartment = new TrafficLightCompartment(static_cast<int>(TrafficLightState::END));

            this->_transition_(compartment);
            return;
        }
        else if (e._message == "systemError")
        {
            TrafficLightCompartment *compartment = new TrafficLightCompartment(static_cast<int>(TrafficLightState::FLASHINGRED));

            this->_transition_(compartment);
            return;
        }
    }

    //===================== Actions Block ===================//
public:
    void enterRed_do();

public:
    void enterGreen_do();
    
public:
    void enterYellow_do();
    

    // Unimplemented Actions

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
    void log_do(std::string msg) { throw std::logic_error("Not implemented"); }

    //===================== Domain Block ===================//

public:
    std::string flashColor = "";

    //=============== Machinery and Mechanisms ==============//

private:
    int _state_;

    void _transition_(TrafficLightCompartment *compartment)
    {
        _nextCompartment_ = compartment;
    }

    void _doTransition_(TrafficLightCompartment *nextCompartment)
    {
        this->_mux_(FrameEvent("<", _compartment_->exitArgs));
        this->_compartment_ = nextCompartment;
        this->_mux_(FrameEvent(">", this->_compartment_->enterArgs));
    }
};
// #endif  // TRAFFICLIGHT_H
class TrafficLightController : public TrafficLight
{
public:
    TrafficLightController(TrafficLightMom *manager) : TrafficLight(manager) {}
public:
    void enterRed_do();
    void enterGreen_do();
    void enterYellow_do();
    void enterFlashingRed_do();
    void exitFlashingRed_do();
    void startWorkingTimer_do();
    void stopWorkingTimer_do();
    void startFlashingTimer_do();
    void stopFlashingTimer_do();
    void changeColor_do(const std::string &color);
    void startFlashing_do();
    void stopFlashing_do();
    void changeFlashingAnimation_do();
    void log_do(const std::string &msg);
};

void TrafficLightController::enterRed_do() 
{

    _manager->enterRed_do();
}

void TrafficLightController::enterGreen_do() 
{
    _manager->enterGreen_do();
}

void TrafficLightController::enterYellow_do()
{
    _manager->enterYellow_do();
}

void TrafficLightController::enterFlashingRed_do() {}

void TrafficLightController::exitFlashingRed_do() {}

void TrafficLightController::startWorkingTimer_do() {}

void TrafficLightController::stopWorkingTimer_do() {}

void TrafficLightController::startFlashingTimer_do() {}

void TrafficLightController::stopFlashingTimer_do() {}

void TrafficLightController::changeColor_do(const std::string &color) {}

void TrafficLightController::startFlashing_do() {}

void TrafficLightController::stopFlashing_do() {}

void TrafficLightController::changeFlashingAnimation_do() {}

void TrafficLightController::log_do(const std::string &msg) {}

TrafficLightController *TrafficLight::loadTrafficLight(TrafficLightMomController *manager, TrafficLightController *data)
{

    data->_manager = manager;
    return data;
}

#endif  // TRAFFICLIGHT_H
