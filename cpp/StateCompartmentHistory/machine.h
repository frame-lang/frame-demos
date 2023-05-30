// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files
#include <iostream>
#include <unordered_map>
#include <string>
#include <stack>
#include <any>
#include "frameEvent.h"
using namespace std;

//=============== Compartment ==============//

class HistoryStateContextCompartment
{
public:
    int state;
    
    HistoryStateContextCompartment(int state)
    {
        this->state = state;
    }
    
    std::unordered_map<std::string, std::any> stateArgs;
    std::unordered_map<std::string, std::any> stateVars;
    std::unordered_map<std::string, std::any> enterArgs;
    std::unordered_map<std::string, std::any> exitArgs;
    FrameEvent *_forwardEvent = nullptr;
};

class HistoryStateContext
{
private:
    HistoryStateContextCompartment *_compartment_;
    HistoryStateContextCompartment *_nextCompartment_;
    
    
    
public:
    HistoryStateContext()
    {
        // Create and intialize start state compartment.
        _state_ = static_cast<int>(HistoryStateContextState::START);
        // Create state stack.
        
        this->_stateStack_ = new stack<HistoryStateContextCompartment>();
        
        
        _compartment_ = new HistoryStateContextCompartment(this->_state_);
        _nextCompartment_ = nullptr;
        
        
        // Send system start event
        FrameEvent frame_event(">", std::unordered_map<std::string, std::any>());
        _mux_(&frame_event);
    
    }
    
    // states enum
private:
    enum class HistoryStateContextState
    {
        START = 0,
        S0 = 1,
        S1 = 2,
        DEADEND = 3
    };
    
    //====================== Multiplexer ====================//

private:
    void _mux_(FrameEvent *e)
    {
        if(this->_compartment_->state == static_cast<int>(HistoryStateContextState::START))
        {
            this->_sStart_(e);
        }
        else if(this->_compartment_->state == static_cast<int>(HistoryStateContextState::S0))
        {
            this->_sS0_(e);
        }
        else if(this->_compartment_->state == static_cast<int>(HistoryStateContextState::S1))
        {
            this->_sS1_(e);
        }
        else if(this->_compartment_->state == static_cast<int>(HistoryStateContextState::DEADEND))
        {
            this->_sDeadEnd_(e);
        }
        
        
        if(this->_nextCompartment_ != nullptr)
        {
            HistoryStateContextCompartment *nextCompartment = this->_nextCompartment_;
            this->_nextCompartment_ = nullptr;
            if (nextCompartment->_forwardEvent != nullptr && 
                nextCompartment->_forwardEvent->_message == ">")
            {
                this->_mux_(new FrameEvent( "<", this->_compartment_->exitArgs));
                this->_compartment_ = nextCompartment;
                this->_mux_(nextCompartment->_forwardEvent);
            }
            else
            {
                this->_doTransition_(nextCompartment);
                if (nextCompartment->_forwardEvent != nullptr)
                {
                    this->_mux_(nextCompartment->_forwardEvent);
                }
            }
            nextCompartment->_forwardEvent = nullptr;
        }
    }
    
        
        //===================== Interface Block ===================//
    public:
        void start() {
            FrameEvent e(">>", unordered_map<string, any>());
            this->_mux_(&e);
        }
        
        void switchState() {
            FrameEvent e("switchState", unordered_map<string, any>());
            this->_mux_(&e);
        }
        
        void gotoDeadEnd() {
            FrameEvent e("gotoDeadEnd", unordered_map<string, any>());
            this->_mux_(&e);
        }
        
        void back() {
            FrameEvent e("back", unordered_map<string, any>());
            this->_mux_(&e);
        }
        
    
    //===================== Machine Block ===================//

private:
    
    void _sStart_(FrameEvent *e)
    {
        if (e->_message == ">>") {
            HistoryStateContextCompartment *compartment =  new HistoryStateContextCompartment(static_cast<int>(HistoryStateContextState::S0));
            compartment->stateVars["enterMsg"] = string("Enter S0");
            
            this->_transition_(compartment);
            return;
        }
    }
    
    void _sS0_(FrameEvent *e)
    {
        if (e->_message == ">") {
            print_do((any_cast<string>(this->_compartment_->stateVars["enterMsg"])));
            return;
        }
        else if (e->_message == "switchState") {
            // Switch\nState
            HistoryStateContextCompartment *compartment =  new HistoryStateContextCompartment(static_cast<int>(HistoryStateContextState::S1));
            
            this->_transition_(compartment);
            return;
        }
        else if (e->_message == "gotoDeadEnd") {
            _stateStack_push_(_compartment_);
            // Goto\nDead End
            HistoryStateContextCompartment *compartment =  new HistoryStateContextCompartment(static_cast<int>(HistoryStateContextState::DEADEND));
            
            this->_transition_(compartment);
            return;
        }
    }
    
    void _sS1_(FrameEvent *e)
    {
        if (e->_message == ">") {
            print_do("Enter S1");
            return;
        }
        else if (e->_message == "switchState") {
            // Switch\nState
            HistoryStateContextCompartment *compartment =  new HistoryStateContextCompartment(static_cast<int>(HistoryStateContextState::S0));
            compartment->stateVars["enterMsg"] = string("Enter S0");
            
            this->_transition_(compartment);
            return;
        }
        else if (e->_message == "gotoDeadEnd") {
            _stateStack_push_(_compartment_);
            // Goto\nDead End
            HistoryStateContextCompartment *compartment =  new HistoryStateContextCompartment(static_cast<int>(HistoryStateContextState::DEADEND));
            
            this->_transition_(compartment);
            return;
        }
    }
    
    void _sDeadEnd_(FrameEvent *e)
    {
        if (e->_message == ">") {
            print_do("Enter $DeadEnd");
            return;
        }
        else if (e->_message == "back") {
            // Go Back
            HistoryStateContextCompartment *compartment = this->_stateStack_pop_();
            this->_transition_(compartment);
            return;
        }
    }

//===================== Actions Block ===================//
    
    
    
public:
    void print_do(string msg)
    {
        cout << msg << endl;
    }
    
    // Unimplemented Actions
    
    
    
    //=============== Machinery and Mechanisms ==============//

private:
    int _state_;
    
    void _transition_(HistoryStateContextCompartment *compartment)
    {
        _nextCompartment_ = compartment;
    }
    
    void _doTransition_(HistoryStateContextCompartment *nextCompartment)
    {
        this->_mux_(new FrameEvent("<", _compartment_->exitArgs));
        this->_compartment_ = nextCompartment;
        this->_mux_(new FrameEvent(">", this->_compartment_->enterArgs));
    }
    
    private:
        std::stack<HistoryStateContextCompartment> *_stateStack_;
    
    public:
        void _stateStack_push_(HistoryStateContextCompartment *compartment)
        {
            _stateStack_->push(*compartment);
        }
        
        
        HistoryStateContextCompartment *_stateStack_pop_()
        {
            HistoryStateContextCompartment *copyCompartment = &_stateStack_->top();
            _stateStack_->pop();
            return copyCompartment;
        }
    
    
};

class HistoryStateContextController : public HistoryStateContext
{
public:
	HistoryStateContextController() : HistoryStateContext() {}
};


















