// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files
#include <iostream>
#include <unordered_map>
#include <string>
#include <any>
#include <thread> // for std::thread
#include <chrono> // for std::this_thread::sleep_for
#include "frameEvent.h"

// #include "TrafficLightController.h"

//=============== Compartment ==============//

class TrafficLightCompartment
{
public:
    int state;

    TrafficLightCompartment(int state)
    {
        this->state = state;
    }

    std::unordered_map<std::string, any> stateArgs;
    std::unordered_map<std::string, any> stateVars;
    std::unordered_map<std::string, any> enterArgs;
    std::unordered_map<std::string, any> exitArgs;
    FrameEvent *_forwardEvent = nullptr;
};

class TrafficLight
{
private:
    TrafficLightCompartment *_compartment_;
    TrafficLightCompartment *_nextCompartment_;

public:
    TrafficLight()
    {
        // Create and intialize start state compartment.
        _state_ = static_cast<int>(TrafficLightState::BEGIN);
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
    };

    //===================== Interface Block ===================//

public:
    void start(TrafficLight *mom)
    {
        std::unordered_map<std::string, any> parameters;
        parameters["mom"] = mom;

        FrameEvent e(">>", parameters);
        this->_mux_(e);
    }

    // ...

    void stop()
    {
        FrameEvent e("stop", unordered_map<string, any>());
        this->_mux_(e);
    }

    void tick()
    {
        FrameEvent e("tick", unordered_map<string, any>());
        this->_mux_(e);
    }

    void systemError()
    {
        FrameEvent e("systemError", unordered_map<string, any>());
        this->_mux_(e);
    }

    void systemRestart()
    {
        FrameEvent e("systemRestart", unordered_map<string, any>());
        this->_mux_(e);
    }

    //===================== Machine Block ===================//

private:
    void _sBegin_(FrameEvent e)
    {
        if (e._message == ">>")
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
    void enterRed_do()
    {
        std::cout << "enterRed_do()";
    }

public:
    void enterGreen_do()
    {
        std::cout << "enterGreen_do()";
    }

public:
    void enterYellow_do()
    {
        std::cout << "enterYellow_do()";
    }

private:
    std::thread *ticker = nullptr;

public:
    void startWorkingTimer_do()
    {
        this->ticker = new std::thread([this]()
                                       {
        while (true) {
            std::this_thread::sleep_for(std::chrono::milliseconds(1000));
            this->tick();
        } });
    }

    void stopWorkingTimer_do()
    {
        if (this->ticker != nullptr)
        {
            this->ticker->join();
            delete this->ticker;
            this->ticker = nullptr;
        }
    }

    // Unimplemented Actions

protected:
    void enterFlashingRed_do() { throw std::logic_error("Not implemented"); }
    void exitFlashingRed_do() { throw std::logic_error("Not implemented"); }
    void startFlashingTimer_do() { throw std::logic_error("Not implemented"); }
    void stopFlashingTimer_do() { throw std::logic_error("Not implemented"); }
    void changeColor_do(std::string color) { throw std::logic_error("Not implemented"); }
    void startFlashing_do() { throw std::logic_error("Not implemented"); }
    void stopFlashing_do() { throw std::logic_error("Not implemented"); }
    void changeFlashingAnimation_do() { throw std::logic_error("Not implemented"); }
    void log_do(std::string msg) { throw std::logic_error("Not implemented"); }

    //===================== Domain Block ===================//

    // public:
    // string flashColor  = "";

    // public:
    // Mom mom  = null;

    // public:
    // *time.Ticker ticker  = null;

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

// #ifndef TRAFFIC_LIGHT_CONTROLLER_H
// #define TRAFFIC_LIGHT_CONTROLLER_H
// namespace myNamespace
// {
class TrafficLightController : public TrafficLight
{
public:
    TrafficLightController() : TrafficLight() {}

    // protected:
    //     void enterFlashingRed_do() {}
    //     void exitFlashingRed_do() {}
    //     void startFlashingTimer_do() {}
    //     void stopFlashingTimer_do() {}
    //     void changeColor_do(std::string color) {}
    //     void startFlashing_do() {}
    //     void stopFlashing_do() {}
    //     void changeFlashingAnimation_do() {}
    //     void log_do(std::string msg) {}
};
//}
// #endif // TRAFFIC_LIGHT_CONTROLLER_H
