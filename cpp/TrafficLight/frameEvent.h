#ifndef FRAME_EVENT_H
#define FRAME_EVENT_H

#include <unordered_map>
#include <string>
#include <any>

using namespace std;

class FrameEvent {
public:
    std::string _message;
    std::unordered_map<std::string, std::any> _parameters;
    void* _return = nullptr;

    FrameEvent(std::string message, std::unordered_map<std::string, std::any> parameters) :
        _message(message), _parameters(parameters) {}

    FrameEvent() {}
};

#endif // FRAME_EVENT_H
