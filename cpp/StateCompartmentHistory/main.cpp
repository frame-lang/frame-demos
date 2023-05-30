#include "machine.h"
#include "stateHistoryBasic.h"
using namespace std;

int main()
{
    HistoryBasicController basic;
    basic.start();
    basic.switchState();
    basic.gotoDeadEnd();
    basic.back();

    HistoryStateContextController compartment;
    compartment.start();
    compartment.switchState();
    compartment.gotoDeadEnd();
    compartment.back();

    return 0;
}