```
#include <iostream>
#include <unordered_map>
#include <any>
#include <string>
#include "frameEvent.h"
using namespace std;
```
#SystemParams $[stateMsg:string] >[enterMsg:string]

-machine-

$Begin [stateMsg:string]
|>|[enterMsg:string]
print(stateMsg + " " + enterMsg) ^

-actions-

print    [msg:string] {`cout << msg << " ";`}

##